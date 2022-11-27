package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	godotenv.Load()
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/login", login)
	log.Println("Server started in port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func login(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.WriteHeader(405)
		rw.Write([]byte("Can not " + r.Method + " /login"))
		return
	}

	rw.Header().Add("Content-Type", "application/json")

	b, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(`{"message": "Bad Request"}`))
		log.Println("Error: ", err.Error())
		return
	}

	var u User

	err = json.Unmarshal(b, &u)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(`{"message": "Bad Request"}`))
		log.Println("Error: ", err.Error())
		return
	}

	db, err := getDB()
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"message": "Internal Server Error"}`))
		log.Println("Error: ", err.Error())
		return
	}

	stmt, err := db.Prepare("SELECT email, password FROM users WHERE email=?")
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"message": "Internal Server Error"}`))
		log.Println("Error: ", err.Error())
		return
	}

	rows, err := stmt.Query(u.Email)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"message": "Internal Server Error"}`))
		log.Println("Error: ", err.Error())
		return
	}

	var email, pass string
	if rows.Next() {
		err = rows.Scan(&email, &pass)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(`{"message": "Internal Server Error"}`))
			log.Println("Error: ", err.Error())
			return
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(u.Password))
	if err != nil || email != u.Email {
		rw.WriteHeader(401)
		rw.Write([]byte(`{"message": "Unauthorized"}`))
		log.Println("Error: ", err.Error())
		return
	}

	claims := jwt.MapClaims{
		"email":     email,
		"expiresAt": time.Now().Add(time.Hour * 24).UnixMilli(),
		"issuedAt":  time.Now().UnixMilli(),
		"idAdmin":   true,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	st, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte(`{"message": "Internal Server Error"}`))
		log.Println("Error: ", err.Error())
		return
	}

	rw.WriteHeader(200)
	rw.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, st)))
}

func getDB() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
