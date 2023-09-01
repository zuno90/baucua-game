package game

type Types string

const (
	// CREATEROOM Types = "CREATEROOM"
	// CHATROOM   Types = "CHATROOM"
	// JOINROOM   Types = "JOINROOM"
	// LEAVEROOM  Types = "LEAVEROOM"
	CHAT     Types = "CHAT"
	WELCOME  Types = "WELCOME"
	GAMEINFO Types = "GAMEINFO"
	RESULT   Types = "RESULT"
	BYEBYE   Types = "BYEBYE"
	LOGIN    Types = "LOGIN"
	LOGOUT   Types = "LOGOUT"
	ERROR    Types = "ERROR"
)

type ReqData struct {
	Type Types  `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

type ResData struct {
	Type Types  `json:"type"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Msg  string `json:"msg"`
}

type ResError struct {
	Type Types  `json:"type"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ResMessage(t Types, msg string) ResData {
	return ResData{
		Type: t,
		Msg:  msg,
	}
}

func ResErrorMessage(t Types, code int, msg string) ResError {
	return ResError{
		Type: t,
		Code: code,
		Msg:  msg,
	}
}
