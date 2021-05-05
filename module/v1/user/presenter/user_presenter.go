package presenter

import (
	"net/http"
	"strconv"
	"svc-users-go/module/v1/user/model"
	"svc-users-go/module/v1/user/usecase"
	"svc-users-go/utils"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase usecase.UseCase
}

func NewUserHandler(e *echo.Echo, userUseCase usecase.UseCase) {
	injectionHandler := &UserHandler{
		UserUseCase: userUseCase,
	}

	groupingPath := e.Group("/api/v1")
	groupingPath.GET("/users", injectionHandler.GetAllUsers)
	groupingPath.GET("/users/:id", injectionHandler.GetDetailUsers)
	groupingPath.POST("/users", injectionHandler.CreateNewUser)
	groupingPath.PUT("/users/:id", injectionHandler.UpdateUser)
	groupingPath.DELETE("/users/:id", injectionHandler.DeleteUser)
}

// GetAllUsers godoc
// @Summary Retrieves user based on given limit, name and page
// @Query limit path integer false "limit"
// @Query page path string false "page"
// @Query name path string false "name"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
// @Failure 400 {object} map[string]interface{}
// @Failure 204 {object} map[string]interface{}
// @Tags Users-endpoints
func (uh *UserHandler) GetAllUsers(ctx echo.Context) error {
	var (
		limitParam      = ctx.QueryParam("limit")
		pagesParam      = ctx.QueryParam("page")
		name            = ctx.QueryParam("name")
		convertLimit, _ = strconv.ParseInt(limitParam, 10, 64)
		convertPage, _  = strconv.ParseInt(pagesParam, 10, 64)
	)

	// Limiting and paging data
	findAllUserUseCase, errorHandlerUseCase := uh.UserUseCase.FindAllUsers(name, convertLimit, convertPage)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCase.Error(),
			"message": "Error when get usecase",
		})
	}

	if findAllUserUseCase == nil {
		return ctx.JSON(http.StatusNoContent, echo.Map{
			"error":   errorHandlerUseCase.Error(),
			"message": "No Content",
		})
	}

	// Count all data
	countAllDataUser, errorHandlerUseCaseCount := uh.UserUseCase.CountAllUsers()
	if !utils.GlobalErrorException(errorHandlerUseCaseCount) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":   errorHandlerUseCaseCount.Error(),
			"message": "Error when get usecase",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"count": len(findAllUserUseCase),
		"data":  findAllUserUseCase,
		"total": countAllDataUser,
		"limit": convertLimit,
		"page":  convertPage,
	})
}

// GetDetailUsers godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [get]
// @Failure 400 {object} map[string]interface{} StatusBadRequest
// @Tags Users-endpoints
func (uh *UserHandler) GetDetailUsers(ctx echo.Context) error {
	id := ctx.Param("id")

	findUserById, errorHandlerUseCase := uh.UserUseCase.FindUserById(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"data": findUserById,
	})
}

// CreateNewUser godoc
// @Summary Create New User
// @Router /users [post]
// @Param name body string true "Name user"
// @Param enumint query int false "int enums" Enums (1, 2, 3)
// @Param enumnumber query number false "int enums "Enums (1.1, 1.2, 1.3)
// @Param string query string false "string valida "minlength (5) maxlength (10)"
// @Param int query int false "int valid "mínimo (1) máximo (10)"
// @Param default query string false "string default" default (A)"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} StatusBadRequest
// @Tags Users-endpoints
func (uh *UserHandler) CreateNewUser(ctx echo.Context) error {
	userPayload := new(model.CreateUser)

	_, err := govalidator.ValidateStruct(userPayload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	errorHandlerBindJSON := ctx.Bind(userPayload)

	if !utils.GlobalErrorDatabaseException(errorHandlerBindJSON) {
		return errorHandlerBindJSON
	}

	saveUser, errorHandlerUseCase := uh.UserUseCase.CreateNewUser(userPayload)

	if !utils.GlobalErrorDatabaseException(errorHandlerUseCase) {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message":      "User created successfully",
		"created_user": saveUser,
	})
}

// UpdateUser godoc
// @Summary Update User
// @Param id path string true "User ID"
// @Router /users/{id} [put]
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} StatusBadRequest
// @Tags Users-endpoints
func (uh *UserHandler) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userUpdate := new(model.UpdateUser)

	errorHandlerBindJSON := ctx.Bind(userUpdate)

	if !utils.GlobalErrorDatabaseException(errorHandlerBindJSON) {
		return errorHandlerBindJSON
	}

	errorHandlerUpdate := uh.UserUseCase.UpdateUser(id, userUpdate)

	if !utils.GlobalErrorDatabaseException(errorHandlerUpdate) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Update gagal",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "User updated sucessfully"})
}

// DeleteUser godoc
// @Summary Delete User
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{} StatusBadRequest
// @Tags Users-endpoints
func (uh *UserHandler) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	errorHandlerDelete := uh.UserUseCase.DeleteUser(id)

	if !utils.GlobalErrorDatabaseException(errorHandlerDelete) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "User not found",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "User deleted sucessfully"})
}
