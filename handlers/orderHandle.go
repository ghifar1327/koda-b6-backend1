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
