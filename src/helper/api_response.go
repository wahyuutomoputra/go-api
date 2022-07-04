package helper

type Meta struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func ApiResponse(code int, status string, data interface{}) Response {
	meta := Meta{
		Code:   code,
		Status: status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
