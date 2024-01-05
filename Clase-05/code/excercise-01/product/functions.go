package product

func GetByID(id int) Product {
	var emptyProduct Product
	for _, p := range Products {
		if p.Id == id {
			return p
		}
	}
	return emptyProduct
}
