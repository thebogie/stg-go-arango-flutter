package resendAuth

import (
	config "back/configs"
	model "back/models"
)

type Repository interface {
	ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *config.DBconnection
}

func NewRepositoryResend(db *config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) ResendRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	//db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	/*
		users.Email = input.Email

		checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

		if checkUserAccount.RowsAffected < 1 {
			errorCode <- "RESEND_NOT_FOUD_404"
			return &users, <-errorCode
		}

		if users.Active {
			errorCode <- "RESEND_ACTIVE_403"
			return &users, <-errorCode
		} else {
			errorCode <- "nil"
		}
	*/
	return &users, <-errorCode
}
