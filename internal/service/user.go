package service

import (
	"eat_box/global"
	"eat_box/internal/model/swagger"
	"eat_box/pkg/app"
	"eat_box/pkg/errcode"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Code string `form:"code" binding:"required"`
}
type UpdateInfoRequest struct {
	ID        string
	Nickname  string `form:"nickname" binding:"min=1,max=30"`
	Tele      string `form:"tele" binding:"len=11"`
	HeadImage string `form:"headimage" binding:"url"`
}
type DetailRequest struct {
	ID   string `form:"id" binding:"len=10"`
	Self bool
}
type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (svc *Service) WXLogin(code string) (*WXLoginResp, *errcode.Error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, global.WechatSetting.AppID, global.WechatSetting.AppSecret, code)
	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, errcode.ServerError
	}
	defer resp.Body.Close()
	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, errcode.JSONUnmarshalError
	}
	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errcode.WXAPIError
	}
	return &wxResp, errcode.Success
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
	user, ok := svc.dao.FindUserByID(params.ID)
	if !ok {
		return errcode.NotFound
	}
	if params.Nickname != "" {
		user.UpdateNickName(params.Nickname)
	}
	if params.Tele != "" {
		user.UpdateTele(params.Tele)
	}
	if params.HeadImage != "" {
		user.UpdateHeadImage(params.HeadImage)
	}
	err := svc.dao.UpdateUserInfo(user)
	if err != nil {
		return errcode.MySQLErr
	}
	return errcode.Success
}
func (svc *Service) GetUserInfo(params *DetailRequest) (swagger.DetailData, *errcode.Error) {
	user, ok := svc.dao.FindUserByID(params.ID)
	if !ok {
		return swagger.DetailData{}, errcode.NotFound
	}
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
