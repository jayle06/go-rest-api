package entity

type Review struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"product_id"`
	Comment   string `json:"comment"`
	Rating    int64  `json:"rating"`
}
