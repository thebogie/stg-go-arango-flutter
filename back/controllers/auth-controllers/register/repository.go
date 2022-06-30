package registerAuth

import (
	config "back/configs"
	model "back/models"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *config.DBconnection
}

func NewRepositoryRegister(db *config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	//db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	/*
		checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

		if checkUserAccount.RowsAffected > 0 {
			errorCode <- "REGISTER_CONFLICT_409"
			return &users, <-errorCode
		}

		users.Fullname = input.Fullname
		users.Email = input.Email
		users.Password = input.Password

		addNewUser := db.Debug().Create(&users)
		db.Commit()

		if addNewUser.Error != nil {
			errorCode <- "REGISTER_FAILED_403"
			return &users, <-errorCode
		} else {
			errorCode <- "nil"
		}
	*/
	return &users, <-errorCode
}
