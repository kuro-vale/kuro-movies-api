package middleware

import (
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/database"
	"github.com/kuro-vale/kuro-movies-api/models"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "ID"

func InitJWTMiddleware() *jwt.GinJWTMiddleware {
	var SECRET_KEY string = os.Getenv("SECRET_KEY")
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "kuro-movies-api",
		Key:         []byte(SECRET_KEY),
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			var user models.User
			database.DB.First(&user, "email = ?", claims[identityKey].(string))
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginRequest models.UserRequest
			if err := c.BindJSON(&loginRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user models.User
			if err := database.DB.First(&user, "email = ?", loginRequest.Email).Error; err == nil {
				if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err == nil {
					return user, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:", err.Error())
	}
	return authMiddleware
}
