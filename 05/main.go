package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"name", binding:"required"`
	Age      int    `json:"age", binding:"omitempty"`
	Email    string `json:"email", binding:"required,email"`
	Password string `json:"password", binding:"required,customPassword"`
}

func validatePassword(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if len(password) > 20 && len(password) < 8 {
		return false
	}

	return true
}

type UserList struct {
	Users []User
}

func (u *UserList) listUsers() []User {
	return u.Users
}

func (u *UserList) addUser(user User) {
	u.Users = append(u.Users, user)
}

func main() {
	router := gin.New()
	userList := UserList{}

	router.Use(gin.Logger(), gin.Recovery())

	v1 := router.Group("/v1")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("customPassword", validatePassword)
	}

	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "hello world")
		})

		v1.POST("/user", func(ctx *gin.Context) {
			name := ctx.Query("name")
			age := ctx.Query("age")

			ctx.JSON(http.StatusOK, gin.H{
				"status":   http.StatusOK,
				"message":  "success",
				"name":     name,
				"age":      age,
				"password": "111",
			})
		})

		v1.POST("/profile", func(ctx *gin.Context) {
			user := User{}
			err := ctx.ShouldBindJSON(&user)

			if err != nil {
				ctx.String(http.StatusBadRequest, "Plese enter valid credentials", err.Error())
				return
			}
			userList.addUser(user)
			ctx.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "you've been added to our database",
			})
		})

		v1.GET("/userlist", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"users": userList.listUsers(),
			})
		})

	}
	router.Run(":4000")
}
