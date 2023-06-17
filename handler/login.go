package handler

import (
	"gin/middleware"
	"gin/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JWT key: dipakai untuk membuat signature JWT token.

// Data key: user, password, role yang bisa digunakan untuk mengakses API
var users = map[string]*User{
	"aditira": {
		Password: "password1",
		Role:     "admin",
	},
	"dito": {
		Password: "password2",
		Role:     "student",
	},
}

type User struct {
	Password string
	Role     string
}

// Struct untuk membaca request body JSON
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Struct Claims digunakan sebagai object yang akan di encode atau di parse oleh JWT
// jwt.StandardClaims ditambahkan sebagai embedded type untuk memudahkan proses encoding, parsing dan validasi JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Exipires string `json:"expires"`
}

func Login(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildResponse(err.Error(), nil))
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword.Password != creds.Password {
		c.JSON(http.StatusUnauthorized, util.BuildResponse("Unauthorized", nil))
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		Role:     expectedPassword.Role,
		StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JWTSecretKey)

	if err != nil {
		// return internal error ketika ada kesalahan saat pembuatan JWT string
		c.JSON(http.StatusInternalServerError, util.BuildResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, util.BuildResponse("Login success", LoginResponse{
		Name:     "token",
		Token:    tokenString,
		Exipires: expirationTime.String(),
	}))

}
