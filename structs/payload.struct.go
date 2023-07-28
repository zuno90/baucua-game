package structs

type Types string

const (
	CHAT       Types = "CHAT"
	CREATEROOM Types = "CREATEROOM"
	CHATROOM   Types = "CHATROOM"
	JOINROOM   Types = "JOINROOM"
	LEAVEROOM  Types = "LEAVEROOM"
)

type Payload struct {
	Type Types  `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

func CustomErr(from, msg string) Payload {
	return Payload{
		From: from,
		Msg:  msg,
	}
}
