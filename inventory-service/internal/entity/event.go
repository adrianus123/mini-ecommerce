package entity

type Event struct {
	Owner  string `json:"owner"`
	Action string `json:"action"`
	Value  []Item `json:"value"`
}

type Item struct {
	ProductID uint `json:"product_id"`
	Qty       uint `json:"qty"`
}
