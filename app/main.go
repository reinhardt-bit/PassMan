package main

import (
	"fmt"

	passMan "github.com/reinhardt-bit/PassMan"
	torsoSql "github.com/reinhardt-bit/PassMan/internal/sql"
	"github.com/reinhardt-bit/passMan/app/views/dashboard"
)

func main() {
	app := passMan.New()

	app.Get("/2fa", Handle2FAIndex)

	app.Start(":3000")
	// torsoUrl := torsoSql.GetTorsoURL("my-db")
	// fmt.Println("Torso url in main : ", torsoUrl) }
	//
	const dbName = "my-db"
	torsoUsers := torsoSql.QueryUsers(dbName)
	for _, u := range torsoUsers {
		fmt.Println(u)
	}
}

func Handle2FAIndex(c *passMan.Context) error {
	// ctx.HTML(templates.Index)
	return c.Render(dashboard.Index())
}
