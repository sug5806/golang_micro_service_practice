package github

type CreateRepoRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Private         bool   `json:"private"`
	LicenseTemplate string `json:"license_template"`
}

type CreateRepoResponse struct {
	Id         uint           `json:"id"`
	NodeId     string         `json:"node_id"`
	Name       string         `json:"name"`
	FullName   string         `json:"full_name"`
	Owner      RepoOwner      `json:"owner"`
	Private    bool           `json:"private"`
	Permission RepoPermission `json:"permission"`
}

type RepoOwner struct {
	Id     uint   `json:"id"`
	Login  string `json:"login"`
	Type   string `json:"type"`
	NodeId string `json:"node_id"`
}

type RepoPermission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
