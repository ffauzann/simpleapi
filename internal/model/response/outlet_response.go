package response

type (
	// Reusable outlet's response
	Outlet struct {
		ID   int    `json:"id" example:"1"`
		Name string `json:"name" example:"outlet 1"`
	}

	OutletWithGross struct {
		ID    int    `json:"id" example:"1"`
		Name  string `json:"name" example:"outlet 1"`
		Gross int    `json:"gross" example:"1000"`
	}

	OutletGross struct {
		Date     string              `json:"date" example:"2021-11-01"`
		Merchant MerchantWithOutlets `json:"merchant"`
	}
)
