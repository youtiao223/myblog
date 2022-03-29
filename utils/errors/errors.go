package errors

const (
	SUCCESS        = 200
	ERROR          = 500
	ErrorUserExits = 4000
	ErrorNameOrPwd = 4001
)

var codeMsg = map[int]string{
	SUCCESS:        "OK",
	ERROR:          "FAIL",
	ErrorUserExits: "用户名已存在",
	ErrorNameOrPwd: "用户名或密码错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
