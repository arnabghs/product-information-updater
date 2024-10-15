package priceUpdatecontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	mocks2 "product-information-updater/app/updatePrice/service/mocks"
	"testing"
)

func Test_UpdatePrice(t *testing.T) {
	validRequest := updatePriceModel.RequestBody{
		ID:      "12",
		Message: "Test",
	}

	testcases := []struct {
		name               string
		isSvcCalled        bool
		svcResponse        error
		productID          string
		reqBody            interface{}
		expectStatusCode   int
		expectResponseBody string
	}{
		{
			name:               "successfully processed request",
			isSvcCalled:        true,
			svcResponse:        nil,
			productID:          "product123",
			reqBody:            validRequest,
			expectStatusCode:   http.StatusOK,
			expectResponseBody: `{"status":"success"}`,
		},
		{
			name:             "throw Bad Request if product ID not found",
			isSvcCalled:      false,
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name:             "throw Bad Request if request body not valid",
			isSvcCalled:      false,
			productID:        "product123",
			reqBody:          "invalid",
			expectStatusCode: http.StatusBadRequest,
		},
		{
			name:             "throw internal server in case of any processing error",
			isSvcCalled:      true,
			svcResponse:      errors.New("some error"),
			productID:        "product123",
			reqBody:          validRequest,
			expectStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testcases {
		mockCtrl := gomock.NewController(t)
		mockSvc := mocks2.NewMockService(mockCtrl)

		recorder := httptest.NewRecorder()
		gCtx, _ := gin.CreateTestContext(recorder)

		mr, _ := json.Marshal(tc.reqBody)
		gCtx.Request = httptest.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(mr))

		t.Run(tc.name, func(t *testing.T) {
			if tc.productID != "" {
				gCtx.Params = []gin.Param{
					{
						Key:   "productID",
						Value: tc.productID,
					},
				}
			}

			if tc.isSvcCalled {
				mockSvc.EXPECT().Process(gCtx, "product123", tc.reqBody).Return(tc.svcResponse)
			}

			priceUpdateHandler := NewPriceUpdateHandler(mockSvc)
			priceUpdateHandler.UpdatePrice(gCtx)

			assert.Equal(t, tc.expectStatusCode, recorder.Code)
			if tc.expectResponseBody != "" {
				assert.Equal(t, tc.expectResponseBody, recorder.Body.String())
			}
		})
	}
}
