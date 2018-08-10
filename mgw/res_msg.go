package mgw

// MessageEnum is enum for response
type MessageEnum string

const (
	// MsgSuccess message
	MsgSuccess MessageEnum = "success"
	// MsgError message
	MsgError MessageEnum = "error"
)

// ResMsg is a response message struct
type ResMsg struct {
	Code    int         `json:"code"`
	Message MessageEnum `json:"message"`
	Data    string      `json:"data"`
}

// StatusOK is success for response
func StatusOK(data string) ResMsg {
	return ResMsg{
		Code:    0,
		Message: MsgSuccess,
		Data:    data,
	}
}

// StatusFailure is error for response
func StatusFailure(data string) ResMsg {
	return ResMsg{
		Code:    -1,
		Message: MsgError,
		Data:    data,
	}
}
