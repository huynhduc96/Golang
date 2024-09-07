package models

import (
	"social/internal/constant"
	"social/internal/utils"
	"time"
)

const DOBLayout = "2006-01-02"
const saltSize = 16

type User struct {
	Id         int    `json:"id,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	Email      string `json:"email,omitempty"`
	UserName   string `json:"userName,omitempty"`
	Password   string `json:"password,omitempty"`
	DOB        string `json:"dob,omitempty"`
	DOBDate    time.Time
	Salt       []byte
	HashedPass string
}

type UserLogin struct {
	UserName string
	Password string
}

type UserUpdate struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	DOB       string
}

func (u *User) IsMatchPassword(curPass string) bool {
	return utils.IsPassMatch(u.HashedPass, curPass, u.Salt)
}

func (u *User) Update(updateRecord *User) {
	if updateRecord.FirstName != "" {
		u.FirstName = updateRecord.FirstName
	}

	if updateRecord.LastName != "" {
		u.LastName = updateRecord.LastName
	}

	if updateRecord.Password != "" {
		salt := utils.GenRandomSalt(constant.SaltSize)

		hashedPass := utils.HashPassword(u.Password, salt)
		u.Password = hashedPass
	}

	if updateRecord.DOB != "" {
		u.DOB = updateRecord.DOB
	}

}
