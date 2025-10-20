package restdb

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	db := ConnectPostgres()
	fmt.Println(db)
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Errorf("Ping: %v\n", err)
	}

	u := User{}
	fmt.Println(u)
	rows, err := db.Query(`SELECT "username" FROM "users"`)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	defer rows.Close()
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			t.Errorf("%v\n", err)
		}

		fmt.Println(username)
	}

	if err != nil {
		t.Errorf("%v\n", err)
	}

	log.Println("Populating PostgreSQL")
	user := User{ID: 0, Username: "mtsouk", Password: "admin", LastLogin: time.Now().Unix(), Admin: 1, Active: 1}
	if InsertUser(user) {
		fmt.Println("User inserted successfully.")
	} else {
		t.Errorf("Insert failed!\n")
	}

	mtsoukUser := FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if DeleteUser(mtsoukUser.ID) {
		fmt.Println("User Deleted.")
	} else {
		t.Errorf("User not Deleted.\n")
	}

	mtsoukUser = FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if DeleteUser(mtsoukUser.ID) {
		t.Errorf("User Deleted.\n")
	} else {
		fmt.Println("User not Deleted.")
	}

	if DeleteUser(mtsoukUser.ID) {
		t.Errorf("User Deleted.")
	} else {
		fmt.Println("User not Deleted.")
	}
}
