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
	list = append(list, &endpoint {
		method:   "POST",
		path:     "/doctor",
		function: Insert(d),
	})

	// R
	list = append(list, &endpoint {
		method:   "GET",
		path:     "/doctors",
		function: FindAll(d),
	})

	// U
	list = append(list, &endpoint {
		method:   "PUT",
		path:     "/doctor/:id",
		function: Update(d),
	})

	// D
	list = append(list, &endpoint {
		method:   "DELETE",
		path:     "/doctor/:id",
		function: Delete(d),
	})

	list = append(list, &endpoint {
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
			fmt.Println(err)
			os.Exit(1)
		}

		var doctor Doctor

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = d.Insert(doctor)

		if err != nil {
			fmt.Println("No se pudo insertar Doctor")
			fmt.Println(err.Error())
		}

	}

}

func FindAll(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Data": d.FindAll(),
		})
	}

}

func Update(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		data, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var doctor Doctor

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		Id, err := strconv.Atoi(c.Param("id"))

		if err = json.Unmarshal(data, &doctor); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		d.Update(Id, doctor)

		c.JSON(http.StatusOK, gin.H {
			"Message": "Se modificó el doctor",
		})

	}

}

func Delete(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		d.Delete(strconv.Atoi(c.Param("id")))
		c.JSON(http.StatusOK, gin.H {
			"Message": "Se eliminó el Doctor",
		})

	}

}

func FindByID(d DoctorService) gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		c.JSON(http.StatusOK, gin.H {
			"Beer": d.FindByID(id),
		})

	}

}