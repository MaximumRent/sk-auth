package worker

type AuthRequestMessage struct {
	MessageSource string `json:"message_source"`
	Email         string `json:"email"`
	Token         string `json:"token"`
}

type AuthResponseMessage struct {
}
