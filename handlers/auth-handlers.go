package handlers

import (
	"github/Shimaa-Ibrahim/todo/db"
	"github/Shimaa-Ibrahim/todo/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Error  []string
	SignUp bool
}

var data Data

func GetAuthHandler(w http.ResponseWriter, r *http.Request) {
	if userId, _ := r.Cookie("user_id"); userId != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	template, _ := template.ParseFiles("templates/auth.html")
	err := template.Execute(w, data)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data.Error = nil
	username := r.FormValue("username")
	password := r.FormValue("password")
	passwordsMatch := (password == r.FormValue("confirm"))

	if username == "" {
		data.Error = append(data.Error, "Error: Username required")
	}
	if password == "" {
		data.Error = append(data.Error, "Error: Passwords required")
	}
	if !passwordsMatch {
		data.Error = append(data.Error, "Error: Passwords do not match")
	}

	if len(data.Error) > 0 {
		handleError(w, r, true, "")
		return
	}
	hashpasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		handleError(w, r, true, "Error Occurred, Please Try Again")
		return
	}

	user := &models.User{Username: username, Password: string(hashpasswordBytes)}
	result := db.DB.Create(&user)
	if result.Error != nil {
		handleError(w, r, true, "Error Occurred, Please Try Again")
		return
	}
	SetCookie(w, strconv.FormatUint(uint64(user.ID), 10), 5)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data.Error = nil
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" {
		data.Error = append(data.Error, "Error: Username required")
	}
	if password == "" {
		data.Error = append(data.Error, "Error: Password required")
	}

	if len(data.Error) > 0 {
		handleError(w, r, false, "")
		return
	}

	var user models.User
	result := db.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		handleError(w, r, false, "Username or Password is incorrect!")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		handleError(w, r, false, "Username or Password is incorrect!")
		return
	}
	SetCookie(w, strconv.FormatUint(uint64(user.ID), 10), 60)
	http.Redirect(w, r, "/", http.StatusFound)
}

func SetCookie(w http.ResponseWriter, value string, expire time.Duration) {
	expiration := time.Now().Add(expire * time.Minute)
	cookie := http.Cookie{Name: "user_id", Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func handleError(w http.ResponseWriter, r *http.Request, SignUp bool, errMsg string) {
	data.SignUp = SignUp
	data.Error = append(data.Error, errMsg)
	http.Redirect(w, r, "/auth", http.StatusFound)
}
