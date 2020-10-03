package controller

import (
	"github.com/oxlb/GoLangPostgresHelloWorld/model"
	db "github.com/oxlb/GoLangPostgresHelloWorld/storage"
	"github.com/oxlb/GoLangPostgresHelloWorld/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	students, err := GetRepoStudents()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := db.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
