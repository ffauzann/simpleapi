package entity

type Pagination struct {
	Page           int    `json:"page" query:"page" validate:"required,number,gte=1"`
	Limit          int    `json:"limit" query:"limit" validate:"required,number,gte=1,lte=100"`
	Offset         int    `json:"-"`
	Total          int    `json:"total" example:"26"`
	From           int    `json:"from" example:"1"`
	To             int    `json:"to" example:"10"`
	LastPage       int    `json:"last_page" example:"3"`
	CurrentPageURL string `json:"current_page_url"`
	NextPageURL    string `json:"next_page_url"`
	PrevPageURL    string `json:"prev_page_url"`
}

// Calc calculates offset.
func (p *Pagination) Calc() {
	if p.Page > 1 {
		p.Offset = (p.Limit * (p.Page - 1))
	}
}
