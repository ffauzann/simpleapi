package response

type (
	Merchant struct {
		ID   int    `json:"id" example:"1"`
		Name string `json:"name" example:"merchant 1"`
	}

	MerchantWithOutlets struct {
		ID      int               `json:"id" example:"1"`
		Name    string            `json:"name" example:"merchant 1"`
		Outlets []OutletWithGross `json:"outlets"`
	}

	MerchantGross struct {
		Date     string   `json:"date" example:"2021-11-01"`
		Merchant Merchant `json:"merchant"`
		Gross    int      `json:"gross" example:"1000"`
	}
)
