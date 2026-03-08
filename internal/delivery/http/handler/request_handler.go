package handler

import (
	"infared-backend/internal/usecase"
	"infared-backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	requestUsecase usecase.RequestUsecase
}

func NewRequestHandler(u usecase.RequestUsecase) *RequestHandler {
	return &RequestHandler{requestUsecase: u}
}

type ChatRequest struct {
	PoskoID    string `json:"posko_id" binding:"required"`
	PromptText string `json:"prompt_text" binding:"required"`
}

func (h *RequestHandler) ChatToAI(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Format input tidak valid", err.Error())
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Sesi tidak valid", nil)
		return
	}

	result, err := h.requestUsecase.ProcessChat(c.Request.Context(), userID.(string), req.PoskoID, req.PromptText)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal memproses laporan", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Permintaan logistik berhasil diproses oleh AI", result)
}

func (h *RequestHandler) GetAllRequests(c *gin.Context) {
	requests, err := h.requestUsecase.GetAllRequests(c.Request.Context())
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data laporan logistik", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Berhasil mengambil data laporan logistik", requests)
}
