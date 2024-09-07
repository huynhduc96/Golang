package repository

import (
	"database/sql"
	"social/internal/models"
)

type NewsFeedRepository struct {
	DB *sql.DB
}

func CreateNewsFeedRepository(db *sql.DB) *NewsFeedRepository {
	return &NewsFeedRepository{
		DB: db,
	}
}

func (nfRepository *NewsFeedRepository) GetUserRecordByUserId(userId int) (models.User, error) {
	stmtOut, err := nfRepository.DB.Prepare("SELECT id, first_name, last_name, user_name, email, salt, hashed_password FROM user WHERE id = ?")
	defer stmtOut.Close()

	var user models.User

	if err != nil {
		return user, err
	}

	err = stmtOut.QueryRow(userId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.UserName,
		&user.Email,
		&user.Salt,
		&user.HashedPass,
	)

	return user, err
}

func (nfRepository *NewsFeedRepository) GenAllNewsFeedPost(id int) ([]models.Post, error) {
	var posts []models.Post

	stmtOut, err := nfRepository.DB.Prepare("SELECT id, content_text, IFNULL(content_image_path, '') AS content_image_path, created_at FROM `user_user` u_u LEFT JOIN `post` p ON u_u.fk_follower_id = p.fk_user_id WHERE u_u.fk_user_id = ? AND visible = 1 ORDER BY created_at DESC")
	defer stmtOut.Close()

	if err != nil {
		return posts, err
	}

	rows, err := stmtOut.Query(id)
	defer rows.Close()

	if err != nil {
		return posts, err
	}

	for rows.Next() {
		var p models.Post

		err := rows.Scan(&p.Id, &p.ContentText, &p.ContentImagePath, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	err = rows.Err()
	if err != nil {
		return posts, err
	}

	return posts, nil
}
