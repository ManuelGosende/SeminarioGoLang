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
	Insert(Doctor) (*Doctor, error)
	FindByID(int) (*Doctor, error)
	Update(int, Doctor) (int, error)
	Delete(int) (int, error)
	FindAll() ([]*Doctor, error)
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
func (s service) FindAll() ([]*Doctor, error) {

	var list []*Doctor

	if err := s.db.Select(&list, "SELECT * FROM Doctor"); err != nil {
		fmt.Println("Error en la consulta, intentelo de nuevo")
		return nil, err
	}

	return list, nil

}

//Insert ...
func (s service) Insert(d Doctor) (*Doctor, error) {

	query := "INSERT INTO Doctor (name, enrollment, age) VALUES (?,?,?)"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println("Error en el método Prepare")
		return nil, err
	}

	_, err = stmtCreate.Exec(d.Name, d.Enrollment, d.Age)

	if err != nil {
		fmt.Println("Error en el método Exec")
		return nil, err
	}

	return &d, nil

}

//FindByID ...
func (s service) FindByID(Id int) (*Doctor, error) {

	var Doctor Doctor

	query := "SELECT * FROM Doctor WHERE ID=?"

	if err := s.db.Get(&Doctor, query, Id); err != nil {
		return nil, err
	}

	return &Doctor, nil

}

//Update ...
func (s service) Update(Id int, d Doctor) (int, error) {

	query := "UPDATE Doctor SET name = ?, enrollment = ?, age = ? WHERE id = :id"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println("Error en el método Prepare")
		fmt.Println(err.Error())
	}

	_, err = stmtCreate.Exec(d.Name, d.Enrollment, d.Age, Id)

	if err != nil {
		fmt.Println("Error en el método Exec")
		fmt.Println(err.Error())
	}

	return Id, nil

}

//Delete ...
func (s service) Delete(Id int) (int, error) {

	query := "DELETE FROM Doctor WHERE id = :id"

	stmtCreate, err := s.db.Prepare(query)

	if err != nil {
		fmt.Println("Error en el método Prepare")
		fmt.Println(err.Error())
	}

	_, err = stmtCreate.Exec(Id)

	if err != nil {
		fmt.Println("Error en el método Exec")
		fmt.Println(err.Error())
	}

	return Id, nil

}
