package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"regexp"

	_ "github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "162004"
	dbname   = "postgres"
)

var (
	Objects map[string]*Object
)

type Object struct {
	ObjectId string
	// Score       int64
	// PlayerName  string
	FirstName   string 
	LastName    string
	Email       string
	Phone       string
	Password    string
	DateOfBirth string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func init() {
	Objects = make(map[string]*Object)
	fmt.Println("WOrking...........................")
	// Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	// Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
}

func AddOne(object Object) (ObjectId string) {
	fmt.Println(object)

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	//Phone Number validation
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	// End


	//Date validation
	ree := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	//End
	

	//Hash Password
	hash, _ := HashPassword(object.Password) // ignore error for the sake of simplicity

	fmt.Println("Password:", object.Password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(object.Password, hash)
	fmt.Println("Match:   ", match)
	//End
	fmt.Println(isEmailValid(object.Email), re.MatchString(object.Phone), ree.MatchString(object.DateOfBirth))
	if isEmailValid(object.Email) && re.MatchString(object.Phone) && ree.MatchString(object.DateOfBirth) {
		sqlStatement := `INSERT INTO "Person_Information" ("FirstName", "LastName", "Email", "Phone", "Password", "DateOfBirth")
			VALUES ($1, $2, $3, $4, $5, $6)`
		_, err = db.Exec(sqlStatement, object.FirstName, object.LastName, object.Email, object.Phone, hash, object.DateOfBirth)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nRow inserted successfully!")
		}
	} else {
		fmt.Println("\nInvalid Email Or Phone Number Or Date Of Birth!")
		object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
		Objects[object.ObjectId] = &object
		return "Invalid Email Or Phone Number Or Date Of Birth!"
	}

	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectId] = &object
	return "Data Inserted Successfully!"
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

func Update(ObjectId string, Email string) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Email = Email
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
