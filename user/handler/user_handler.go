package handle

import (
	"context"
	"grpc-client/model"
	"grpc-client/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)


type UserHandler struct{
	userClient model.UsersClient
}

func CreateUserHandler(r *gin.Engine, userClient model.UsersClient) {
	userHandler := &UserHandler{userClient}

	r.POST("/user", userHandler.addUser)
	r.GET("/user", userHandler.viewUser)
	r.GET("/user/:id", userHandler.viewUserById)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.deleteUser)
}

func (e *UserHandler) addUser(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user.Name == "" || user.Email == "" || user.Alamat == "" || user.Password == "" {
		utils.HandleError(c, http.StatusBadRequest, "fields are required")
		return
	}
	_, err = e.userClient.InsertUser(context.Background(), &user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, "Sucess Add Data")
}

func (e *UserHandler) viewUser(c *gin.Context) {
	userList, err := e.userClient.GetUserList(context.Background(), new(empty.Empty))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, userList.List)
}

func (e *UserHandler) viewUserById(c *gin.Context) {
	id := c.Param("id")
	userid := model.UserId{
		Id: id,
	}
	user, err := e.userClient.GetUserById(context.Background(), &userid)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSuccess(c, user)
}

func (e *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var up = model.UserUpdate{
		Id: id,
		User: &user,
	}
	_, err = e.userClient.UpdateUser(context.Background(), &up)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, "Update User Success")
}

func(e *UserHandler) deleteUser(c *gin.Context) {
	id := c.Param("id")
	var up = model.UserId{
		Id: id,
	}
	_, err := e.userClient.DeleteUser(context.Background(), &up)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSuccess(c, "Success delete data")
}