package pkg

import (
	"UserTask/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() {

	urlDB := "postgres://postgres:@kim1001@localhost:5432/postgres?sslmode=disable"
	var err error
	db, err = sqlx.Connect("postgres", urlDB)
	if err != nil {
		fmt.Printf("error open db: %v\n", err)
		return
	}
}

func GetPGInsance() *sqlx.DB {
	return db
}

func GetUsersDB() []models.Users {
	pg := GetPGInsance()
	var users []models.Users
	var user models.Users
	rows, err := pg.Query("select login, first_name, last_name, middle_name , job_title from employeer")
	if err != nil {
		fmt.Println("Error select from employeer func GetUsersDB")
		return nil
	}
	for rows.Next() {

		err = rows.Scan(&user.Login, &user.FirstName, &user.LastName, &user.MiddleName, &user.JobTitle)
		users = append(users, user)
	}
	err = rows.Err()

	return users
}

func GetCode(login string) int32 {
	pg := GetPGInsance()
	row := pg.QueryRow("select code from employeer_x where login =$1", login)
	var code int32
	err := row.Scan(&code)
	if err != nil {
		return -1
	}
	return code
}
