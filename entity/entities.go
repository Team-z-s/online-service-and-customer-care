package entity

import (
	"time"
)

//Admin stract containing admin info
type Admin struct {
	username string
	password string
}

//User struct containing users
type User struct {
	ID    int    `json: "id"`
	FName string `json: "fname" gorm:"type:varchar(255); not null"`
	LName string `json: "lname" gorm:"type:varchar(255); not null"`
	Email string `json: "email" gorm:"type:varchar(255);unique"`
	Phone string `json: "phone" gorm:"type:varchar(255)"`

	Username string `json: "username" gorm:"type:varchar(255); not null; unique"`
	Password string `json: "password" gorm:"type:varchar(255)"`
}

//Address contains country,city, woreda,kebele
type Address struct {
	country, city, woreda, kebele string
}

//Employee struct contains employees
type Employee struct {
	ID        int    `json:"id"`
	companyID int    `json:"company_id"`
	FName     string `json: "fname" gorm:"type:varchar(255)" `
	LName     string `json: "lname" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255);not null; unique"`

	Address  string  `json: "address"`
	Sallary  float64 `json: "sallary"`
	Phone    string  `json: "phone" gorm:"type:varchar(255)"`
	photo    string  `json: "photo" gorm:"type:varchar(255)"`
	Username string  `json: "username" gorm:"type:varchar(255);not null; unique"`
	Password string  `json: "password" gorm:"type:varchar(255) "`
}

//Company struct contain fields of the company
type Companie struct {
	ID       int    `json: "id"`
	FullName string `json: "fullname"`
	Logo     string `json: "logo"`
	Email    string `json: "email"`
	Phone    string `json: "phone"`
	Address  string `json: "address"`
	Moto     string `json: "moto"`
	Password string `json: "password" gorm:"type:varchar(255) "`
}

//Service contains fields of service
type Service struct {
	ID          int    `json: "id"`
	CompanyID   int    `json: "company_id"`
	Name        string `json: "service_name"`
	Description string `json: "description"`
}

//Date c
type Date struct {
	Day   int
	Month int
	Year  int
}

//Attendance contain date and time
type Attendance struct {
	ID         int       `json: "id"`
	EmployeeID int       `json: "employee_id"`
	Shift      string    `json: "shift"`
	Date       Date      `json: "data_a"`
	Status     string    `json: "status"`
	postedAt   time.Time `json: "posted_at"`
}

//Comment coment
type Comment struct {
	ID       uint   `json:"id"`
	username string `json:"usename" gorm:"type:varchar(255)"`
	Message  string `json:"message"`
	//Phone    string    `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null; unique"`
	PlacedAt time.Time `json:"posted_at"`
}

//Jobs employee's work
type Employee_job struct {
	EmployeeID int     `json: "employee_id"`
	CompanyID  int     `json: "company_id"`
	Job        string  `json: "job"`
	DayRate    float64 `json: "day_rate"`
	NightRate  float64 `json: "night_rate"`
}

//Roles role
type Roles struct {
	ID   int
	Name string
}

//Task taskes assigned for the employee
type Task struct {
	ID          int    `json: "id"`
	EmployeeID  int    `json: "employee_id"`
	Name        string `json: "task_name"`
	Description string `json: "description"`
	progress    string `json: "progress"`
}
