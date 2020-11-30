package service

import (
	"SeminarioGoLang/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//Doctor ...
type Doctor struct {
	Id         int
	Name       string
	Enrollment string
	Age        int
}

//DoctorService ...
type DoctorService interface {
	Insert(Doctor) error
	FindByID(int) *Doctor
	Update(int, Doctor) int
	Delete(int) int
	FindAll() []*Doctor
}

type service struct {
	db     *sqlx.DB
	config *config.Config
}

//New ...
func New(db *sqlx.DB, c *config.Config) (DoctorService, error) {
	return service{db, c}, nil
}

//FindAll ...
func (s service) FindAll() []*Doctor {

	var list []*Doctor

	if err := s.db.Select(&list, "SELECT * FROM Doctor"); err != nil {
		panic(err)
	}

	return list

}

//Insert ...
func (s service) Insert(d Doctor) error {

	query := "INSERT INTO Doctor (name, enrollment, age) VALUES (?,?,?)"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println("Error in Prepare")
		return err
	}

	fmt.Println(d)
	_, err = stmtCreate.Exec(d.Name, d.Enrollment, d.Age)

	if err != nil {
		fmt.Println("Error in Exec")
		return err
	}

	return nil

}

//FindByID ...
func (s service) FindByID(Id int) *Doctor {

	var Doctor Doctor

	query := "SELECT * FROM Doctor WHERE ID=?"

	if err := s.db.Get(&Doctor, query, Id); err != nil {
		return nil
	}

	return &Doctor

}

//Update ...
func (s service) Update(Id int, d Doctor) int {

	query := "UPDATE Doctor SET name = ?, enrollment = ?, age = ? WHERE id = :id"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmtCreate.Exec(d.Name, d.Enrollment, d.Age, Id)

	if err != nil {
		fmt.Println(err.Error())
	}

	return Id

}

//Delete ...
func (s service) Delete(Id int) int {

	query := "DELETE FROM Doctor WHERE id = :id"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmtCreate.Exec(Id)

	if err != nil {
		fmt.Println(err.Error())
	}

	return Id

}
