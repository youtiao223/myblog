package errors

// 通用错误
const (
	SUCCESS = 200
	ERROR   = 500
)

// 用户模块错误
const (
	ErrorUserExits      = 4000
	ErrorNameOrPwd      = 4001
	ErrorUserIdNotExits = 4002
)

// 分类模块错误
const (
	ErrorCateExits      = 4100
	ErrorCateIdNotExits = 4101
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ErrorUserExits:      "用户名已存在",
	ErrorUserIdNotExits: "用户不存在",
	ErrorNameOrPwd:      "用户名或密码错误",
	ErrorCateExits:      "分类名已存在",
	ErrorCateIdNotExits: "分类不存在",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
