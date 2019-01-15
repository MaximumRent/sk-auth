package api

type Message struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Payload interface{} `json:"payload"`
}
