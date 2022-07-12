package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var Setting *JwtSetting

type JwtSetting struct {
	SecretKey   string
	Issuer      string
	TokenExpire int
}

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"userName"`
}

func SetJwtSetting(s *JwtSetting) {
	Setting = s
}

type MyClaims struct {
	UserInfo
	jwt.StandardClaims
}

// GenToken Create a new token
func GenToken(user UserInfo) (string, error) {
	var jwtKey = []byte(Setting.SecretKey)

	c := MyClaims{
		UserInfo: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(Setting.TokenExpire)).Unix(),
			Issuer:    Setting.Issuer,
			Subject:   Setting.Issuer,
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// ParseToken Parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	var jwtKey = []byte(Setting.SecretKey)
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware Middleware of JWT
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Authorization is null in Header",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Invalid Token.",
			})
			c.Abort()
			return
		}
		// Store Account info into Context
		c.Set("id", mc.ID)
		c.Set("name", mc.Name)

		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
