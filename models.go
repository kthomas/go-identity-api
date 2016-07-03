package identity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Model struct {
	ID        	uint 		`gorm:"primary_key",json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	Errors		Errors		`gorm:"-",json:"-"`
}

type Error struct {
	Message     string      `json:"message"`
	Status      int         `json:"status"`
}
type Errors []Error

func (m *User) Save() (bool) {
	db := DatabaseConnection()
	if db.NewRecord(m) {
		result := db.Create(&m)
		rowsAffected := result.RowsAffected
		dbErrors := result.GetErrors()

		errors := make(Errors, len(dbErrors))
		for i := range dbErrors {
			errors[i] = Error{ Message: dbErrors[i].Error() }
		}
		m.Errors = errors

		if !db.NewRecord(m) {
			return rowsAffected > 0
		}
	} else {
		result := db.Save(&m)
		rowsAffected := result.RowsAffected
		dbErrors := result.GetErrors()

		errors := make(Errors, len(dbErrors))
		for i := range dbErrors {
			errors[i] = Error{ Message: dbErrors[i].Error() }
		}
		m.Errors = errors

		return rowsAffected > 0
	}
	return false
}

func (m *Token) Create() (bool) {
	db := DatabaseConnection()
	if db.NewRecord(m) {
		result := db.Create(&m)
		rowsAffected := result.RowsAffected
		//dbErrors := result.GetErrors()
		if !db.NewRecord(m) {
			return rowsAffected > 0
		}
	}
	return false
}

type Token struct {
	Model
	UserID    	uint      	`json:"user_id"`
	Token     	string    	`json:"token"`
}
type Tokens []Token

type User struct {
	Model
	Email     	string    `gorm:"unique_index",json:"email"`
	Name      	string    `json:"name"`
	Password  	string    `json:"password"`
	Tokens    	Tokens    `json:"-"`
}
type Users []User

func (u *User) Authenticate(password string) (bool) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
