package models

import "time"

type Method struct {
	Id       int
	Name     string
	AddPrice int
}

type Size struct {
	Id       int
	Name     string
	AddPrice int
}
type Variant struct {
	Id       int
	Name     string
	AddPrice int
}

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	BasePrice int    `json:"base_price"`
	Variants  []int  `json:"variants"`
	Sizes     []int  `json:"sizes"`
	Methods   []int  `json:"methods"`
	Stock     int    `json:"stock"`
}

type CartItem struct {
	Id        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	VariantID int `json:"variant_id"`
	SizeID    int `json:"size_id"`
	MethodID  int `json:"method_id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
}

type OrderItem struct {
	Id        int
	OrderID   int
	ProductID int
	Qty       int
	Price     int
}

type Order struct {
	Id            int         `json:"id"`
	UserID        int         `json:"user_id"`
	Total         int         `json:"total"`
	Status        string      `json:"status"`
	Address       string      `json:"address"`
	PaymentMethod string      `json:"payment_method"`
	CreatedAt     time.Time   `json:"created_at"`
	Items         []OrderItem `json:"items"`
}
