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

// SearchAllUsers busca todos os usuários que atendem o filtro no Banco
func (repository users) SearchFilter(filter string) ([]models.User, error) {
	filter = fmt.Sprintf("%%%s%%", filter) // %filter%
	lines, erro := repository.db.Query(
		"select id, username, nick, email, createDate from user where username LIKE ? or nick LIKE ?", filter, filter)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if erro := lines.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.CreateDate); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}

// SearchUser busca o usuário no banco com id informado
func (repository users) Search(id uint64) (models.User, error) {
	var user models.User
	line, erro := repository.db.Query(
		"select id,username,nick,email,createdate from user where id = ?", id)
	if erro != nil {
		return user, erro
	}
	defer line.Close()
	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.CreateDate); erro != nil {
			return user, erro
		}
	}
	if user.ID == 0 {
		erro = errors.New(fmt.Sprintf("Nenhum user encontrado no banco com id %d", id))
		return user, erro
	}

	return user, nil
}

// UpdateUser verifica a existencia de usuário com id no banco
// e atualiza dados do usuário no banco com id informado
func (repository users) Update(id uint64, userReq models.User) error {
	userDB, erro := repository.Search(id)
	if erro != nil {
		return erro
	}
	if userReq.Username != "" {
		userDB.Username = userReq.Username
	}
	if userReq.Nick != "" {
		userDB.Nick = userReq.Nick
	}
	if userReq.Email != "" {
		userDB.Email = userReq.Email
	}

	statement, erro := repository.db.Prepare(
		"update user set username = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro := statement.Exec(userDB.Username, userDB.Nick, userDB.Email, id); erro != nil {
		return erro
	}

	return nil
}

// DeleteUser verifica a existencia de usuário com id no banco
// e deleta o usuário do banco com o id informado
func (repository users) Delete(id uint64) error {
	_, erro := repository.Search(id)
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

// SearchEmail busca um esuario pelo e-mail iformado e retorna seu id e senha com hash
func (repository users) SearchEmail(email string) (models.User, error) {
	var user models.User
	line, erro := repository.db.Query(
		"select id, pass from user where email = ?", email)
	if erro != nil {
		return user, erro
	}
	defer line.Close()
	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Pass); erro != nil {
			return user, erro
		}
	}
	return user, nil
}

// Follow permite que um usuário siga outro
func (repository users) Follow(userId, followerUserId uint64) error {
	statement, erro := repository.db.Prepare(
		"insert ignore into follower(user_id, follower_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(userId, followerUserId)
	if erro != nil {
		return erro
	}
	return nil
}

// Unfollow permite que um usuário pare de seguir o outro
func (repository users) Unfollow(userId, followerUserId uint64) error {
	statement, erro := repository.db.Prepare(
		"delete from follower where user_id = ? and follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(userId, followerUserId)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository users) SearchFollowers(id uint64) ([]models.User, error) {
	var users []models.User
	lines, erro := repository.db.Query(`
		select u.id, u.username, u.nick, u.email, u.createdate 
		from user u inner join follower f on u.id = f.follower_id where f.user_id = ?`, id,
	// foi unido a tabela user e follower, e busca os dados do usuario que o user.id(usuario)
	// seja igual ao FOLLOWER_ID(seguidor)
	// mas com uma condições que o usuario seguido seja o ID que desejamos...
	)
	if erro != nil {
		return users, erro
	}
	var user models.User
	defer lines.Close()
	for lines.Next() {
		if erro := lines.Scan(&user.ID, &user.Username, &user.Nick, &user.Email, &user.CreateDate); erro != nil {
			return users, erro
		}
		users = append(users, user)
	}
	if id == 0 {
		erro = errors.New(fmt.Sprintf("Nenhum user encontrado no banco com id %d", id))
		return users, erro
	}

	return users, nil
}
