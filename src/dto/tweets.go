package dto

type Tweet struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Value string `json:"value"`
}

type GetAllTweetsResponse struct {
	Data  []Tweet `json:"data"`
	Total int     `json:"total"`
}
