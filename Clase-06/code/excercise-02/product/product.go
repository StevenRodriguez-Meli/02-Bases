package product

type Product interface {
	Price() float64
}

func CreateProduct(productType string, price float64) (Product, string) {
	var product Product
	var errMsg string
	switch productType {
	case "Small":
		product = Small{PriceOfProduct: price}
	case "Medium":
		product = Medium{PriceOfProduct: price}
	case "Large":
		product = Large{PriceOfProduct: price}
	default:
		errMsg = "Unknown product type"
	}
	return product, errMsg
}
