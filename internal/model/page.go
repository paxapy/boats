package model

type Page struct {
    Id int
    Title, Slug, Body string
}

func (m *Model) Pages() ([]*Page, error) {
    return m.SelectPages()
}
