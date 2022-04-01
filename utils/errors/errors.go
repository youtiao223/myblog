package errors

const (
	SUCCESS = 200
	ERROR   = 500

	ErrorUserNameExits  = 4000
	ErrorNameOrPwd      = 4001
	ErrorUserIdNotExits = 4002

	ErrorCateExits      = 4100
	ErrorCateIdNotExits = 4101

	ErrorArtIdNotExits = 4200

	ErrorTokenExits    = 4300
	ErrorTokenTime     = 4301
	ErrorTokenFormat   = 4302
	ErrorTokenValidate = 4303

	ErrorUploadFile = 4400
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ErrorUserNameExits:  "用户名已存在",
	ErrorUserIdNotExits: "用户不存在",
	ErrorNameOrPwd:      "用户名或密码错误",

	ErrorCateExits:      "分类名已存在",
	ErrorCateIdNotExits: "分类不存在",

	ErrorArtIdNotExits: "文章不存在",

	ErrorTokenExits:    "Token 不存在",
	ErrorTokenTime:     "Token 已过期",
	ErrorTokenFormat:   "Token 格式错误",
	ErrorTokenValidate: "Token 验证错误",

	ErrorUploadFile: "文件上传错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
