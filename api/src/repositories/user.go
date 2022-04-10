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

// Create inseri os dados do usuário no Banco
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

// SearchAllUsers busca todos os usuários no Banco
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

// SearchUser busca o usuário no banco com id informado
func (repository users) SearchUser(id uint64) (models.User, error) {
	var user models.User
	line, erro := repository.db.Query("select * from user where id = ?", id)
	if erro != nil {
		return user, erro
	}
	defer line.Close()
	for line.Next() {
		if erro := line.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.Pass, &user.CreateDate); erro != nil {
			return user, erro
		}
	}
	if user.ID == 0 {
		erro = errors.New(fmt.Sprintf("Nenhum user encontrado no banco com id %d", id))
		return user, erro
	}

	return user, nil
}

// UpdateUser atualiza dados do usuário no banco com id informado
func (repository users) UpdateUser(id uint64, userReq models.User) error {
	userDB, erro := repository.SearchUser(id)
	if erro != nil {
		return erro
	}
	switch {
	case userReq.Username != "":
		userDB.Username = userReq.Username
	case userReq.Nick != "":
		userDB.Nick = userReq.Nick
	case userReq.Email != "":
		userDB.Email = userReq.Email
	case userReq.Pass != "":
		userDB.Pass = userReq.Pass
	}
	statement, erro := repository.db.Prepare(
		"update user set username = ?, nick = ?, email = ?, pass = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro := statement.Exec(userDB.Username, userDB.Nick, userDB.Email, userDB.Pass, id); erro != nil {
		return erro
	}

	return nil
}

// DeleteUser deleta o usuário do banco com o id informado
func (repository users) DeleteUser(id uint64) error {
	_, erro := repository.SearchUser(id)
	if erro != nil {
		return erro
	}
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
