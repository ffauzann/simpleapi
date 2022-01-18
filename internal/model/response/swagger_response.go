package response

type (
	ExampleLogin struct {
		Meta MetaPagination `json:"meta"`
		Data Login          `json:"data"`
	}

	ExampleMerchantGross struct {
		Meta MetaPagination `json:"meta"`
		Data MerchantGross  `json:"data"`
	}

	ExampleOutletGross struct {
		Meta MetaPagination `json:"meta"`
		Data OutletGross    `json:"data"`
	}
)
