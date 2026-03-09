package handler

import (
	"infared-backend/internal/usecase"
	"infared-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	inventoryUsecase usecase.InventoryUsecase
}

func NewInventoryHandler(u usecase.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{inventoryUsecase: u}
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	poskoID := c.Param("id")

	if poskoID == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID Posko tidak valid", nil)
		return
	}

	inventory, err := h.inventoryUsecase.GetInventoryByPosko(c.Request.Context(), poskoID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data inventaris", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil data inventaris posko", inventory)
}
