package models

import (
	"crud/models/objects"
	"crud/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) ([]objects.Student, error) {
	var result []objects.Student

	query := `select Students::Students {
		class,
		id,
		name,
		email,
		school
	  };`

	err := utils.DbClient.Query(c, query, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func GetStudent(c *gin.Context, uuid interface{}) (objects.Student, error) {

	var result objects.Student

	u := map[string]interface{}{
		"id": uuid,
	}

	query := `select Students::Students {
		id,
		name,
		email,
		school,
		class
	  }
	  filter .id = <std::uuid>$id`

	err := utils.DbClient.QuerySingle(c, query, &result, u)

	return result, err

}

func UpdateStudent(c *gin.Context, body map[string]interface{}) (objects.CreateStudent, error) {

	var result objects.CreateStudent

	query := `UPDATE Students::Students FILTER .id = <uuid>$id SET {
		name := <optional str>$name,
        email := <optional str>$email,
        school := <optional str>$school,
        class := <optional str>$class
	};`

	err := utils.DbClient.QuerySingle(c, query, &result, body)

	return result, err

}

func DeleteStudent(c *gin.Context, uuid interface{}) (objects.CreateStudent, error) {

	var result objects.CreateStudent

	u := map[string]interface{}{
		"id": uuid,
	}

	q := `DELETE Students::Students filter .id = <uuid>$id ;`

	err := utils.DbClient.QuerySingle(c, q, &result, u)

	return result, err
}

func AddNewStudent(c *gin.Context, body map[string]interface{}) (objects.CreateStudent, error) {
	var result objects.CreateStudent

	fmt.Println("🚀 ~ file: students.go ~ line 57 ~ funcAddNewStudent ~ result : ----1  ", &result)
	fmt.Println("🚀 ~ file: students.go ~ line 57 ~ funcAddNewStudent ~ result :  ----2 ", body)

	query := `INSERT Students::Students {
        name := <str>$name,
        email := <str>$email,
        school := <optional str>$school,
        class := <optional str>$class
    };`

	err := utils.DbClient.QuerySingle(c, query, &result, body)
	// fmt.Println("🚀 ~ file: students.go ~ line 57 ~ funcAddNewStudent ~ result : ", result)

	return result, err
}

func CheckStudentExist(c *gin.Context, email interface{}) (bool, error) {
	var result []objects.Student

	e := map[string]interface{}{
		"email": email,
	}

	fmt.Println("email >>>>>>>>>> ", e)

	query := `select Students::Students {
		id,
		name,
		email,
		school,
		class
	  }
	  filter .email = <str>$email`

	err := utils.DbClient.Query(c, query, &result, e)
	
	if len(result) > 0 {
		return true, err
	}

	return false, err
}
