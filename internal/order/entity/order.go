package entity

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	TotalPrice float64 `json:"totalPrice"`
}
