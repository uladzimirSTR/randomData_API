package objects

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUsersResponse struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Data   []User `json:"data"`
	Error  string `json:"error,omitempty"`
}
