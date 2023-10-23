package handlers

import (
	"database/sql"
	"fmt"
	"password_strength_api/services"
	"time"

	"github.com/gin-gonic/gin"
)

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

type PasswordResponse struct {
	NumOfSteps int `json:"num_of_steps"`
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (request PasswordRequest) String() string {
	return fmt.Sprintf("{\"init_password\": \"%s\"}", request.InitPassword)
}

func (response PasswordResponse) String() string {
	return fmt.Sprintf("{\"num_of_steps\": %d}", response.NumOfSteps)
}

func ResponseHandler(c *gin.Context, db *sql.DB) {
	var request PasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	actionsNeeded := services.ActionsNeededToMakeStrong(request.InitPassword)
	response := PasswordResponse{
		NumOfSteps: actionsNeeded,
	}
	insertdb := `insert into  logger (request,response,method,code,accesstime) values($1,$2,$3,$4,$5)`
	_, e := db.Exec(insertdb, request.String(), response.String(), c.Request.Method, c.Writer.Status(), time.Now().Add(7*time.Hour))
	CheckError(e)
	c.JSON(200, response)
}
