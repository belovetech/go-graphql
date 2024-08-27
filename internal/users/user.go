package users

import (
	"database/sql"
	"log"

	database "github.com/belovetech/go-graphql/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("INSERT INTO Users (username, password) VALUES(?,?)")
	if err != nil {
		log.Panic(err)
	}
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		log.Panic("Hash password err: ", err)
	}
	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Panic(err)
	}
}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = ?")
	if err != nil {
		log.Panic(err)
	}
	row := stmt.QueryRow(username)

	var Id int
	err = row.Scan(&Id)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Panic(err)
		}
		return 0, err
	}

	return Id, nil
}

func (user *User) Authenticate() bool {
	stmt, err := database.Db.Prepare("SELECT Password FROM Users WHERE Username = ?")
	if err != nil {
		log.Panic(err)
	}

	row := stmt.QueryRow(user.Username)
	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Panic(err)
	}

	return CheckPasswordHash(hashedPassword, user.Password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
