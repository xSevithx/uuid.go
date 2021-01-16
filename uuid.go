package main

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/user"
	time2 "time"
)

func main() {
	database, _ := sql.Open("sqlite3","./Client.db")
	createTable, _ := database.Prepare("CREATE TABLE IF NOT EXISTS connection (id integer primary key, UUID text, customerCode text, installedBy text, time text, computerName text)")
	createTable.Exec()

	UUID := createUUID()
	RFC822 := getTime()
	Username := getUsername()
	computerName := getComputerName()

	statement, _ := database.Prepare("INSERT INTO connection (UUID, installedBy, time, computerName) VALUES (?,?,?,?)")
	statement.Exec(UUID, Username, RFC822, computerName)
}

func createUUID() string{
	id := uuid.New()
	var UUID string
	UUID = id.String()
	return  UUID
}
func getTime() string  {
	time := time2.Now()
	RFC822 := time.Format(time2.RFC822)
	return RFC822
}
func getUsername() string{
	newUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	username := newUser.Username
	return username
}
func getComputerName() string {
	computerName, err := os.Hostname()
	if err != nil{
		log.Fatal(err)
	}
	return  computerName
}
