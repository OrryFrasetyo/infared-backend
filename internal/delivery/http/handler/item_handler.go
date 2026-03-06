package handler

import (
	"infared-backend/internal/usecase"
	"infared-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemUsecase usecase.ItemUsecase
}

func NewItemHandler(u usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{itemUsecase: u}
}

type CreateItemRequest struct {
	Name string `json:"name" binding:"required"`
	Unit string `json:"unit" binding:"required"` 
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var req CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Format input tidak valid", err.Error())
		return
	}

	item, err := h.itemUsecase.CreateItem(c.Request.Context(), req.Name, req.Unit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal menambahkan barang", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Barang berhasil ditambahkan", item)
}

func (h *ItemHandler) GetAllItems(c *gin.Context) {
	items, err := h.itemUsecase.GetAllItems(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data barang", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Daftar barang berhasil diambil", items)
}
