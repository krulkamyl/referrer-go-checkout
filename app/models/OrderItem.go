package models

type OrderItem struct {
	Model
	OrderId         uint    `json:"order_id"`
	ProductTitle    string  `json:"product_title"`
	Price           float64 `json:"price"`
	Quantity        uint    `json:"quantity"`
	AdminRevenue    float64 `json:"admin_revenue"`
	ReferrerRevenue float64 `json:"referrer_revenue"`
}
