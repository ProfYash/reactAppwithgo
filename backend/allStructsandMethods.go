package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var jwtKey = []byte("secretkey")

var adminUsers = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
	"user3": "pass3",
}

type adminModel struct {
	gorm.Model
	Username string
	Password string
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}
type User struct {
	gorm.Model
	UID     string `json:"UID"`
	FName   string `json:"FName"`
	RollNo  string `json:"RollNo"`
	Contact string `json:"Contact"`
}
type Address struct {
	AddID            string `json:"addid"`
	UID              string `json:"UID"`
	AddressName      string `json:"addressname"`
	FirstLineAddress string `json:"firstlineadd"`
	City             string `json:"city"`
	Pincode          string `json:"pincode"`
	gorm.Model
}

var flagfordelete int
var user []User

func _deleteUserAtUid(RollNo string) {
	for index, u := range user {
		if u.RollNo == RollNo {
			user = append(user[:index], user[index+1:]...)
			flagfordelete = 1
			break
		}
	}
	if flagfordelete == 1 {
		flagfordelete = 0
	} else {
		fmt.Println("Invalid Roll No")
		flagfordelete = 2
	}
}
func getByRoll(RollNo string) bool {
	db, err := gorm.Open(sqlite.Open("address.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var usersForDisplay []User
	db.Find(usersForDisplay)
	for _, u := range usersForDisplay {
		if RollNo == u.RollNo {
			return true
		}
	}
	return false
}

func isValidCoockie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
func refreshToken(w http.ResponseWriter, r *http.Request) {
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
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
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
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 30)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "tokenRefreshed",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
