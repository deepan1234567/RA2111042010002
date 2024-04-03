func selectTopNProducts(products []Product, n int) []Product {
	if n > len(products) {
	  n = len(products)
	}
	return products[:n]
  }
  