package structs

type Payload struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}
