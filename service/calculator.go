package service

import "simple-cart/domain"

type Calculator interface {
	Execute(req domain.CalculatorReq) domain.CalculatorRes
}

type calculator struct {
	promotionsPerProduct []domain.PromotionsPerProductPrice
	promotionsTotal      []domain.PromotionTotalPrice
}

func NewCalculator(promotionsPerProduct []domain.PromotionsPerProductPrice, promotionsTotal []domain.PromotionTotalPrice) Calculator {
	return &calculator{
		promotionsPerProduct: promotionsPerProduct,
		promotionsTotal:      promotionsTotal,
	}
}

func (c *calculator) Execute(req domain.CalculatorReq) domain.CalculatorRes {

	carts := c.FindPrice(req.CartItemReq, c.promotionsPerProduct)

	totalPriceParams := domain.TotalPriceParams{
		IsMember: req.IsMember,
	}
	totalPrice, totalDiscount := c.TotalPrice(carts, totalPriceParams, c.promotionsTotal)

	result := domain.CalculatorRes{
		CartItems:     carts,
		TotalPrice:    totalPrice,
		TotalDiscount: totalDiscount,
	}

	return result
}

func (c *calculator) FindPrice(carts []domain.CartItemReq, promotions []domain.PromotionsPerProductPrice) []domain.CartItem {
	result := []domain.CartItem{}
	for _, item := range carts {

		product := domain.ProductMaster[item.Name]
		if product.Name == item.Name {
			price := float32(item.Count) * product.Price
			data := domain.CartItem{
				Product: product,
				Price:   price,
				Count:   item.Count,
			}

			for _, promotion := range promotions {
				data = promotion(data)
			}

			result = append(result, data)
		}
	}
	return result
}

func (c *calculator) TotalPrice(carts []domain.CartItem, params domain.TotalPriceParams, promotions []domain.PromotionTotalPrice) (float32, float32) {
	var price, discount float32
	for _, item := range carts {
		price = price + item.Price
		discount = discount + item.Discount
	}

	var newPrice, newDiscount float32
	for _, promotion := range promotions {
		newPrice, newDiscount = promotion(params, price)
	}
	return newPrice, discount + newDiscount
}
