package torsoSql

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type User struct {
	Name string
	ID   int
}

func getTorsoURL(dbName string) string {
	// get torso db url eg. turso db show <database-name> --url
	cmd := exec.Command("turso", "db", "show", dbName, "--url")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getTorsoToken(dbName string) string {
	cmd := exec.Command("turso", "db", "tokens", "create", dbName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getTorsoConnection(dbName string) *sql.DB {
	// get Torso URL
	torsoURL := getTorsoURL(dbName)
	if torsoURL == "" {
		fmt.Fprintf(os.Stderr, "failed to get Torso URL for db %s", dbName)
		os.Exit(1)
	}

	// get Torso Token
	torsoToken := getTorsoToken(dbName)
	if torsoToken == "" {
		fmt.Fprintf(os.Stderr, "failed to get Torso token for db %s", dbName)
		os.Exit(1)
	}

	// combine url and token
	compURL := fmt.Sprintf("%s?authToken=%v", torsoURL, torsoToken)

	db, err := sql.Open("libsql", compURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", compURL, err)
		os.Exit(1)
	}

	return db
}

func QueryUsers(dbName string) []User {
	db := getTorsoConnection(dbName)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			fmt.Println("Error scanning row:", err)
			return users
		}

		users = append(users, user)
		fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return users
}
