package core

type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CookieIds string `json:"cookie_ids"`
	Prev      string `json:"prev"`
	Next      string `json:"next"`
	Head      bool   `json:"head"`
}

func (pa *Page) Save() {
	db.Create(pa)
}

func (pa *Page) Update() {
	db.Model(pa).Update(pa)

}
func (pa *Page) Delete() {
	db.Delete(pa)
}

func GetPages() []Page {
	var pages []Page
	db.Where("head = ?", true).Find(&pages)
	return pages
}

func (pa *Page) next() Page {
	var page Page
	db.Where("id = ?", pa.Next).First(&page)
	return page
}

func GetPageLink(head string) []Page {
	var pages []Page
	var page Page
	db.Where("id = ?", head).First(&page)
	if page.Head {
		pages = append(pages, page)
		for {
			if page.Next != "" || page.Head {
				page = page.next()
				if page.Id != "" {
					pages = append(pages, page)
				}
			} else {
				break
			}
		}
	}

	return pages
}
