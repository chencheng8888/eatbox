package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的用户")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败,token错误")
	UnauthorizedTokenExpired  = NewError(10000005, "鉴权失败,token已过期")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败,token生成失败")
	Unauthorized              = NewError(10000007, "登录失败")
	MySQLErr                  = NewError(10000008, "数据库操作失败")
	TooManyRequests           = NewError(10000009, "请求过多")
	UnauthorizedAuthIsEmpty   = NewError(100000010, "鉴权失败，请求头中的auth为空")
	ToJSONError               = NewError(100000011, "转换成JSON格式错误")
	ErrRedisSet               = NewError(100000012, "redis设置key-value失败")
	JSONUnmarshalError        = NewError(100000013, "JSON unmarshal error")
	ErrRedisDel               = NewError(100000014, "redis删除key-val失败")
	WXAPIError                = NewError(100000015, "微信api调用错误或者code已过期")
	PageInvalid               = NewError(100000016, "页码超限")
)
