package response

const SUCCESS = 0
const FAILED = 1

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJsonResponse(data interface{}, message string, statusType int) JsonResponse {
	return JsonResponse{
		Status:  statusType,
		Message: message,
		Data:    data,
	}
}
