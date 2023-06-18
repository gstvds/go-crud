package usercontroller

import (
	"encoding/json"
	"errors"
	"go-crud/src/domain/entities"
	"go-crud/src/shared"
	"go-crud/src/shared/responses"
	"go-crud/src/usecases"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Create a new User
func Create(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
			"Something went wrong while reading the body",
			"something_went_wrong",
			err,
		))
		return
	}

	var data entities.User
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return
	}

	if err = usecases.CreateUserUseCase(&data); err != nil {
		if err.Error() == "user already exists" {
			responses.Error(response, http.StatusInternalServerError, shared.Error(
				"User already exists",
				"user_already_exits",
				err,
			))
		} else {
			responses.Error(response, http.StatusInternalServerError, shared.Error(
				"Failed to create a user",
				"creation_failed",
				err,
			))
		}

		return
	}

	responses.JSON(response, http.StatusCreated, data)
}

// Update an existing User
func Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["id"]

	if userId == "" {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
			"Something went wrong while reading the body",
			"something_went_wrong",
			err,
		))
		return
	}

	var data entities.User
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return
	}
	log.Println(data)
	data.Id, err = uuid.Parse(userId)
	log.Println(data)

	if err != nil {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Something went wrong. Try again",
			"something_went_wrong",
			err,
		))
	}

	if err := usecases.UpsertUserUseCase(&data); err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
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
	userId := vars["id"]
	if userId == "" {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Missing user id",
			"missing_user_id",
			errors.New("missing_user_id"),
		))
		return
	}

	var user = entities.User{}
	convertedUserId, err := uuid.Parse(userId)

	if err != nil {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Something went wrong. Try again",
			"something_went_wrong",
			err,
		))
	}

	user.Id = convertedUserId
	err = usecases.GetUserUseCase(&user)

	if err != nil {
		responses.Error(response, http.StatusNotFound, shared.Error(
			"Unable to find user. Try again or check if the user exists",
			"unable_to_find_user",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusOK, user)
}

// Delete an User by its ID
func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["id"]
	if userId == "" {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return
	}

	var data entities.User
	userUuid, err := uuid.Parse(userId)

	if err != nil {
		responses.Error(response, http.StatusBadRequest, shared.Error(
			"Something went wrong. Try again",
			"something_went_wrong",
			err,
		))
	}

	data.Id = userUuid

	if err := usecases.DeleteUserUseCase(&data); err != nil {
		responses.Error(response, http.StatusInternalServerError, shared.Error(
			"Error while deleting user",
			"deletion_failed",
			err,
		))
		return
	}

	responses.JSON(response, http.StatusOK, nil)
}
