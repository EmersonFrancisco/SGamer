package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

// NewRepositoriesUser cria um repositório de publicação
func NewRepositoryPost(db *sql.DB) *posts {
	return &posts{db}
}

// Create inseri os dados da publicação no Banco
func (repository posts) Create(post models.Post) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into post (title, content, authorid) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	insert, erro := statement.Exec(post.Title, post.Content, post.AuthorId)
	if erro != nil {
		return 0, erro
	}
	IdInsert, erro := insert.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(IdInsert), nil
}

// SeachId tras uma unica publicação do banco de dados de acordo com ID
func (repository posts) SearchId(id uint64) (models.Post, error) {
	var post models.Post
	line, erro := repository.db.Query(`
		select p.*, u.nick from
		post p inner join user u
		on u.id = p.authorid where p.id = ?`, id,
	)
	if erro != nil {
		return post, erro
	}
	if line.Next() {
		if erro = line.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.DatePost,
			&post.AuthorNick); erro != nil {
			return post, erro
		}
	}
	return post, nil
}

func (repository posts) SearchAll(id uint64) ([]models.Post, error) {
	var posts []models.Post
	lines, erro := repository.db.Query(`
		select distinct p.*, u.nick from
		post p inner join user u
		on u.id = p.authorid 
		inner join follower f 
		on f.user_id = p.authorid
		where u.id = ? or f.follower_id = ? ORDER BY datepost desc`, id, id,
	)
	if erro != nil {
		return posts, erro
	}
	for lines.Next() {
		var post models.Post
		if erro = lines.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.DatePost,
			&post.AuthorNick); erro != nil {
			return nil, erro
		}
		posts = append(posts, post)
	}
	return posts, nil
}
