func getProductData(companyName string, categoryName string) ([]Product, error) {
	// Use the appropriate client library to call the e-commerce company's API
	// Parse the response and extract relevant product data into Product structs
  
	// Example (replace with actual API calls)
	if companyName == "AMZ" {
	  products, err := company1.GetProducts(categoryName)
	  if err != nil {
		return nil, err
	  }
	  // ... convert products to Product structs
	  return convertedProducts, nil
	}
	// ... similar logic for other companies
	return nil, fmt.Errorf("Unsupported company: %s", companyName)
  }
  
  func filterProductsByPrice(products []Product, minPrice, maxPrice float64) []Product {
	var filteredProducts []Product
	for _, product := range products {
	  if product.Price >= minPrice && product.Price <= maxPrice {
		filteredProducts = append(filteredProducts, product)
	  }
	}
	return filteredProducts
  }
  
  