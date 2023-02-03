package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/ent"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/hashings"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/service"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var uname string
var expectedPassword string
var upassword string

// jwt key..
var jwtKey = []byte("secret_key")

// user cretentials..
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// claims..
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// signup handler..
func Signup(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uname = credentials.Username

	var upasswordd, hasherr = hashings.HashPassword(credentials.Password)
	upassword = upasswordd
	if hasherr != nil {
		println(fmt.Println("Error hashing password"))
		return
	}
	if _, err = service.Createnewuser(ctx, client, uname, upassword); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

// login handler..
func Login(w http.ResponseWriter, r *http.Request) {
	client, err01 := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Densityasmt password=Anirudh@123 sslmode=disable")
	if err01 != nil {
		log.Fatalf("failed opening connection to postgres: %v", err01)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err02 := client.Schema.Create(context.Background()); err02 != nil {
		log.Fatalf("failed creating schema resources: %v", err02)
	}
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uname = credentials.Username
	var u, erro = service.Getuserbyname(ctx, client, uname)
	log.Println("users returened: ", u)
	if len(u) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expectedPassword = u[0].Password
	if !hashings.DoPasswordsMatch(expectedPassword, credentials.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

// home handler..
func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

// refresh handler..
func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

// logout..
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "logout_token",
			Value:   tokenString,
			Expires: time.Now(),
		})
}
