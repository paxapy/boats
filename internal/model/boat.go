package model

type Boat struct {
    Id int8
    Title, Slug, Type string
    Length, Width, Height, Thickness, Weight float32
    Capacity int8
}

func (m *Model) Boats() ([]*Boat, error) {
    return m.SelectBoats()
}
