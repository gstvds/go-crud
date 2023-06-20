package usercontroller

import (
	"encoding/json"
	"errors"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	userrepository "go-crud/src/external/repositories/user_repository"
	"go-crud/src/shared"
	"go-crud/src/shared/responses"
	"go-crud/src/usecases"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Create a new User
func Create(context *fiber.Ctx) error {
	body := context.Body()
	db := database.Get()
	userRepository := userrepository.New(db)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)

	data := usecases.CreateUserInputDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		responses.Error(context, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return nil
	}

	if createdUser, err := createUserUseCase.Exec(data); err != nil {
		if err.Error() == "user already exists" {
			responses.Error(context, http.StatusInternalServerError, shared.Error(
				"User already exists",
				"user_already_exits",
				err,
			))
		} else {
			responses.Error(context, http.StatusInternalServerError, shared.Error(
				"Failed to create a user",
				"creation_failed",
				err,
			))
		}

		return nil
	} else {
		responses.JSON(context, http.StatusCreated, createdUser)
		return nil
	}
}

// Update an existing User
func Update(context *fiber.Ctx) error {
	userId := context.Params("id")

	if userId == "" {
		responses.Error(context, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	body := context.Body()

	var data entities.User
	if err := json.Unmarshal(body, &data); err != nil {
		responses.Error(context, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return nil
	}

	var err error
	log.Println(data)

	if err != nil {
		responses.Error(context, http.StatusBadRequest, shared.Error(
			"Something went wrong. Try again",
			"something_went_wrong",
			err,
		))
	}

	if err := usecases.UpsertUserUseCase(&data); err != nil {
		responses.Error(context, http.StatusInternalServerError, shared.Error(
			"Failed to upser user",
			"upsert_failed",
			err,
		))
		return nil
	}

	responses.JSON(context, http.StatusOK, data)
	return nil
}

// Get an User by its ID
func Get(context *fiber.Ctx) error {
	userId := context.Params("id")
	if userId == "" {
		responses.Error(context, http.StatusBadRequest, shared.Error(
			"Missing user id",
			"missing_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	var user = entities.User{}
	user.Id = userId

	err := usecases.GetUserUseCase(&user)

	if err != nil {
		responses.Error(context, http.StatusNotFound, shared.Error(
			"Unable to find user. Try again or check if the user exists",
			"unable_to_find_user",
			err,
		))
		return nil
	}

	responses.JSON(context, http.StatusOK, user)
	return nil
}

// Delete an User by its ID
func Delete(context *fiber.Ctx) error {
	userId := context.Params("id")
	if userId == "" {
		responses.Error(context, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	var data entities.User
	data.Id = userId

	if err := usecases.DeleteUserUseCase(&data); err != nil {
		responses.Error(context, http.StatusInternalServerError, shared.Error(
			"Error while deleting user",
			"deletion_failed",
			err,
		))
		return nil
	}

	responses.JSON(context, http.StatusOK, nil)
	return nil
}
