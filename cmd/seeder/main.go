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

	log.Println("🏕️ Memulai proses Seeding Posko...")

	address1 := "Jl. Sudirman, Padang Barat (Balaikota Lama)"
	address2 := "Kampus Politeknik Negeri Padang, Limau Manis"

	poskos := []domain.Posko{
		{
			ID:        "psk-12345", 
			Name:      "Posko Utama Balaikota Padang",
			Address:   &address1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "psk-67890", 
			Name:      "Posko Darurat PNP",
			Address:   &address2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	queryPosko := `
		INSERT INTO posko (id, name, address, latitude, longitude, coordinator_id, created_at, updated_at)
		VALUES (:id, :name, :address, :latitude, :longitude, :coordinator_id, :created_at, :updated_at)
		ON CONFLICT (id) DO NOTHING
	`

	for _, p := range poskos {
		_, err := db.NamedExecContext(ctx, queryPosko, &p)
		if err != nil {
			log.Printf("⚠️ Gagal insert posko %s: %v\n", p.Name, err)
		} else {
			log.Printf("✅ Berhasil membuat Posko: %s (ID: %s)\n", p.Name, p.ID)
		}
	}

	log.Println("🎉 Proses Seeding selesai!")
}
