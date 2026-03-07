package main

import (
	"infared-backend/internal/config"
	"infared-backend/internal/delivery/http/handler"
	"infared-backend/internal/delivery/http/router"
	"infared-backend/internal/repository"
	"infared-backend/internal/usecase"
	"infared-backend/pkg/gemini"
	"log"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	itemRepo := repository.NewItemRepository(db)
	itemUsecase := usecase.NewItemUsecase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUsecase)

	aiClient, err := gemini.NewGeminiClient()
	if err != nil {
		log.Fatalf("Gagal inisiasi Gemini: %v", err)
	}

	requestRepo := repository.NewRequestRepository(db)
	requestUsecase := usecase.NewRequestUsecase(requestRepo, itemRepo, aiClient)
	requestHandler := handler.NewRequestHandler(requestUsecase)

	r := router.SetupRouter(userHandler, itemHandler, requestHandler)

	log.Println("🚀 Menjalankan server InfaRed di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server berhenti paksa: %v", err)
	}
}
