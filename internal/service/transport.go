package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

//Register ...
func (h httpService) Register(gin *gin.Engine) {
	for _, e := range h.endpoints {
		gin.Handle(e.method, e.path, e.function)
	}
}

// NewHTTPTransport ...
func NewHTTPTransport(d DoctorService) HTTPService {
	endpoints := makeEndpoints(d)
	return httpService{endpoints}
}

func makeEndpoints(d DoctorService) []*endpoint {

	list := []*endpoint{}

	// C
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/doctor",
		function: Insert(d),
	})

	// R
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/doctors",
		function: FindAll(d),
	})

	// U
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/doctor/:id",
		function: Update(d),
	})

	// D
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/doctor/:id",
		function: Delete(d),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/doctor/:id",
		function: FindByID(d),
	})

	return list

}

func Insert(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		data, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			fmt.Println("Error en el método ReadAll")
			os.Exit(1)
		}

		var doctor Doctor

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println("Error en el método Unmarshal")
			os.Exit(1)
		}

		doc, err := d.Insert(doctor)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error en método Insert",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"person": doc,
			})
		}

	}

}

func FindAll(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		doctors, err := d.FindAll()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Error en la consulta, intentelo de nuevo",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Data": doctors,
			})
		}

	}

}

func Update(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		data, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			fmt.Println("Error en el método ReadAll")
			os.Exit(1)
		}

		var doctor Doctor

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println("Error en el método Unmarshal")
			os.Exit(1)
		}

		Id, err := strconv.Atoi(c.Param("id"))

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println("Error en el método Unmarshal")
			os.Exit(1)
		}

		doc, err := d.Update(Id, doctor)

		if err != nil {
			fmt.Println("Error en el método Update")
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Se modificó el doctor",
				"ID":      doc,
			})
		}

	}

}

func Delete(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			fmt.Println("Error en el método Atoi")
			os.Exit(1)
		}

		doc, err := d.Delete(id)

		if err != nil {
			fmt.Println("Error en el método Delete")
			os.Exit(1)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Message": "Se eliminó el Doctor",
				"ID":      doc,
			})
		}

	}

}

func FindByID(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		doc, err := d.FindByID(id)

		if err != nil {
			fmt.Println("Error en el método FindByID")
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H{
			"Doctor": doc,
		})

	}

}
