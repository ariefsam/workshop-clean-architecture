package repository

import "database/sql"

type mysqluserRepository struct {
	db *sql.DB
}

func NewMySQLUser(dsn string) (repo *mysqluserRepository, err error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	repo = &mysqluserRepository{
		db: db,
	}
	return
}

type User struct {
	ID      string
	Name    string
	Address string
}

func (m *mysqluserRepository) Save(id string, data User) (err error) {

	query := "INSERT INTO user (id, name, address) VALUES (?, ?, ?)"
	_, err = m.db.Exec(query, id, data.Name, data.Address)

	return
}

func (m *mysqluserRepository) Get(id string) (user User, err error) {
	query := "SELECT id, name, address FROM user WHERE id = ?"
	err = m.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Address)

	return
}
