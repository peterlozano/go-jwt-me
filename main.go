package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(
		`-----BEGIN PUBLIC KEY-----
<YOUR PUBLIC KEY HERE>
-----END PUBLIC KEY-----`))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    publicKey,
		SigningMethod: "RS256",
	}))

	e.GET("/", me)

	e.Logger.Fatal(e.Start(":1323"))
}

func me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	b, _ := json.Marshal(user)

	return c.JSONBlob(http.StatusOK, b)
}
