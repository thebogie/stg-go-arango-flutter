package route

import (
	config "back/configs"
	createStudent "back/controllers/student-controllers/create"
	handlerCreateStudent "back/handlers/student-handlers/create"

	resultsStudent "back/controllers/student-controllers/results"
	handlerResultsStudent "back/handlers/student-handlers/results"

	resultStudent "back/controllers/student-controllers/result"
	handlerResultStudent "back/handlers/student-handlers/result"

	deleteStudent "back/controllers/student-controllers/delete"
	handlerDeleteStudent "back/handlers/student-handlers/delete"

	updateStudent "back/controllers/student-controllers/update"
	handlerUpdateStudent "back/handlers/student-handlers/update"

	middleware "back/middlewares"

	"github.com/gin-gonic/gin"
)

func InitStudentRoutes(db *config.DBconnection, route *gin.Engine) {

	/**
	@description All Handler Student
	*/

	createStudentRepository := createStudent.NewRepositoryCreate(*db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(*db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(*db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(*db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(*db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
