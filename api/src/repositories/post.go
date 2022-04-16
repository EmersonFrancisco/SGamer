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

// SeachId traz uma unica publicação do banco de dados de acordo com ID
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

// SearchFeed traz as publicações do feed do usuario, com suas e as dos que segue
func (repository posts) SearchFeed(id uint64) ([]models.Post, error) {
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

// SearchUser traz as publicações de um usuário especifico
func (repository posts) SearchByUser(id uint64) ([]models.Post, error) {
	lines, erro := repository.db.Query(`
		select p.*, u.nick from
		post p inner join user u
		on u.id = p.authorid where p.authorid = ?`, id,
	)
	if erro != nil {
		return nil, erro
	}
	var posts []models.Post
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

// UpdatePost efetua a alteração dos dados da publicação
func (repository posts) Update(postAlter models.Post, id uint64) error {
	statement, erro := repository.db.Prepare(
		"update post set title = ?, content = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(postAlter.Title, postAlter.Content, id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository posts) Delete(id uint64) error {
	statement, erro := repository.db.Prepare(
		"delete from post where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository posts) Like(id uint64) error {
	statement, erro := repository.db.Prepare(
		"update post set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository posts) Unlike(id uint64) error {
	statement, erro := repository.db.Prepare(`
			UPDATE post SET likes = 
			CASE 
				WHEN likes > 0 
					THEN likes - 1 
				ELSE 0 
			END WHERE id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(id)
	if erro != nil {
		return erro
	}
	return nil
}
