package models


// ======================================================================================== dummy varrian
var Variants = []Variant{
	// ================== drinks
	{1, "hot", 0},
	{2, "ice", 0},
	// ================== foods
	{3, "spicy", 3000},
	{4, "tasty", 0},
}

// ========================================================================================= dummy size
var Sizes = []Size{
	{1,"reguar",0},
	{2,"medium",3000},
	{3,"lerge",5000},

}

// ==========================================================================================  methode
var Methods = []Method{
	{1,"dine in", 0},
	{2,"delivery door", 10000},
	{3,"pick up", 2000},

}

// ======================================================================================== dummy product
var Products = []Product{
	{
		Id:        1,
		Name:      "Cappuccino",
		BasePrice: 20000,
		Stock:     50,
		Variants: []int{1,2,},
		Sizes: []int{1,2,3},
		Methods: []int{1,2,3,},
	},
	{
		Id:        2,
		Name:      "Chiken Burger",
		BasePrice: 22000,
		Stock:     40,
		Variants: []int{3,4,},
		Sizes: []int{1,2,3,},
		Methods: []int{1},
	},
}

// // ======================================================================================================== dummy cart

// var CartItems = []CartItem{
// 	{
// 		Id:        1,
// 		UserID:    1,
// 		ProductID: 1, // Cappuccino
// 		VariantID: 3, // Extra Shot
// 		SizeID:    3, // Large
// 		MethodID:  1, // Dine In
// 		Qty:       2,
// 		Price:     32000, // 20000 + 7000 + 5000
// 	},
// 	{
// 		Id:        2,
// 		UserID:    1,
// 		ProductID: 2,
// 		VariantID: 5,
// 		SizeID:    5,
// 		MethodID:  2,
// 		Qty:       1,
// 		Price:     35000,
// 	},
// }

// // =================================================================================================================  dummy orders

// var Orders = []Order{
// 	{
// 		Id:            1,
// 		UserID:        1,
// 		Total:         99000,
// 		Status:        "completed",
// 		Address:       "Jl. Sudirman No.10",
// 		PaymentMethod: "transfer",
// 		CreatedAt:     time.Now(),
// 		Items: []OrderItem{
// 			{
// 				Id:        1,
// 				OrderID:   1,
// 				ProductID: 1,
// 				Qty:       2,
// 				Price:     32000,
// 			},
// 			{
// 				Id:        2,
// 				OrderID:   1,
// 				ProductID: 2,
// 				Qty:       1,
// 				Price:     35000,
// 			},
// 		},
// 	},
// }
