package users

// User Model of user
type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ID          int64  `json:"Id"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}
