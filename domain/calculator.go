package domain

var (
	ProductMaster = map[string]Product{
		"Red": Product{
			Name:  "Red",
			Price: 50,
		},
		"Green": Product{
			Name:  "Green",
			Price: 40,
		},
		"Blue": Product{
			Name:  "Blue",
			Price: 30,
		},
		"Yellow": Product{
			Name:  "Yellow",
			Price: 50,
		},
		"Pink": Product{
			Name:  "Pink",
			Price: 80,
		},
		"Purple": Product{
			Name:  "Purple",
			Price: 90,
		},
		"Orange": Product{
			Name:  "Orange",
			Price: 120,
		},
	}

	DoublePromotionMaster = map[string]bool{"Orange": true, "Pink": true, "Green": true}
	PromotionsPerProduct  = []PromotionsPerProductPrice{PromotionDouble}
	PromotionsTotal       = []PromotionTotalPrice{PromotionMember}
)

type Product struct {
	Name  string
	Price float32
}

type CalculatorReq struct {
	CartItemReq []CartItemReq
	IsMember    bool
}

type CalculatorRes struct {
	CartItems     []CartItem
	TotalPrice    float32
	TotalDiscount float32
}

type CartItemReq struct {
	Name  string
	Count int
}

type CartItem struct {
	Count    int
	Product  Product
	Discount float32
	Price    float32
}

type TotalPriceParams struct {
	IsMember bool
}

type PromotionTotalPrice func(params TotalPriceParams, currentPrice float32) (float32, float32)
type PromotionsPerProductPrice func(item CartItem) CartItem

func PromotionMember(params TotalPriceParams, currentPrice float32) (float32, float32) {
	if params.IsMember {
		return CalWithPersen(currentPrice, 10)
	}
	return currentPrice, 0
}

func PromotionDouble(item CartItem) CartItem {
	if DoublePromotionMaster[item.Product.Name] {
		if item.Count >= 2 {
			priceDouble, discountDouble := CalWithPersen(item.Price, 5)
			item.Price = priceDouble
			item.Discount = discountDouble
		}
	}
	return item
}

func CalWithPersen(currentPrice float32, persen float32) (float32, float32) {
	discount := (currentPrice * (persen / 100))
	price := currentPrice - discount
	return price, discount
}
