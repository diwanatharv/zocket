package handler

import (
	"awesomeProject6/api/service"
	"awesomeProject6/pkg/enums"
	"awesomeProject6/pkg/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func Updateproduct(e echo.Context) error {
	str, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	err = service.UpdateProduct(str)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, enums.Statusok)
}

func Createuser(e echo.Context) error {
	var reqbody models.User
	err := json.NewDecoder(e.Request().Body).Decode(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	err = service.CreateUser(reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, enums.Statusok)

}
func CreateProduct(e echo.Context) error {
	var reqbody models.Product
	err := e.Bind(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Validation)
	}
	err = service.CreateProduct(reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, enums.Statusok)
}
