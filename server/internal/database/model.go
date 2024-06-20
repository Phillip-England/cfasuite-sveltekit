package database

type DataModel interface {
	Table() string
	Create() string
	Read() string
	Update() string
	Delete() string
}
