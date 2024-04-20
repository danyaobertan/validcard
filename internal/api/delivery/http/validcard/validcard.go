package validcard

import (
	"net/http"

	"github.com/danyaobertan/validcard/internal/api/domain"
	"github.com/danyaobertan/validcard/internal/api/services"
	"github.com/danyaobertan/validcard/pkg/errorops"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	CardNumber      string `json:"cardNumber"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
}

type ResponseBody struct {
	Error *Error `json:"error,omitempty"`
	Valid bool   `json:"valid"`
}

type Error struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type Handler struct {
	validCardService services.ValidCard
}

// ValidateCardInfo validates card information and returns validation results.
// @Summary Validate card information
// @Description Validates the card number, expiration month, and expiration year, and returns whether the card is valid or not based on various checks including empty fields, format validation, and the Luhn algorithm.
// @Tags CardValidation
// @Accept json
// @Produce json
// @Param requestBody body RequestBody true "Card Information" example({"cardNumber": "5346464212502892", "expirationMonth": 1, "expirationYear": 2025})
// @Success 200 {object} ResponseBody {"valid": true} "Response when the card information is valid, error field may be null if there are no errors."
// @Failure 400 {object} ResponseBody {"error": {"code": 400, "message": "invalid card number"}, "valid": false} "Response when there is an error in card validation."
// @Router /validate [post]
func (h Handler) ValidateCardInfo() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request RequestBody
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorops.NewError(
				http.StatusInternalServerError,
				"failed to bind request body",
				nil,
			))

			return
		}

		if err := request.validate(); err != nil {
			ctx.JSON(err.Code, err)

			return
		}

		valid, err := h.validCardService.IsValidCardInfo(domain.Card{
			CardNumber:      request.CardNumber,
			ExpirationMonth: request.ExpirationMonth,
			ExpirationYear:  request.ExpirationYear,
		})
		if err != nil {
			ctx.JSON(err.Code, err)
			return
		}

		ctx.JSON(http.StatusOK, ResponseBody{
			Valid: valid,
		})
	}
}

func NewHandler(service services.ValidCard) *Handler {
	return &Handler{
		validCardService: service,
	}
}
