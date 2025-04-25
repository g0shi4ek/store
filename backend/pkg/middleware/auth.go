package middleware

import (
	"log"

	"github.com/g0shi4ek/store/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(requiredRole string, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)
		log.Println(tokenString)

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.StoreConf.SecretKey), nil
		})

		log.Println(claims.Audience, claims.Subject)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		if claims.Audience != requiredRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "Permissions denied"})
			return
		}

		c.Next()
	}
}

func ExtractToken(c *gin.Context) string {
	bearer := c.GetHeader("Authorization")
	log.Println(bearer)
	if len(bearer) > 7 && bearer[:7] == "Bearer " {
		return bearer[7:]
	}
	log.Println(bearer)
	return ""
}
