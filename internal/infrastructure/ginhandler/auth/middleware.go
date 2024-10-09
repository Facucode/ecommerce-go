package auth

import (
	"ecommerce-go/internal/core/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"regexp"
)

var allowList = []*regexp.Regexp{
	regexp.MustCompile("^/swagger/.*"),
}

func TokenMiddleware(envVariables domain.Environment) gin.HandlerFunc {
	jwtSecretKey := envVariables.SecretKeyJWT
	if jwtSecretKey == "" {
		panic("Please set jwt environment variable")
	}

	return func(c *gin.Context) {
		for _, regex := range allowList {
			if regex.MatchString(c.Request.URL.Path) {
				c.Next()
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			respondWithError(c, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		tokenString := authHeader[len("Bearer "):]
		parsedTokenInfo, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil {
			respondWithError(c, http.StatusUnauthorized, fmt.Sprintf("Failed to parse token: %s", err))
			return
		}

		if !parsedTokenInfo.Valid {
			respondWithError(c, http.StatusForbidden, "Invalid token")
			return
		}

		claims := parsedTokenInfo.Claims.(jwt.MapClaims)

		subject := claims["name"]

		fmt.Println(subject)

		c.Next()
	}
}

func respondWithError(ctx *gin.Context, code int, message interface{}) {
	fmt.Println(message)
	ctx.AbortWithStatusJSON(code, gin.H{"error": message})
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "token" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
