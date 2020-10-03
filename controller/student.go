package controller

import (
	model "github.com/oxlb/GoLangPostgresHelloWorld/model/student"
	db "github.com/oxlb/GoLangPostgresHelloWorld/storage"
)

func GetStudents(c echo.Context) error {
	students, err := GetRepoStudents()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model, error) {
	db := db.GetDBInstance()
	students := []model.Student{}

	if err := db.Find(&student).Error; err != nil {
		return nil, err
	}

	return students, nil
}


