package errutils

import (
	cons "github.com/zpdev/zins/common/constance"
)

type ZinError struct {
	Errno int
	Msg   string
	CnMsg string
}

func (err *ZinError) Error() string {
	return err.Msg
}

func (err *ZinError) GetMsg(lang string) string {
	if lang == cons.LanguageCN {
		return err.CnMsg
	}
	return err.Msg
}

func JsonFormatError() *ZinError {
	return &ZinError{
		Errno: 1001,
		Msg:   "JSON format error",
		CnMsg: "JSON格式化失败",
	}
}

func UserAlreadyExit(username string) *ZinError {
	return &ZinError{
		Errno: 1002,
		Msg:   "User " + username + " already exist",
		CnMsg: "用户" + username + "已经存在",
	}
}

func EmailAlreadyExit(email string) *ZinError {
	return &ZinError{
		Errno: 1003,
		Msg:   "Mail " + email + " already exist",
		CnMsg: "邮箱" + email + "已经存在",
	}
}

func SpecifiedUserNotFound(username string) *ZinError {
	return &ZinError{
		Errno: 1004,
		Msg:   "User " + username + " not exist",
		CnMsg: "用户" + username + "不存在",
	}
}

func UserNotFound() *ZinError {
	return &ZinError{
		Errno: 1004,
		Msg:   "User not exist",
		CnMsg: "用户不存在",
	}
}

func PasswordEncryptError() *ZinError {
	return &ZinError{
		Errno: 1005,
		Msg:   "Encrypt password error",
		CnMsg: "密码加密失败",
	}
}

func PasswordVerifyError() *ZinError {
	return &ZinError{
		Errno: 1006,
		Msg:   "Verify password error",
		CnMsg: "密码校验失败",
	}
}

func UserPassError() *ZinError {
	return &ZinError{
		Errno: 1007,
		Msg:   "Username or Password is not correct",
		CnMsg: "用户名或密码不正确",
	}
}

func LoginFailed() *ZinError {
	return &ZinError{
		Errno: 1008,
		Msg:   "Login failed",
		CnMsg: "登录失败",
	}
}

func DBOperationsFailed() *ZinError {
	return &ZinError{
		Errno: 1009,
		Msg:   "DB Operations failed",
		CnMsg: "数据库操作失败",
	}
}

func InvaildToken() *ZinError {
	return &ZinError{
		Errno: 1010,
		Msg:   "Invaild token",
		CnMsg: "非法的令牌",
	}
}

func NotLogin() *ZinError {
	return &ZinError{
		Errno: 1011,
		Msg:   "Not login",
		CnMsg: "未登录",
	}
}

func LoginTimeOut() *ZinError {
	return &ZinError{
		Errno: 1011,
		Msg:   "Login timeout",
		CnMsg: "登录超时",
	}
}

func PermissionDenied() *ZinError {
	return &ZinError{
		Errno: 1011,
		Msg:   "Permission denied",
		CnMsg: "权限不足",
	}
}
