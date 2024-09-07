package repository

import (
	"database/Assignment/http-client/internal/constant"
	"database/Assignment/http-client/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/go-redis/redis"
)

// ErrUserNotFound is returned when a user is not found
var ErrUserNotFound = errors.New("user not found")

// UserRepository performs database operations for users
type UserRepository struct {
	db  *sql.DB
	rdb *redis.Client
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB, rdb *redis.Client) *UserRepository {
	return &UserRepository{
		db:  db,
		rdb: rdb,
	}
}

// LoginUser checks if the user exists in the database and returns it or an error
func (r *UserRepository) LoginUser(userLogin models.UserLogin) (*string, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM User WHERE username =?", userLogin.UserName).Scan(&user.ID, &user.Name, &user.UserName, &user.Password, &user.Address)

	if err != nil || user.Password != userLogin.Password {
		return nil, ErrUserNotFound
	}

	sessionID := "session-" + userLogin.UserName

	err = r.rdb.Set(sessionID, user.ID, constant.REDIS_TTL*time.Second).Err()

	if err != nil {
		return nil, err
	}
	return &sessionID, nil
}

func (r *UserRepository) Ping(sessionKey string) (*string, error) {
	userId, err := r.rdb.Get(sessionKey).Result()
	nowTimestamp := time.Now().Unix()
	var userPingCountKey = userId + "_ping_count"
	if err != nil {
		return nil, fmt.Errorf("not found session")
	}

	pingLock, err := r.rdb.Get(constant.PingLock).Result()

	if err != nil {
		errLock := r.rdb.Set(constant.PingLock, true, 0).Err()

		if errLock != nil {
			return nil, errLock
		}
	}

	if pingLock == "1" {
		return nil, fmt.Errorf("another people is pinging")
	}

	var keyFormat = "user_" + userId + "_*"

	existKeys, err := r.rdb.Keys(keyFormat).Result()

	if len(existKeys) > 2 {
		return nil, fmt.Errorf("you call more than 2 times in 60 seconds")
	}

	// Tạo key với format "user_" + "timestamp"
	userTime := fmt.Sprintf("user_%s_%d", userId, nowTimestamp)

	// get ping time
	pingTime, _ := r.rdb.Get(userPingCountKey).Result()

	if pingTime == "" {
		r.rdb.Set(userPingCountKey, 1, 0)
	} else {
		r.rdb.Incr(userPingCountKey)
	}

	r.rdb.PFAdd("users:ping:unique", userId)
	r.rdb.Set(userTime, 1, 60*time.Second)
	// Simulate processing time
	time.Sleep(1 * time.Second)

	r.rdb.Set(constant.PingLock, false, 0)

	// get new value
	count, _ := r.rdb.Get(userPingCountKey).Result()
	var result = "Total Count: " + count

	return &result, nil
}

func (r *UserRepository) TopPing(sessionKey string) (*string, error) {
	_, err := r.rdb.Get(sessionKey).Result()
	if err != nil {
		return nil, fmt.Errorf("not found session")
	}

	var keyFormat = "*_ping_count"
	re := regexp.MustCompile(`(\d+)_ping_count`)

	existKeys, err := r.rdb.Keys(keyFormat).Result()

	var result = ""

	for i := 0; i < len(existKeys); i++ {
		key := existKeys[i]
		matches := re.FindStringSubmatch(key)
		userID := matches[1]
		var user models.User
		err := r.db.QueryRow("SELECT * FROM User WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.UserName, &user.Password, &user.Address)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		result = result + user.Name + "\n"
	}

	return &result, nil
}

func (r *UserRepository) Count(sessionKey string) (*string, error) {
	_, err := r.rdb.Get(sessionKey).Result()
	if err != nil {
		return nil, fmt.Errorf("not found session")
	}

	count, err := r.rdb.PFCount("users:ping:unique").Result()
	if err != nil {
		// Xử lý lỗi
	}

	result := fmt.Sprintf("Số người gọi API /ping: %d", count)

	return &result, nil
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
		err = rows.Scan(&user.ID, &user.Name, &user.UserName, &user.Password, &user.Address)
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
	err := r.db.QueryRow("SELECT * FROM User WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Address, &user.Name, &user.Password)
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
		err := rows.Scan(&user.ID, &user.Name, &user.Address, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
