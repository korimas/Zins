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
		Errno: 1002,
		Msg:   "Mail " + email + " already exist",
		CnMsg: "邮箱" + email + "已经存在",
	}
}

func UserNotFound(username string) *ZinError {
	return &ZinError{
		Errno: 1002,
		Msg:   "User " + username + " not exist",
		CnMsg: "用户" + username + "不存在",
	}
}

func DBOperationsFailed() *ZinError {
	return &ZinError{
		Errno: 1003,
		Msg:   "DB Operations failed",
		CnMsg: "数据库操作失败",
	}
}
