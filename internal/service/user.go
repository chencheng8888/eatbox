package service

import (
	"eat_box/global"
	"eat_box/internal/model/swagger"
	"eat_box/pkg/app"
	"eat_box/pkg/errcode"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"net/http"
)

type LoginRequest struct {
	Code string `form:"code" binding:"required"`
}
type UpdateInfoRequest struct {
	ID        string
	Nickname  string `form:"nickname" binding:"omitempty,min=1,max=30"`
	Tele      string `form:"tele" binding:"omitempty,len=11"`
	HeadImage string `form:"headimage" binding:"omitempty,url"`
}
type DetailRequest struct {
	ID   string `form:"id" binding:"omitempty,len=10"`
	Self bool
}
type GetBusinessesRequest struct {
	Page int `form:"page" binding:"required,number,gte=1"`
}
type ScoreRequest struct {
	BusinessID int     `form:"business_id" binding:"required,number"`
	Score      float64 `form:"score" binding:"required,gte=0,lte=10"`
}
type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (svc *Service) WXLogin(code string) (WXLoginResp, *errcode.Error) {
	//var datachan = make(chan WXLoginResp)
	var respchan = make(chan *http.Response)
	var errchan = make(chan *errcode.Error)
	go func() {
		// 合成url, 这里的appId和secret是在微信公众平台上获取的
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", global.WechatSetting.AppID, global.WechatSetting.AppSecret, code)
		// 创建http get请求
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			errchan <- errcode.ServerError
			return
			//return WXLoginResp{}, errcode.ServerError
		}
		errchan <- errcode.Success
		respchan <- resp
	}()
	err := <-errchan
	if err.Code() != errcode.Success.Code() {
		return WXLoginResp{}, err
	}
	resp := <-respchan
	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return WXLoginResp{}, errcode.JSONUnmarshalError
	}
	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return WXLoginResp{}, errcode.WXAPIError
	}
	return wxResp, errcode.Success
}
func (svc *Service) Login(params *LoginRequest) (bool, string, *errcode.Error) {
	first := false
	wxresp, err := svc.WXLogin(params.Code)
	if err.Code() != errcode.Success.Code() {
		return false, "", err
	}
	user, ok := svc.dao.FindUserByOpenID(wxresp.OpenId)
	if !ok {
		first = true
		user.Create(wxresp.OpenId)
		err1 := svc.dao.CreateUser(user)
		if err1 != nil {
			return false, "", errcode.MySQLErr
		}
	}
	token, _ := app.GenerateToken(user.ID)
	return first, token, errcode.Success
}
func (svc *Service) UpdateInfo(params *UpdateInfoRequest) *errcode.Error {
	ok := svc.dao.IsExistUser(params.ID)
	mp := make(map[string]interface{})
	if !ok {
		return errcode.NotFound
	}
	if params.Nickname != "" {
		mp["nick_name"] = params.Nickname
	}
	if params.Tele != "" {
		mp["tele"] = params.Tele
	}
	if params.HeadImage != "" {
		mp["head_image"] = params.HeadImage
	}
	_, err1 := svc.cache.GetUserFromCache(params.ID)
	if err1.Code() == errcode.Success.Code() {
		err2 := svc.cache.DeleteOneUser(params.ID)
		fmt.Println(err2)
	}
	err := svc.dao.UpdateUserInfo(params.ID, mp)
	if err != nil {
		return errcode.MySQLErr
	}
	return errcode.Success
}
func (svc *Service) GetUserInfo(params *DetailRequest) (swagger.DetailData, *errcode.Error) {
	//user := model.NewUser()
	var ok bool
	user, err := svc.cache.GetUserFromCache(params.ID)
	//如果失败，就从数据库中获取user
	if err.Code() != errcode.Success.Code() {
		user, ok = svc.dao.FindUserByID(params.ID)
		if !ok {
			return swagger.DetailData{}, errcode.NotFound
		}
	}
	//如果缓存中没有就设置一个
	go func() {
		if err.Code() == errcode.NotFound.Code() {
			err = svc.cache.SetOneUser(user)
			if err.Code() != errcode.Success.Code() {
				fmt.Println(err)
			}
		}
	}()

	data := swagger.DetailData{
		ID:        user.ID,
		NickName:  user.NickName,
		HeadImage: user.HeadImage,
		Level:     user.Level,
	}
	if params.Self {
		data.Points = user.Points
		data.Tele = user.Tele
	}
	return data, errcode.Success
}
func (svc *Service) GetBusinesses(params *GetBusinessesRequest) (swagger.BusinessData, *errcode.Error) {
	var Maxpage int
	var data swagger.BusinessData

	//先看page是否合法
	var err1 *errcode.Error
	total := svc.dao.GetBusinessNum()
	if total%global.Pagesize == 0 {
		Maxpage = int(total) / global.Pagesize
	} else {
		Maxpage = int(total)/global.Pagesize + 1
	}
	if params.Page > Maxpage {
		return swagger.BusinessData{}, errcode.PageInvalid
	}
	//先从缓存中获取
	data, err1 = svc.cache.GetBusinessData(int64(params.Page))
	if err1.Code() == errcode.Success.Code() {
		return data, errcode.Success
	}
	limit := global.Pagesize
	offset := (params.Page - 1) * limit
	businesses, err := svc.dao.GetBusinesses(limit, offset)
	if err != nil {
		return swagger.BusinessData{}, errcode.MySQLErr
	}
	data = swagger.BusinessData{
		Total:       total,
		Businesses:  businesses,
		CurrentPage: int64(params.Page),
		PageSize:    int64(global.Pagesize),
	}

	//开个协程来设置缓存
	go func() {
		err1 := svc.cache.SetBusinessData(data)
		if err1.Code() != errcode.Success.Code() {
			fmt.Println("set businesslist err")
			fmt.Println(err1)
		}
	}()

	return data, errcode.Success
}
func (svc *Service) Score(params *ScoreRequest) *errcode.Error {
	data := swagger.ScoreData{
		BusinessID: params.BusinessID,
		Score:      params.Score,
	}
	//向kafka发送消息
	content, err := json.Marshal(&data)
	if err != nil {
		return errcode.ToJSONError
	}
	msg := &sarama.ProducerMessage{Topic: "score", Value: sarama.StringEncoder(content), Partition: 0}
	return svc.producer.SendMsg(msg)
}
