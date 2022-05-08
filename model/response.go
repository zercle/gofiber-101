package model

// result object for response
// https://jsonapi.org/examples/#sparse-fieldsets
type ResponseForm struct {
	Success    bool            `json:"success"`
	Result     interface{}     `json:"result"`
	Messages   []string        `json:"messages"`
	Errors     []*ResposeError `json:"errors"`
	ResultInfo *ResultInfo     `json:"result_info,omitempty"`
}

// error object for response
// https://jsonapi.org/examples/#error-objects
type ResposeError struct {
	Code    int         `json:"code"`
	Source  interface{} `json:"source,omitempty"`
	Title   string      `json:"title,omitempty"`
	Message string      `json:"message"`
}

// result info object for response page-based strategy
// https://jsonapi.org/examples/#pagination
type ResultInfo struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Count     int `json:"count"`
	TotalCont int `json:"total_count"`
}
