package api

type Success struct {
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"message"`
}

func Ok(data interface{}) *Success {
	return &Success{
		Data: data,
		Msg:  "success",
	}
}
