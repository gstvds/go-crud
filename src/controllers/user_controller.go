package controllers

import (
	"encoding/json"
	"errors"
	"go-crud/src/external/providers"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/repositories"
	"go-crud/src/shared"
	"go-crud/src/shared/responses"
	"go-crud/src/usecases"
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
	messagingProvider := providers.NewKafkaMessagingProvider()
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository, messagingProvider)

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
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	upsertUserUseCase := usecases.NewUpsertUserUseCase(userRepository)

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

	data := usecases.UpsertUserInputDTO{
		Id: userId,
	}
	if err := json.Unmarshal(body, &data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Invalid request body",
			"invalid_request_body",
			err,
		))
		return nil
	}

	if updatedUser, err := upsertUserUseCase.Exec(data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Failed to upser user",
			"upsert_failed",
			err,
		))
		return nil
	} else {
		responses.JSON(fiberContext, http.StatusOK, updatedUser)
		return nil
	}
}

// Get an User by its ID
func (UserController) Get(fiberContext *fiber.Ctx) error {
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	getUserUseCase := usecases.NewGetUserUseCase(userRepository)

	userId := fiberContext.Params("id")

	if userId == "" {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
			"Missing user id",
			"missing_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	user := usecases.GetUserInputDTO{
		UserId: userId,
	}

	if foundUser, err := getUserUseCase.Exec(user); err != nil {
		responses.Error(fiberContext, http.StatusNotFound, shared.Error(
			"Unable to find user. Try again or check if the user exists",
			"unable_to_find_user",
			err,
		))
		return nil
	} else {
		responses.JSON(fiberContext, http.StatusOK, foundUser)
		return nil
	}
}

// Delete an User by its ID
func (UserController) Delete(fiberContext *fiber.Ctx) error {
	db := database.Get()
	userRepository := repositories.NewPrismaUserRepository(db)
	deleteUserUseCase := usecases.NewDeleteUserUseCase(userRepository)

	userId := fiberContext.Params("id")
	if userId == "" {
		responses.Error(fiberContext, http.StatusBadRequest, shared.Error(
			"Invalid user id",
			"invalid_user_id",
			errors.New("missing_user_id"),
		))
		return nil
	}

	data := usecases.DeleteUserInputDTO{
		UserId: userId,
	}

	if err := deleteUserUseCase.Exec(data); err != nil {
		responses.Error(fiberContext, http.StatusInternalServerError, shared.Error(
			"Error while deleting user",
			"deletion_failed",
			err,
		))
		return nil
	} else {
		responses.JSON(fiberContext, http.StatusNoContent, nil)
		return nil
	}
}
