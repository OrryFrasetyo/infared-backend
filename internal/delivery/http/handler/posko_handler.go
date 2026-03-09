package handler

import (
	"infared-backend/internal/usecase"
	"infared-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PoskoHandler struct {
	poskoUsecase usecase.PoskoUsecase
}

func NewPoskoHandler(u usecase.PoskoUsecase) *PoskoHandler {
	return &PoskoHandler{poskoUsecase: u}
}

func (h *PoskoHandler) GetAllPoskos(c *gin.Context) {
	poskos, err := h.poskoUsecase.GetAllPoskos(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil daftar posko", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil daftar posko", poskos)
}
