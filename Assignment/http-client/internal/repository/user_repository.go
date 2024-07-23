package repository

import (
	"database/Assignment/http-client/internal/models"
	"database/sql"
	"errors"
)

// ErrUserNotFound is returned when a user is not found
var ErrUserNotFound = errors.New("user not found")

// UserRepository performs database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetAllUsers returns all users
func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetUserByID returns a user by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM User WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *models.User) error {
	result, err := r.db.Exec("INSERT INTO User(name, address) VALUES(?, ?)", user.Name, user.Address)
	if err != nil {
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(lastID)
	return nil
}

// UpdateUser updates an existing user
func (r *UserRepository) UpdateUser(user *models.User) error {
	result, err := r.db.Exec("UPDATE User SET name = ?, address = ? WHERE id = ?", user.Name, user.Address, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(id int) error {
	result, err := r.db.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) SearchUsersByAddress(address string, page, pageSize int) ([]models.User, error) {
	// Query the database for users with the provided address
	// ... (Implement the database query)
	value := "%" + address + "%"
	offset := (page - 1) * pageSize
	rows, err := r.db.Query("SELECT * FROM User WHERE address LIKE ? LIMIT ? OFFSET ?", value, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
