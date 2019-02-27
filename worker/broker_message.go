package worker

type AuthRequestMessage struct {
	MessageSource string `json:"message_source"`
	Path          string `json:"path"`
	Email         string `json:"email"`
	Nickname      string `json:"nickname"`
	Token         string `json:"token"`
	Access		  string `json:"access"`
}

type AuthResponseMessage struct {
	HasAccess       bool                `json:"has_access"`
	ReturnedMessage *AuthRequestMessage `json:"returned_message"`
}
