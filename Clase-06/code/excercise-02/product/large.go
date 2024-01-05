package product

type Large struct {
	PriceOfProduct float64
}

func (l Large) Price() float64 {
	return l.PriceOfProduct + (0.06 * l.PriceOfProduct) + 2500
}
