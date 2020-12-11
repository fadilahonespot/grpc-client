package main

import (
	"log"

	"grpc-client/model"
	userHandler "grpc-client/user/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)


func main() {
	port := "8827"
	targetPort := "8957"
	conn, err := grpc.Dial(":" + targetPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cloud not conect to %v %v", targetPort, err)
	}
	user := model.NewUsersClient(conn)
	router := gin.Default()

	userHandler.CreateUserHandler(router, user)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}