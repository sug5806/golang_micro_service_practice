package github

type ErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Errors           []Error
}

func (r ErrorResponse) Error() string {
	return r.Message
}

type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
