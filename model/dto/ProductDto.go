package dto

type ProductDto struct {
	Id           int64 `json:"id"`
	ProductId    int64 `json:"product_id"`
	ProductPrice int64 `json:"product_price"`
	CreatedAt    int64 `json:"created_at"`
}
