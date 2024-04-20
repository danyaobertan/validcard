package validcard_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	vchttp "github.com/danyaobertan/validcard/internal/api/delivery/http/validcard"
	"github.com/danyaobertan/validcard/internal/api/services/validcard"
	"github.com/gin-gonic/gin"
)

// TestValidateCardInfo using the real service.
func TestValidateCardInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	realCardService := validcard.NewService()
	router.POST("/validate-card", vchttp.NewHandler(realCardService).ValidateCardInfo())

	tests := []struct { //nolint:govet // it's a test
		name           string
		body           vchttp.RequestBody
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "empty card number",
			body: vchttp.RequestBody{
				CardNumber:      "",
				ExpirationMonth: 1,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 400,
			expectedBody:   `{"code":400,"message":"card number is empty"}`,
		},
		{
			name: "empty expiration month",
			body: vchttp.RequestBody{
				CardNumber:      "5335724310512959",
				ExpirationMonth: 0,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 400,
			expectedBody:   `{"code":400,"message":"expiration month is empty"}`,
		},
		{
			name: "empty expiration year",
			body: vchttp.RequestBody{
				CardNumber:      "5335724310512959",
				ExpirationMonth: 1,
				ExpirationYear:  0,
			},
			expectedStatus: 400,
			expectedBody:   `{"code":400,"message":"expiration year is empty"}`,
		},
		{
			name: "valid card",
			body: vchttp.RequestBody{
				CardNumber:      "5346464212502892",
				ExpirationMonth: 1,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 200,
			expectedBody:   `{"valid":true}`,
		},
		{
			name: "card number with spaces",
			body: vchttp.RequestBody{
				CardNumber:      "5346 4642 1250 2892",
				ExpirationMonth: 1,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 200,
			expectedBody:   `{"valid":true}`,
		},
		{
			name: "invalid symbol in card number",
			body: vchttp.RequestBody{
				CardNumber:      "534i 4642 1250 2892",
				ExpirationMonth: 1,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 400,
			expectedBody:   `{"code":400,"message":"invalid card number symbols","description":"card number could contain only digits and spaces, but contains invalid character: i"}`,
		},
		{
			name: "invalid card by luhn",
			body: vchttp.RequestBody{
				CardNumber:      "5346464212502893",
				ExpirationMonth: 1,
				ExpirationYear:  time.Now().Year() + 1,
			},
			expectedStatus: 400,
			expectedBody:   `{"code":400,"message":"invalid card number by Luhn algorithm"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			bodyBytes, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/validate-card", bytes.NewBuffer(bodyBytes))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != tc.expectedStatus {
				t.Errorf("expected status %d; got %d", tc.expectedStatus, resp.Code)
			}

			if resp.Body.String() != tc.expectedBody {
				t.Errorf("expected body %s; got %s", tc.expectedBody, resp.Body.String())
			}
		})
	}
}
