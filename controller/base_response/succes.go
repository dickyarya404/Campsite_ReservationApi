package base_response

type BaseSuccessResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(message string, data any) *BaseSuccessResponse {
	return &BaseSuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

type BaseSuccessResponseWithMetadata struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// func NewSuccessResponseWithMetadata(message string, data any, metadata Metadata) *BaseSuccessResponseWithMetadata {
// 	return &BaseSuccessResponseWithMetadata{
// 		Status:   true,
// 		Message:  message,
// 		Metadata: metadata,
// 		Data:     data,
// 	}
// }
