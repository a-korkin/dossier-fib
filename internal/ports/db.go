package ports

type Repo interface {
	GetAll() []interface{}
	GetByID(id uint32) interface{}
}
