package models

type Order struct {
	Model
	TransactionId string      `json:"transaction_id" gorm:"null"`
	UserId        uint        `json:"user_id"`
	Code          string      `json:"code"`
	ReferrerEmail string      `json:"referrer_email"`
	FirstName     string      `json:"-"`
	LastName      string      `json:"-"`
	Name          string      `json:"name" gorm:"-"`
	Email         string      `json:"email"`
	Address       string      `json:"address" gorm:"null"`
	City          string      `json:"city" gorm:"null"`
	Country       string      `json:"country" gorm:"null"`
	Zip           string      `json:"zip" gorm:"null"`
	Complete      bool        `json:"-" gorm:"default:false"`
	Total         float64     `json:"total" gorm:"-"`
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

func (order *Order) FullName() string {
	return order.FirstName + " " + order.LastName
}

func (order *Order) GetTotal() float64 {
	var total float64 = 0

	for _, orderItem := range order.OrderItems {
		total += orderItem.Price * float64(orderItem.Quantity)
	}
	return total
}
