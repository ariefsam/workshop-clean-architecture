package repository

import "log"

type mockRepository struct {
	saved map[string]any
}

func NewMock() (repo *mockRepository) {
	repo = &mockRepository{}
	return
}

func (m *mockRepository) Save(id string, tableName string, data any) (err error) {
	if m.saved == nil {
		m.saved = map[string]any{}
	}
	m.saved[tableName] = data
	log.Println("data saved table: ", tableName, "id:", id)
	return
}

func (m *mockRepository) Get(id string, tableName string, data any) (err error) {
	return
}

func (m *mockRepository) IsSaved(tableName string) (isSaved bool) {
	if m.saved == nil {
		return
	}
	if m.saved[tableName] != nil {
		isSaved = true
		return
	}
	return
}
