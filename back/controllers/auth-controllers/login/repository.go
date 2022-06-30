package loginAuth

import (
	config "back/configs"
	model "back/models"
	util "back/utils"
	"context"
	"log"

	"github.com/arangodb/go-driver"
)

type Repository interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *config.DBconnection
}

func NewRepositoryLogin(db *config.DBconnection) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var player model.EntityPlayer
	var cursor driver.Cursor
	//db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	log.Printf("LOGINREPOISTORY INPUT:%v", input)

	player.Email = input.Email
	player.Password = input.Password

	ctx := context.Background()
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": player.Email,
	}
	cursor, err := r.db.Connect.Query(ctx, query, bindVars)
	if err != nil {
		errorCode <- "Incorrect Query"
		return input, <-errorCode
	}
	defer cursor.Close()

	var returndoc model.EntityPlayer
	//var metadata driver.DocumentMeta

	_, err = cursor.ReadDocument(ctx, &returndoc)

	if err != nil {
		errorCode <- "User not found"
		return input, <-errorCode
	}

	for {
		var doc model.EntityPlayer
		var metadata driver.DocumentMeta

		metadata, err = cursor.ReadDocument(nil, &doc)

		if driver.IsNoMoreDocuments(err) {
			break

		} else if err != nil {
			log.Fatalf("doc returned %v", err)
		} else {
			log.Println("Doc: ", metadata, doc)
		}

	}

	comparePassword := util.ComparePassword(returndoc.Password, player.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return input, <-errorCode
	} else {
		errorCode <- "nil"
	}

	/*
		checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

		if checkUserAccount.RowsAffected < 1 {
			errorCode <- "LOGIN_NOT_FOUND_404"
			return &users, <-errorCode
		}

		if !users.Active {
			errorCode <- "LOGIN_NOT_ACTIVE_403"
			return &users, <-errorCode
		}


	*/

	return input, <-errorCode
}
