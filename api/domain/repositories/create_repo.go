package repositories

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateRepoResponse struct {
	Id    uint64 `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}
