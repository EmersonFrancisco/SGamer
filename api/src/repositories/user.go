package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewRepositoriesUser cria um repositório de usuários
func NewRepositoryUser(db *sql.DB) *users {
	return &users{db}
}

// Create inseri os dados do usuario no Banco
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into user (username, nick, email, pass) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	insert, erro := statement.Exec(user.Username, user.Nick, user.Email, user.Pass)
	if erro != nil {
		return 0, erro
	}
	IdInsert, erro := insert.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(IdInsert), nil
}

func (repository users) SearchAllUsers() ([]models.User, error) {
	lines, erro := repository.db.Query("select * from user")
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if erro := lines.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.Pass, &user.CreateDate); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) SearchUser(id uint64) (models.User, error) {
	var user models.User
	line, erro := repository.db.Query("select * from user where id = ?", id)
	if erro != nil {
		return user, erro
	}
	defer line.Close()
	if erro := line.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.Pass, &user.CreateDate); erro != nil {
		return user, erro
	}
	if user.ID == 0 {
		erro = errors.New(fmt.Sprintf("Nenhum user encontrado no banco com id %d", id))
		return user, erro
	}

	return user, nil
}

func (repository users) UpdateUser(id uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"update user set username = ?, nick = ?, email = ?, pass = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro := statement.Exec(user.Username, user.Nick, user.Email, user.Pass, id); erro != nil {
		return erro
	}

	return nil
}

func (repository users) DeleteUser(id uint64) error {
	statement, erro := repository.db.Prepare("delete from user where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}
	return nil
}
