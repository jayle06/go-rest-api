package entity

type Product struct {
	Id         int64    `json:"id"`
	CategoryId int64    `json:"category_id"`
	Image      []string `json:"image"`
	Name       string   `json:"name"`
	Price      int64    `json:"price"`
	IsSale     bool     `json:"is_sale"`
	CreatedAt  int64    `json:"created_at"`
	ModifiedAt int64    `json:"modified_at"`
	Rating     float64  `json:"rating"`
}
