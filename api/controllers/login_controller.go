package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kenshin579/analyzing-Go-JWT-Postgres-Mysql-Restful-API/api/auth"
	"github.com/kenshin579/analyzing-Go-JWT-Postgres-Mysql-Restful-API/api/models"
	"github.com/kenshin579/analyzing-Go-JWT-Postgres-Mysql-Restful-API/api/responses"
	"github.com/kenshin579/analyzing-Go-JWT-Postgres-Mysql-Restful-API/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	//todo : server.DB.Debug() 코드상으로 이렇게 되어 있어서 production인 경우에는 코드를 수정해줘야 하는 단점이 있음
	//이렇게는 사용하지는 말자
	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
