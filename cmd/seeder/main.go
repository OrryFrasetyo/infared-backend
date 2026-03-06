package main

import (
	"context"
	"infared-backend/internal/config"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
	"infared-backend/pkg/utils"
	"log"
	"time"
)

func main() {
	log.Println("🌱 Memulai proses Seeding Database InfaRed...")

	db := config.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	ctx := context.Background()

	admins := []struct {
		Name  string
		Email string
	}{
		{"Admin BPBD Padang", "admin_padang@infared.com"},
		{"Admin BNPB Pusat", "admin_pusat@infared.com"},
	}

	hashedPassword, err := utils.HashPassword("password123")
	if err != nil {
		log.Fatalf("Gagal nge-hash password: %v", err)
	}

	for _, a := range admins {
		user := &domain.User{
			ID:           utils.GenerateID("usr"),
			Name:         a.Name,
			Email:        a.Email,
			PasswordHash: hashedPassword,
			Role:         domain.RoleAdmin, 
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		err := userRepo.Create(ctx, user)
		if err != nil {
			log.Printf("⚠️ Melewati %s (Mungkin email sudah ada/duplicate): %v\n", a.Email, err)
		} else {
			log.Printf("✅ Berhasil membuat akun Admin: %s\n", a.Email)
		}
	}

	log.Println("🎉 Proses Seeding selesai!")
}
