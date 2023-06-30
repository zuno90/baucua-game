package structs

type Levels string

const (
	LOW  Levels = "LOW"
	MID  Levels = "MID"
	HIGH Levels = "HIGH"
)

type Match struct {
	ID    string `json:"id"`
	Level Levels `json:"level"`
}
