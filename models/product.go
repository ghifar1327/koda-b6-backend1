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
	ProductID int `json:"product_id"`
	VariantID int `json:"variant_id"`
	SizeID    int `json:"size_id"`
	MethodID  int `json:"method_id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
}

type Order struct {
	Id            int        `json:"id"`
	UserID        int        `json:"user_id"`
	Total         int        `json:"total"`
	Status        string     `json:"status"`
	Address       string     `json:"address"`
	PaymentMethod string     `json:"payment_method"`
	CreatedAt     time.Time  `json:"created_at"`
	Items         []CartItem `json:"items"`
}

// ===================================================================== methode

func (s Size) Render(ids []int) []string {
	var result []string

	for _, id := range ids {
		for _, size := range Sizes {
			if size.Id == id {
				result = append(result, size.Name)
				break
			}
		}
	}

	return result
}

func (v Variant) Render(idv []int) []string {
	var result []string
	for _, id := range idv {
		for _, variant := range Variants {
			if variant.Id == id {
				result = append(result, variant.Name)
				break
			}
		}
	}
	return result
}

func (m Method) Render(idm []int) []string {
	var result []string
	for _, id := range idm {
		for _, method := range Methods {
			if method.Id == id {
				result = append(result, method.Name)
			}
		}
	}
	return result
}

type Renders interface {
	Render() []string
}
