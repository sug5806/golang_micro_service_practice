package oauth

import "time"

type AccessToken struct {
	AccessToken string    `json:"access_token"`
	UserId      int64     `json:"user_id"`
	Expires     time.Time `json:"expires"`
}
