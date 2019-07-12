package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "请求参数错误",
	Unauthorized:                   "未授权",
	Created:                        "创建成功",
	Accepted:                       "请求被接受",
	No_Content:                     "No Content",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
