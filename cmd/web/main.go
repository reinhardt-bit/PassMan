package main

import (
	"fmt"

	torsoSql "github.com/reinhardt-bit/PassMan/internal/sql"
)

func main() {
	// torsoUrl := torsoSql.GetTorsoURL("my-db")
	// fmt.Println("Torso url in main : ", torsoUrl) }
	//
	const dbName = "my-db"
	torsoUsers := torsoSql.QueryUsers(dbName)
	for _, u := range torsoUsers {
		fmt.Println(u)
	}
}
