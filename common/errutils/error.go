package errutils

type ZinError struct {
	Errno int
	Msg   string
}

func (err *ZinError) Error() string {
	return err.Msg
}

func JsonFormatError() *ZinError {
	return &ZinError{
		Errno: 1001,
		Msg:   "Json format error",
	}
}

func UserAlreadyExit() *ZinError {
	return &ZinError{
		Errno: 1002,
		Msg:   "User already exist",
	}
}
