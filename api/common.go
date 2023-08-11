package api

type Common struct {
	Code    int    `json:"code" dc:"响应码，1失败，0成功"`
	Message string `json:"message" dc:"响应信息"`
}
