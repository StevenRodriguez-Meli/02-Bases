package product

type Medium struct {
	PriceOfProduct float64
}

func (m Medium) Price() float64 {
	return m.PriceOfProduct + (0.03 * m.PriceOfProduct)
}
