package types

// Page 分页参数
type Page struct {
	Limit int `validate:"gt=0"`
	Page  int `validate:"gt=0"`
}

func (p *Page) Offset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Page) End() int {
	return (p.GetPage()) * p.GetLimit()
}

func (p *Page) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

// GetLimit 默认每页10条
func (p *Page) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

func (p *Page) Next() *Page {
	p.Page = p.Page + 1
	return p
}
