package swagger

import "eat_box/internal/model"

type CreateOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type GetOrderResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data model.Order `json:"data"`
}
