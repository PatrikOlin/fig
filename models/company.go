package models

type Company struct {
	CompanyName     string `json:"CompanyName"`
	OrgNum          string `json:"OrgNum"`
	VatCode         string `json:"VatCode"`
	BeneficialOwner string `json:"BeneficialOwner"`
}
