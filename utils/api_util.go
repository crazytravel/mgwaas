package utils

// ResMsg is a response message struct
type ResMsg struct {
	Code    int             `json:"code"`
	Message ResponseMessage `json:"message"`
	Data    string          `json:"data"`
}

// StatusOK is success for response
func StatusOK(data string) ResMsg {
	return ResMsg{
		Code:    0,
		Message: ResSuccess,
		Data:    data,
	}
}

// StatusFailure is error for response
func StatusFailure(data string) ResMsg {
	return ResMsg{
		Code:    -1,
		Message: ResError,
		Data:    data,
	}
}
