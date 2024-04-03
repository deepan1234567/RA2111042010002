func main() {
	router := gin.Default()
  
	// API to retrieve top N products within a price range
	router.GET("/test/companies/:companyname/categories/:categoryname/products/top-n&minPrice=:minPrice&maxPrice=:maxPrice", getTopProducts)
  
	// Start the server
	router.Run(":8080")
  }
  
  func getTopProducts(c *gin.Context) {
	companyName := c.Param("companyname")
	categoryName := c.Param("categoryname")
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for parameter 'n'"})
	  return
	}
  
	minPrice, err := strconv.ParseFloat(c.Param("minPrice"), 64)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for parameter 'minPrice'"})
	  return
	}
  
	maxPrice, err := strconv.ParseFloat(c.Param("maxPrice"), 64)
	if err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for parameter 'maxPrice'"})
	  return
	}
  
	// Validate company and category based on your accepted values lists
	if !isValidCompany(companyName) || !isValidCategory(categoryName) {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company or category"})
	  return
	}
  
	// ... Retrieve product data from the specified company (see next step)
	// ... Filter products based on price range
	// ... Select top N products
  
	c.JSON(http.StatusOK, products)
  }
  
  