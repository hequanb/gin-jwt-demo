package code

const (
	SUCCESS = 1001
	FAIL    = 1002

	ERR_AUTH_NULL    = 2001
	ERR_AUTH_INVALID = 2002 // 无效TOKEN
	ERR_AUTH_EXPIRED = 2003 // TOKEN过期

	ERR_LOGIN_USERNAME = 3001 // 用户名不正确
)

var CMsg = map[int]string{
	SUCCESS:            "成功",
	FAIL:               "失败",
	ERR_AUTH_NULL:      "请求头中没有Auth信息",
	ERR_AUTH_INVALID:   "无效的TOKEN",
	ERR_AUTH_EXPIRED:   "Token时间过期",
	ERR_LOGIN_USERNAME: "用户名不正确",
}

func GetMsg(code int) string {
	if s, ok := CMsg[code]; !ok {
		return CMsg[FAIL]
	} else {
		return s
	}
}
