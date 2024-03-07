package handlers

import (
	"crud/models"
	"crud/models/objects"
	"crud/utils"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {

	results, err := models.GetAllStudents(c)
	if err != nil {
		utils.ErrorResponse(c, 400, err)
		return
	}
	utils.SuccessResponse(c, 200, results)

}

func GetStudent(c *gin.Context) {
	uuid, err := edgedb.ParseUUID(c.Param("id"))

	if err != nil {

		utils.ErrorResponse(c, 404, err.Error())
		return
	}

	results, err := models.GetStudent(c, uuid)
	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	utils.SuccessResponse(c, 200, results)

}

func AddNewStudent(c *gin.Context) {
	var students objects.Student

	if err := c.ShouldBindJSON(&students); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	fmt.Println("**************************add new student ***********************")

	body := map[string]interface{}{
		"name":   students.Name,
		"email":  students.Email,
		"class":  students.Class,
		"school": students.School,
	}
	

	userD, _ := models.CheckStudentExist(c, body["email"])

	if userD == true {
		utils.ErrorResponse(c, 404, "a user exist with this email")
		return
	}

	return

	res, err := models.AddNewStudent(c, body)

	fmt.Println(err)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	utils.SuccessResponse(c, 200, res)

}

func UpdateStudent(c *gin.Context) {

	uuid, err := edgedb.ParseUUID(c.Param("id"))

	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	_, checkUserErr := models.GetStudent(c, uuid)

	if checkUserErr != nil {
		utils.ErrorResponse(c, 400, checkUserErr.Error())
		return
	}

	var students objects.Student
	// add validators here
	if err := c.BindJSON(&students); err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	body := map[string]interface{}{
		"name":   students.Name,
		"email":  students.Email,
		"class":  students.Class,
		"school": students.School,
		"id":     uuid,
	}
	res, err := models.UpdateStudent(c, body)

	fmt.Println(err)
	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}
	utils.SuccessResponse(c, 200, res)

}

func DeleteStudent(c *gin.Context) {
	uuid, err := edgedb.ParseUUID(c.Param("id"))

	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	_, checkUserErr := models.GetStudent(c, uuid)

	if checkUserErr != nil {
		utils.ErrorResponse(c, 400, checkUserErr.Error())
		return
	}

	res, err := models.DeleteStudent(c, uuid)

	if err != nil {
		utils.ErrorResponse(c, 400, err.Error())
		return
	}

	utils.SuccessResponse(c, 200, res)
}
