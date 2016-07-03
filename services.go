package identity

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

func CreateUser(name string, email string, password string) (*User) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		user := User{
			Name: name,
			Email: email,
			Password: string(hashedPassword),
		}
		if user.Save() {
			return &user
		}
	}
	return nil
}

func AuthenticateUser(email string, password string) (*User) {
	var user User
	db := DatabaseConnection()
	db.Where(&User{ Email: email }).First(&user)
	if &user != nil && user.ID != 0 {
		if user.Authenticate(password) {
			return &user
		}
	}
	return nil
}

func CreateToken(user *User) (*Token) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"name": user.Name,
		"email": user.Email,
	})
	tokenString, _ := jwtToken.SignedString(config.JwtSharedSecret)

	token := Token{
		UserID: jwtToken.Claims.(jwt.MapClaims)["user_id"].(uint),
		Token: tokenString,
	}

	if token.Create() {
		return &token
	}
	return nil
}

func DeleteToken(id uint) (bool) {
	var token Token
	db := DatabaseConnection()
	db.Find(&token, id)
	if &token != nil && token.ID == id {
		return db.Delete(&token).RowsAffected == 1
	}
	return false
}
