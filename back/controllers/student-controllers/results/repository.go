package resultsStudent

import (
	config "back/configs"
	model "back/models"
)

type Repository interface {
	ResultsStudentRepository() (*[]model.EntityStudent, string)
}

type repository struct {
	db config.DBconnection
}

func NewRepositoryResults(db config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) ResultsStudentRepository() (*[]model.EntityStudent, string) {

	var students []model.EntityStudent
	//db := r.db.Model(&students)
	errorCode := make(chan string, 1)
	/*
		resultsStudents := db.Debug().Select("*").Find(&students)

		if resultsStudents.Error != nil {
			errorCode <- "RESULTS_STUDENT_NOT_FOUND_404"
			return &students, <-errorCode
		} else {
			errorCode <- "nil"
		}
	*/
	return &students, <-errorCode
}
