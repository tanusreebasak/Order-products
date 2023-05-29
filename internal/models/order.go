package models

type Order struct {
	ID            string  `json:"id"`
	ProductID     string  `json:"productId"`
	Quantity      int     `json:"quantity"`
	OrderValue    float64 `json:"orderValue"`
	DispatchDate  string  `json:"dispatchDate"`
	Status        string  `json:"status"`
	DiscountApplied bool   `json:"discountApplied"`
}
