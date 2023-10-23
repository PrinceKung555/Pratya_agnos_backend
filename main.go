package main

import (
	"database/sql"
	"fmt"
	handlers "password_strength_api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydb"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	handlers.CheckError(err)
	defer db.Close()
	r := gin.Default()
	api := r.Group("/api")
	{
		strongPasswordSteps := api.Group("/strong_password_steps")
		{
			strongPasswordSteps.POST("/", func(c *gin.Context) {
				handlers.ResponseHandler(c, db)
			})
		}
	}

	r.Run(":8080")
}
