package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"rein/final/controllers"
	"rein/final/database"
	"rein/final/repository"
	"rein/final/structs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	jwt "github.com/appleboy/gin-jwt/v2"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func main() {
	//psqlInfo := fmt.Sprintf("host=#{host} port=#{port} dbname=#{dbname} user=#{user} password=#{password} sslmode=disable")
	//psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load config")
	} else {
		fmt.Println("loaded config")
	}
	//psqlInfo := fmt.Sprintf("host=#{os.Getenv("DB_HOST")} port=#{os.Getenv("DB_PORT")} dbname=#{os.Getenv("DB_NAME")} user=#{os.Getenv("DB_USER")} password=#{os.Getenv("DB_PASSWORD")} sslmode=disable")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("PGHOST"),
	// 	os.Getenv("PGPORT"),
	// 	os.Getenv("PGUSER"),
	// 	os.Getenv("PGPASSWORD"),
	// 	os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection established")
	}
	database.DbMigrate(DB)
	defer DB.Close()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*structs.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &structs.User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			password := loginVals.Password

			var user structs.User
			user.UserName = loginVals.Username

			users, _ := repository.GetUserByUsername(database.DbConnection, user)
			if password == users.Password {
				return &structs.User{
					UserName: users.UserName,
					Email:    users.Email,
					Name:     users.Name,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	if v, ok := data.(*structs.User); ok && v.UserName == "admin" {
		// 		return true
		// 	}

		// 	return false
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	router := gin.Default()
	router.GET("/user", authMiddleware.MiddlewareFunc(), controllers.GetAllUser)
	router.GET("/user/:id", authMiddleware.MiddlewareFunc(), controllers.GetUserById)
	router.POST("/user", authMiddleware.MiddlewareFunc(), controllers.InsertUser)
	router.PUT("/user/:id", authMiddleware.MiddlewareFunc(), controllers.UpdateUser)
	router.DELETE("/user/:id", authMiddleware.MiddlewareFunc(), controllers.DeleteUser)

	router.POST("/register", controllers.InsertUser)
	router.POST("/login", authMiddleware.LoginHandler)

	router.GET("/article", controllers.GetAllArticle)
	router.GET("/article/:id", controllers.GetArticleById)
	router.POST("/article", authMiddleware.MiddlewareFunc(), controllers.InsertArticle)
	router.PUT("/article/:id", authMiddleware.MiddlewareFunc(), controllers.UpdateArticle)
	router.DELETE("/article/:id", authMiddleware.MiddlewareFunc(), controllers.DeleteArticle)

	router.GET("/comment", controllers.GetAllComment)
	router.GET("/comment/:id", controllers.GetCommentById)
	router.POST("/comment", authMiddleware.MiddlewareFunc(), controllers.InsertComment)
	router.PUT("/comment/:id", authMiddleware.MiddlewareFunc(), controllers.UpdateComment)
	router.DELETE("/comment/:id", authMiddleware.MiddlewareFunc(), controllers.DeleteComment)

	router.GET("/like", controllers.GetAllLike)
	router.GET("/like/:id", controllers.GetLikeById)
	router.POST("/like", authMiddleware.MiddlewareFunc(), controllers.InsertLike)
	router.PUT("/like/:id", authMiddleware.MiddlewareFunc(), controllers.UpdateLike)
	router.DELETE("/like/:id", authMiddleware.MiddlewareFunc(), controllers.DeleteLike)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.Run("localhost:8080")
}
