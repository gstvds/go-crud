package usercontroller

import (
	"encoding/json"
	"errors"
	"go-crud/src/infra/repositories/models"
	"go-crud/src/services"
	"go-crud/src/utils"
	"go-crud/src/utils/responses"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a new User
func Create(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Something went wrong while reading the body",
			"something_went_wrong",
			err,
		))
		return
	}

	var data models.User
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return
	}

	if err = services.CreateUserService(&data); err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Failed to create a user",
			"creation_failed",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusCreated, data)
}

// Update an existing User
func Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]

	if email == "" {
		responses.Error(response, http.StatusBadRequest, utils.Error(
			"Invalid user email",
			"invalid_user_email",
			errors.New("missing_user_email"),
		))
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Something went wrong while reading the body",
			"something_went_wrong",
			err,
		))
		return
	}

	var data models.User
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return
	}

	data.Email = email

	if err := services.UpsertUserService(&data); err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Failed to upser user",
			"upsert_failed",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusOK, data)
}

// Get an User by its ID
func Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]
	if email == "" {
		responses.Error(response, http.StatusBadRequest, utils.Error(
			"Missing user email",
			"missing_user_email",
			errors.New("missing_user_email"),
		))
		return
	}

	var data models.User
	data.Email = email

	if err := services.GetUserService(&data); err != nil {
		responses.Error(response, http.StatusNotFound, utils.Error(
			"Unable to find user. Try again or check if the user exists",
			"unable_to_find_user",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusOK, data)
}

// Delete an User by its ID
func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]
	if email == "" {
		responses.Error(response, http.StatusBadRequest, utils.Error(
			"Invalid user email",
			"invalid_user_email",
			errors.New("missing_user_email"),
		))
		return
	}

	var data models.User
	data.Email = email

	if err := services.DeleteUserService(&data); err != nil {
		responses.Error(response, http.StatusInternalServerError, utils.Error(
			"Error while deleting user",
			"deletion_failed",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusOK, nil)
}
