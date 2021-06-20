package entity

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Birthday   string `json:"birthday"`
	Status     bool   `json:"status"`
	CreatedAt  int64  `json:"created_at"`
	ModifiedAt int64  `json:"modified_at"`
}
