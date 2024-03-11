package main

import (
	"fmt"

	passMan "github.com/reinhardt-bit/PassMan"
	"github.com/reinhardt-bit/PassMan/app/views/dashboard"
	torsoSql "github.com/reinhardt-bit/PassMan/internal/sql"
)

func main() {
	app := passMan.New()

	app.Get("/dashboard", Handle2FAIndex)

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
	// return c.Render(dashboard.index())
}
