package users

type FindUsersResponse struct {
	Data  *[]User `json:"data"`
	Total *int    `json:"total"`
	Error *string `json:"error"`
}
