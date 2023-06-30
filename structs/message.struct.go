package structs

type Payload struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}
