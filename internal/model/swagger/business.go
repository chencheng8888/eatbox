package swagger

import "eat_box/internal/model"

type BusinessData struct {
	Businesses  []model.Business `json:"businesses"`   //商家信息
	Total       int64            `json:"total"`        //商家总数
	CurrentPage int64            `json:"current_page"` //当前页数
	PageSize    int64            `json:"page_size"`    //一页的数据量
}
type BusinessListSwagger struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data BusinessData `json:"data"`
}
type ScoreData struct {
	BusinessID int     `json:"business_id"`
	Score      float64 `json:"score"`
}
