package repository

type mongoRepository struct{}

func NewMongo() (repo *mongoRepository) {
	repo = &mongoRepository{}
	return
}

func (m *mongoRepository) Save(id string, tableName string, data any) (err error) {
	return
}

func (m *mongoRepository) Get(id string, tableName string, data any) (err error) {
	return
}
