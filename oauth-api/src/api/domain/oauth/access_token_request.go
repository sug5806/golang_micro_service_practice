package oauth

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
