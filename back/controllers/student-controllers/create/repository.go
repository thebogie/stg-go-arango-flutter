package createStudent

import (
	config "back/configs"
	model "back/models"
)

type Repository interface {
	CreateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string)
}

type repository struct {
	db config.DBconnection
}

func NewRepositoryCreate(db config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) CreateStudentRepository(input *model.EntityStudent) (*model.EntityStudent, string) {

	var students model.EntityStudent
	//db := r.db.Model(&students)
	errorCode := make(chan string, 1)

	/*
		checkStudentExist := db.Debug().Select("*").Where("npm = ?", input.Npm).Find(&students)

		if checkStudentExist.RowsAffected > 0 {
			errorCode <- "CREATE_STUDENT_CONFLICT_409"
			return &students, <-errorCode
		}

		students.Name = input.Name
		students.Npm = input.Npm
		students.Fak = input.Fak
		students.Bid = input.Bid

		addNewStudent := db.Debug().Create(&students)
		db.Commit()

		if addNewStudent.Error != nil {
			errorCode <- "CREATE_STUDENT_FAILED_403"
			return &students, <-errorCode
		} else {
			errorCode <- "nil"
		}
	*/
	return &students, <-errorCode
}
