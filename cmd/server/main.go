package main

import (
	"log"
	"meeting-scheduler/internal/handler"
	"meeting-scheduler/internal/repository"
	"meeting-scheduler/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eventRepo := repository.NewInMemoryEventRepository()
	availabilityRepo := repository.NewInMemoryAvailabilityRepository()
	userRepo := repository.NewInMemoryUserRepository()
	svc := service.NewSchedulerService(userRepo, eventRepo, availabilityRepo)
	h := handler.NewHandler(svc)

	h.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
