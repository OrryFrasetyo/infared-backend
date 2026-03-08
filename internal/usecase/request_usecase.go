package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"infared-backend/internal/domain"
	"infared-backend/internal/repository"
	"infared-backend/pkg/gemini"
	"infared-backend/pkg/utils"
	"strings"
	"time"
)

type RequestUsecase interface {
	ProcessChat(ctx context.Context, userID, poskoID, promptText string) (*domain.LogisticsRequest, error)
	GetAllRequests(ctx context.Context) ([]domain.RequestDetail, error)
}

type requestUsecase struct {
	requestRepo repository.RequestRepository
	itemRepo    repository.ItemRepository
	aiClient    gemini.GeminiClient
}

func NewRequestUsecase(reqRepo repository.RequestRepository, itemRepo repository.ItemRepository, aiClient gemini.GeminiClient) RequestUsecase {
	return &requestUsecase{
		requestRepo: reqRepo,
		itemRepo:    itemRepo,
		aiClient:    aiClient,
	}
}

type AIResponse struct {
	Items []struct {
		ItemID   string `json:"item_id"`
		Quantity int    `json:"quantity"`
		Urgency  string `json:"urgency"`
	} `json:"items"`
}

func (u *requestUsecase) ProcessChat(ctx context.Context, userID, poskoID, promptText string) (*domain.LogisticsRequest, error) {
	masterItems, err := u.itemRepo.GetAll(ctx)
	if err != nil || len(masterItems) == 0 {
		return nil, errors.New("master barang kosong, tidak bisa mencocokkan data")
	}

	var catalogBuilder strings.Builder
	for _, itm := range masterItems {
		catalogBuilder.WriteString(fmt.Sprintf("- ID: %s | Nama: %s | Satuan: %s\n", itm.ID, itm.Name, itm.Unit))
	}

	systemPrompt := fmt.Sprintf(`Anda adalah asisten logistik kebencanaan InfaRed. 
		Tugas Anda adalah mengekstrak permintaan barang dari teks relawan.
		Berikut adalah KATALOG BARANG yang tersedia di sistem kami:
		%s

		Instruksi:
		1. Cocokkan barang yang diminta relawan dengan KATALOG BARANG di atas.
		2. Jika barang yang diminta mirip (misal: "air putih" -> "Air Mineral"), gunakan item_id dari katalog.
		3. Tentukan tingkat urgensi: "rendah", "sedang", "tinggi", atau "kritis" berdasarkan nada pesan.
		4. KEMBALIKAN HANYA FORMAT JSON SEPERTI INI, TANPA MARKDOWN ATAU TEKS LAIN:
		{"items": [{"item_id": "...", "quantity": 0, "urgency": "..."}]}
		Jika ada barang yang diminta tapi tidak ada di katalog, abaikan saja barang tersebut.`, catalogBuilder.String())

	aiResultJSON, err := u.aiClient.ExtractLogisticsData(ctx, systemPrompt, promptText)
	if err != nil {
		return nil, fmt.Errorf("gagal memproses AI: %v", err)
	}

	var aiData AIResponse
	if err := json.Unmarshal([]byte(aiResultJSON), &aiData); err != nil {
		return nil, fmt.Errorf("gagal membaca format output AI: %v", err)
	}

	requestID := utils.GenerateID("req")

	newRequest := &domain.LogisticsRequest{
		ID:             requestID,
		PoskoID:        poskoID,
		RequestedBy:    userID,
		OriginalPrompt: &promptText,
		Status:         domain.StatusPending,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	var requestItems []domain.RequestItem
	for _, aiItem := range aiData.Items {
		requestItems = append(requestItems, domain.RequestItem{
			ID:        utils.GenerateID("rit"),
			RequestID: requestID,
			ItemID:    aiItem.ItemID,
			Quantity:  aiItem.Quantity,
			Urgency:   domain.UrgencyLevel(aiItem.Urgency),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	if len(requestItems) == 0 {
		return nil, errors.New("tidak ada barang yang cocok dengan katalog kami")
	}

	if err := u.requestRepo.CreateRequestWithItems(ctx, newRequest, requestItems); err != nil {
		return nil, fmt.Errorf("gagal menyimpan ke database: %v", err)
	}

	return newRequest, nil
}

func (u *requestUsecase) GetAllRequests(ctx context.Context) ([]domain.RequestDetail, error) {
	return u.requestRepo.GetAllWithDetails(ctx)
}
