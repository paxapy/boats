package model

type db interface {
    SelectBoats() ([]*Boat, error)
    SelectPages() ([]*Page, error)
}

type Model struct {
    db
}

func New(db db) *Model {
    return &Model{
        db: db,
    }
}
