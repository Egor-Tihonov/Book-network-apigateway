package service

import (
	"net/http"
	"time"
)

func SetCookies(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	return cookie
}
