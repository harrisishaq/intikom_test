package controller

import (
	"intikom_test/model"
	"intikom_test/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{service}
}

func (ctrl *userController) InitRoute(g *gin.Engine) {

	g.Use(cors.Default())

	userGroup := g.Group("/api/users")
	{
		userGroup.POST("/", ctrl.CreateUser)
		userGroup.GET("/", ctrl.ListUsers)
		userGroup.DELETE("/:id", ctrl.DeleteUser)
		userGroup.GET("/:id", ctrl.GetUser)
		userGroup.PUT("/:id", ctrl.UpdateUser)
	}
}

func (ctrl *userController) CreateUser(c *gin.Context) {
	request := new(model.CreateUserRequest)
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
		return
	}

	c.JSON(http.StatusOK, model.NewJsonResponse(true))
}

func (ctrl *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", "Bad Request"))
		return
	}

	err := ctrl.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
		return
	}

	c.JSON(http.StatusOK, model.NewJsonResponse(true))
}

func (ctrl *userController) GetUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", "Bad Request"))
		return
	}

	data, err := ctrl.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
		return
	}

	c.JSON(http.StatusOK, model.NewJsonResponse(true).SetData(data))
}

func (ctrl *userController) ListUsers(c *gin.Context) {
	data, total, err := ctrl.service.ListUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
		return
	}

	c.JSON(http.StatusOK, model.NewJsonResponse(true).List(data, total))
}

func (ctrl *userController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, model.NewJsonResponse(false).SetError("400", "Bad Request"))
		return
	}

	request := new(model.UpdateUserRequest)
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.ID = id
	err := ctrl.service.UpdateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.(*model.JsonResponse))
		return
	}

	c.JSON(http.StatusOK, model.NewJsonResponse(true))
}
