package User

type User struct {
	ID        int               `json:"user_id"`
	Username     string         `json:"Username"`
	Password     string         `json:"password"`
	Secret       string         `json:"secret"`
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Email     string            `json:"Email"`
	OrganizationName  string    `json:"OrganizationName"`
	Created  string            `json:"Created"`
	Updated  string            `json:"Updated"`
}

type UserResponse struct {
	ID        int               `json:"user_id"`
	Username     string         `json:"-"`
	Password     string         `json:"-"`
	Secret       string         `json:"secret"`
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Email     string            `json:"Email"`
	OrganizationName  string    `json:"OrganizationName"`
	Created  string            `json:"Created"`
	Updated  string            `json:"Updated"`
}