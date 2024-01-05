package product

type Small struct {
	PriceOfProduct float64
}

func (s Small) Price() float64 {
	return s.PriceOfProduct
}
