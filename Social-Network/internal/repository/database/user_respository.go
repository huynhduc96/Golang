package repository

import (
	"database/sql"
	"log"
	"social/internal/models"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) CreateNewUser(newUser *models.User, bOBDate time.Time, salt []byte, hashedPass string) (int64, error) {
	stmtIn, err := u.DB.Prepare("INSERT INTO user (first_name, last_name, user_name, email, salt, hashed_password, dob) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmtIn.Close()

	result, err := stmtIn.Exec(&newUser.FirstName, &newUser.LastName, &newUser.UserName, &newUser.Email, salt, hashedPass, bOBDate)
	if err != nil {
		return -1, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil

}

func (u *UserRepository) UpdateUser(id int, updateUserRecord *models.UserUpdate, salt []byte, hashedPass string, bOBDate time.Time) error {
	stmtIn, err := u.DB.Prepare("UPDATE `user` SET first_name = ?, last_name = ?, email = ?, salt = ?, hashed_password = ?, dob = ? WHERE id = ? ")
	if err != nil {
		return err
	}
	defer stmtIn.Close()

	_, err = stmtIn.Exec(
		&updateUserRecord.FirstName,
		&updateUserRecord.LastName,
		&updateUserRecord.Email,
		salt, hashedPass, bOBDate,
		id,
	)

	return nil

}

func (u *UserRepository) GetUserRecordByUsername(username string) (models.User, error) {
	stmtOut, err := u.DB.Prepare("SELECT id, first_name, last_name, user_name, email, salt, hashed_password FROM user WHERE user_name = ?")
	defer stmtOut.Close()

	var user models.User

	if err != nil {
		return user, err
	}

	err = stmtOut.QueryRow(username).Scan(
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

func (u *UserRepository) GetUserRecordByUserId(userId int) (models.User, error) {
	stmtOut, err := u.DB.Prepare("SELECT id, first_name, last_name, user_name, email, salt, hashed_password FROM user WHERE id = ?")
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

func (u *UserRepository) ViewFollowers(id int) ([]models.User, error) {
	var followers []models.User

	stmtOut, err := u.DB.Prepare("SELECT first_name, last_name, email, user_name FROM `user_user` u_u LEFT JOIN `user` u ON(u_u.fk_follower_id = u.id) WHERE u_u.fk_user_id = ?")
	defer stmtOut.Close()

	if err != nil {
		return followers, err
	}

	rows, err := stmtOut.Query(id)
	defer rows.Close()

	if err != nil {
		return followers, err
	}

	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.FirstName, &u.LastName, &u.Email, &u.UserName)
		if err != nil {
			return nil, err
		}

		followers = append(followers, u)
	}

	err = rows.Err()
	if err != nil {
		return followers, err
	}

	return followers, nil
}

func (u *UserRepository) FollowUser(id int, friendId int) error {
	stmtIn, err := u.DB.Prepare("INSERT INTO `user_user` (fk_user_id, fk_follower_id) VALUES (?, ?)")
	defer stmtIn.Close()

	if err != nil {
		return err
	}

	_, err = stmtIn.Exec(id, friendId)
	return err
}

func (u *UserRepository) UnFollowHandler(id int, friendId int) error {
	stmtIn, err := u.DB.Prepare("DELETE FROM `user_user` WHERE fk_user_id=? AND fk_follower_id=?")
	defer stmtIn.Close()

	if err != nil {
		return err
	}

	_, err = stmtIn.Exec(id, friendId)
	return err
}

func (u *UserRepository) ViewFriendPost(id int) ([]models.Post, error) {
	var posts []models.Post

	stmtOut, err := u.DB.Prepare("SELECT id, content_text, IFNULL(content_image_path, '') AS content_image_path, created_at FROM post WHERE fk_user_id = ? AND visible = 1")
	defer stmtOut.Close()

	if err != nil {
		log.Printf("err: %+v\n", err)
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

func (u *UserRepository) CreatePost(userId int, createPost *models.CreatePost) (int64, error) {
	stmtIn, err := u.DB.Prepare("INSERT INTO post (fk_user_id, content_text) VALUES (?, ?)")
	defer stmtIn.Close()

	if err != nil {
		return -1, err
	}

	result, err := stmtIn.Exec(
		userId,
		&createPost.ContentText,
	)

	if err != nil {
		return -1, err
	}

	newPostId, err := result.LastInsertId()
	return newPostId, err
}

func (u *UserRepository) GetPostById(postId int) (models.Post, error) {
	stmtOut, err := u.DB.Prepare("SELECT id, fk_user_id, content_text, IFNULL(content_image_path, '') AS content_image_path, created_at FROM `post` WHERE id = ? AND visible = 1")
	defer stmtOut.Close()

	var p models.Post

	if err != nil {
		return p, err
	}

	err = stmtOut.QueryRow(postId).Scan(
		&p.Id, &p.UserId, &p.ContentText, &p.ContentImagePath,
		&p.CreatedAt,
	)

	return p, err
}

func (u *UserRepository) UpdatePost(postId int, createPost *models.CreatePost) error {
	stmtIn, err := u.DB.Prepare("UPDATE `post` SET content_text = ? WHERE id = ?")
	defer stmtIn.Close()

	if err != nil {
		return err
	}

	_, err = stmtIn.Exec(
		&createPost.ContentText,
		&postId,
	)

	return err
}

func (u *UserRepository) DeletePost(postId int) error {
	stmtIn, err := u.DB.Prepare("UPDATE `post` SET visible = 0 WHERE id = ?")
	defer stmtIn.Close()

	if err != nil {
		return err
	}

	_, err = stmtIn.Exec(
		&postId,
	)

	return err

}

func (u *UserRepository) CommentPost(postId int, userId int, content string) (int64, error) {
	stmtIn, err := u.DB.Prepare("INSERT INTO comment (fk_post_id, fk_user_id, content) VALUES (?, ?, ?)")
	defer stmtIn.Close()

	if err != nil {
		return -1, err
	}

	result, err := stmtIn.Exec(
		&postId,
		&userId,
		&content,
	)

	if err != nil {
		return -1, err
	}

	newCmtId, err := result.LastInsertId()
	return newCmtId, err
}

func (u *UserRepository) LikePost(postId int, userId int) error {

	stmtFind, err := u.DB.Prepare("SELECT fk_post_id, fk_user_id FROM `like` WHERE fk_post_id = ? AND fk_user_id = ?")
	defer stmtFind.Close()
	if err != nil {
		return err
	}

	var like models.Like

	err = stmtFind.QueryRow(
		&postId,
		&userId,
	).Scan(&like.PostId, &like.UserId)

	log.Printf("like: %+v\n", like)
	if err != nil || like.PostId == 0 {
		stmtIn, err := u.DB.Prepare("INSERT INTO `like` (fk_post_id, fk_user_id) VALUES (?, ?)")
		defer stmtIn.Close()

		if err != nil {
			return err
		}

		_, err = stmtIn.Exec(
			&postId,
			&userId,
		)

		if err != nil {
			return err
		}
	} else {
		stmtDel, err := u.DB.Prepare("DELETE FROM `like` WHERE fk_user_id =? AND fk_post_id =?")
		defer stmtDel.Close()

		if err != nil {
			return err
		}

		_, err = stmtDel.Exec(

			userId,
			postId,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UserRepository) ViewCommentPost(postId int) ([]models.Like, error) {
	var likes []models.Like

	stmtOut, err := u.DB.Prepare("SELECT fk_post_id, fk_user_id, content, created_at FROM comment WHERE fk_post_id = ?")
	defer stmtOut.Close()

	if err != nil {
		return likes, err
	}

	rows, err := stmtOut.Query(postId)
	defer rows.Close()

	if err != nil {
		return likes, err
	}

	for rows.Next() {
		var like models.Like

		err := rows.Scan(&like.PostId, &like.UserId)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	err = rows.Err()
	if err != nil {
		return likes, err
	}

	return likes, nil
}

func (u *UserRepository) ViewLikePost(postId int) ([]models.Comment, error) {
	var comments []models.Comment

	stmtOut, err := u.DB.Prepare("SELECT  fk_post_id, fk_user_id FROM like WHERE fk_post_id = ?")
	defer stmtOut.Close()

	if err != nil {
		return comments, err
	}

	rows, err := stmtOut.Query(postId)
	defer rows.Close()

	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var cmt models.Comment

		err := rows.Scan(&cmt.Id, &cmt.PostId, &cmt.UserId, &cmt.ContentText, &cmt.CreatedAt)
		if err != nil {
			return nil, err
		}

		comments = append(comments, cmt)
	}

	err = rows.Err()
	if err != nil {
		return comments, err
	}

	return comments, nil
}
