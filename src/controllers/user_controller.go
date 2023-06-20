package controllers

import (
	"encoding/json"
	"errors"
	"go-crud/src/domain/entities"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/repositories"
	"go-crud/src/shared"
	"go-crud/src/shared/responses"
	"go-crud/src/usecases"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

// NewUserController return a new instance of UserController
func NewUserController() *UserController {
	return &UserController{}
}

// Create a new User
func (UserController) Create(fiberContext *fiber.Ctx) error {
	body := fiberContext.Body()
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)

	data := usecases.CreateUserInputDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return nil
	}

	if createdUser, err := createUserUseCase.Exec(data); err != nil {
		if err.Error() == "user already exists" {
			responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
				"User already exists",
				"user_already_exits",
				err,
			))
		} else {
			responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
				"Failed to create a user",
				"creation_failed",
				err,
			))
		}

		return nil
	} else {
		responses.JSON(fiberContext, http.StatusCreated, createdUser)
		return nil
	}
}

// Update an existing User
func (UserController) Update(fiberContext *fiber.Ctx) error {
	userId := fiberContext.Params("id")

	if userId == "" {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	body := fiberContext.Body()

	var data entities.User
	if err := json.Unmarshal(body, &data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return nil
	}

	var err error
	log.Println(data)

	if err != nil {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
			"Something went wrong. Try again",
			"something_went_wrong",
			err,
		))
	}

	if err := usecases.UpsertUserUseCase(&data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Failed to upser user",
			"upsert_failed",
			err,
		))
		return nil
	}

	responses.JSON(fiberContext, http.StatusOK, data)
	return nil
}

// Get an User by its ID
func (UserController) Get(fiberContext *fiber.Ctx) error {
	userId := fiberContext.Params("id")
	if userId == "" {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
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
		responses.Error(fiberContext, http.StatusNotFound, shared.Error(
			"Unable to find user. Try again or check if the user exists",
			"unable_to_find_user",
			err,
		))
		return nil
	}

	responses.JSON(fiberContext, http.StatusOK, user)
	return nil
}

// Delete an User by its ID
func (UserController) Delete(fiberContext *fiber.Ctx) error {
	userId := fiberContext.Params("id")
	if userId == "" {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	var data entities.User
	data.Id = userId

	if err := usecases.DeleteUserUseCase(&data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Error while deleting user",
			"deletion_failed",
			err,
		))
		return nil
	}

	responses.JSON(fiberContext, http.StatusOK, nil)
	return nil
}
