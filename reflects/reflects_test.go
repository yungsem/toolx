package reflects

import (
	"log"
	"testing"
	"time"
)

type Son struct {
	Son1 string
	Son2 string
}

type Model struct {
	Son
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
}
type User struct {
	Id       int    `ent:"primaryKey;autoIncrement;autoIncrementIncrement:2"`
	Name     string `ent:"notNull;default:'ys'"`
	Age      int    `ent:"unsigned"`
	RealName string `ent:"notNull"`
	Model
}


func TestFields(t *testing.T) {
	u := &User{
		Id:   1,
		Name: "yangsen",
		Age:  30,
	}
	u.CreateUser = "1"
	u.CreateTime = "1"
	u.UpdateUser = "1"
	u.UpdateTime = "1"

	fields, err := Fields(u)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", fields)
}


type User3 struct {
	CreateTime time.Time
}

func TestValues(t *testing.T) {
	u := &User3{
		CreateTime: time.Now(),
	}

	fields, err := Values(u)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", fields)
}


type Model2 struct {
	CreateTime time.Time
	CreateUser string
	UpdateTime time.Time
	UpdateUser string
}

type User2 struct {
	Id       int
	Username string
	Password string
	RealName string
	Model2
}

func TestFieldValue(t *testing.T) {
	u := &User2{
		Id:       1,
		Username: "yangsen",
		Password: "123456",
		RealName: "杨森",
	}
	u.CreateUser = "yangsen"
	u.CreateTime = time.Now()
	u.UpdateUser = "yangsen"
	u.UpdateTime = time.Now()

	m, err := FieldValue(u)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", m)
}