package controllers

import (
	"context"
	"log"
	"net/http"
	d "studentApi/database"
	m "studentApi/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func ListUsers(ctx *gin.Context) {
	var users []m.User
	query := "SELECT * FROM users"
	rows, err := d.DBEngine.Pool.Query(context.Background(), query)
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
	//var user []m.User
	user := m.InsertUserRequest{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Geçersiz istek! Eksik veya hatalı alan var"})
		return
	}

	var roleID uuid.UUID
	roleIdQuery := `SELECT ID FROM roles WHERE role_name = $1`
	if err := d.DBEngine.Pool.QueryRow(context.Background(), roleIdQuery, user.RoleName).Scan(&roleID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error":"Belirtilen rol bulunamadı"})
		return
	}

	insertQuery := `
		INSERT INTO users (ID, name, surname, username, email, roleID) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	userID := uuid.New()
	_, err := d.DBEngine.Pool.Exec(
		context.Background(), 
		insertQuery, 
		userID,
		user.Name,
		user.Surname, 
		user.Username,
		user.Email,
		roleID,
	)

	if err != nil {
		log.Println("Kullanıcı ekleme hatası:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Kullanıcı eklenemedi"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id" : userID,
		"name" : user.Name,
		"surname" : user.Surname, 
		"username" : user.Username,
		"email" : user.Email,
		"role_id" : roleID,
		"role_name" : user.RoleName,
	})
}
