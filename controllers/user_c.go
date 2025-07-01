package controllers

import (
	"context"
	"log"
	"net/http"
	d "studentApi/database"
	m "studentApi/models"

	"github.com/gin-gonic/gin"
)

var users []m.User

func ListUsers(ctx *gin.Context) {
	rows, err := d.DBEngine.Pool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Println("DB sorgu hatası: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB hatası"})
	}
	defer rows.Close()

	for rows.Next() {
		user := m.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Username, &user.Email, &user.RoleId); err != nil {
			log.Println("Okuma hatası", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Veri okuma hatası"})
			return
		}

		users = append(users, user)
	}
	ctx.JSON(http.StatusOK, users)
}

func CreateUser(ctx *gin.Context) {

}
