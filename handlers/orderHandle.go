package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

type ProductResponse struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	BasePrice int      `json:"base_price"`
	Stock     int      `json:"stock"`
	Variants  []string `json:"variants"`
	Sizes     []string `json:"sizes"`
	Methods   []string `json:"methods"`
}

var Cart []models.CartItem

func GetProducts(ctx *gin.Context) {
	defer mu.Unlock()
	mu.Lock()

	var result []ProductResponse

	for _, product := range models.Products {
		result = append(result, ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			BasePrice: product.BasePrice,
			Stock:     product.Stock,
			Variants:  models.Variant{}.Render(product.Variants),
			Sizes:     models.Size{}.Render(product.Sizes),
			Methods:   models.Method{}.Render(product.Methods),
		})
	}
	ctx.JSON(200, Response{true, "List of Product", result})
}

func AddChart(ctx *gin.Context) {
	defer mu.Unlock()
	mu.Lock()
	
	var input models.CartItem
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(404, gin.H{"error": "invalid request body"})
		return
	}

	// =========================================================== search product
	var product *models.Product
	for i := range models.Products {
		if models.Products[i].Id == input.ProductID {
			product = &models.Products[i]
			break
		}
	}

	if product == nil {
		ctx.JSON(404, gin.H{"error": "Product not found"})
	}

	// =========================================================== add price
	var variantPrice, sizePrice, methodPrice int

	for _, v := range models.Variants {
		if v.Id == input.VariantID {
			variantPrice = v.AddPrice
			break
		}
	}

	for _, s := range models.Sizes {
		if s.Id == input.SizeID {
			sizePrice = s.AddPrice
			break
		}
	}
	for _, m := range models.Methods {
		if m.Id == input.MethodID {
			methodPrice = m.AddPrice
		}
	}
	totalPrice := (variantPrice + sizePrice + methodPrice + product.BasePrice) * input.Qty
	input.Id = len(Cart) + 1
	input.Price = totalPrice

	Cart = append(Cart, input)
	ctx.JSON(200, Response{true, "Product added successfuly", Cart})
}
