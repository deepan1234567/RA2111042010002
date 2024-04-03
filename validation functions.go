func isValidCompany(companyName string) bool {
	validCompanies := []string{"AMZ", "FLP", "SNP", "MYN", "AZO"}
	for _, validCompany := range validCompanies {
	  if companyName == validCompany {
		return true
	  }
  
  