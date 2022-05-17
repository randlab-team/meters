package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/randlab-team/meters/handlers"
	"github.com/randlab-team/meters/mocks"
	"github.com/randlab-team/meters/models"
)

func TestMeters_GetAll_ReturnMany(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockedMetersRepo := mock_repository.NewMockMeters(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/meters", nil)
	rec := httptest.NewRecorder()
	echoCtx := e.NewContext(req, rec)

	// given
	var meterLogs []models.MeterLog
	if err := faker.FakeData(&meterLogs); err != nil {
		t.Fail()
	}

	expectedMeterLogsJson, err := json.Marshal(meterLogs)
	if err != nil {
		t.Fail()
	}

	mockedMetersRepo.EXPECT().GetAll().Return(
		meterLogs,
		nil,
	)

	metersHandlers := handlers.NewMeters(mockedMetersRepo)

	// when
	err = metersHandlers.GetAll(echoCtx)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expectedMeterLogsJson)+"\n", rec.Body.String())
}

func TestMeters_GetAll_ReturnOne(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockedMetersRepo := mock_repository.NewMockMeters(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/meters", nil)
	rec := httptest.NewRecorder()
	echoCtx := e.NewContext(req, rec)

	// given
	var meterLogs []models.MeterLog
	if err := faker.FakeData(&meterLogs); err != nil {
		t.Fail()
	}

	meterLogs = meterLogs[:1]

	expectedMeterLogsJson, err := json.Marshal(meterLogs)
	if err != nil {
		t.Fail()
	}

	mockedMetersRepo.EXPECT().GetAll().Return(
		meterLogs,
		nil,
	)

	metersHandlers := handlers.NewMeters(mockedMetersRepo)

	// when
	err = metersHandlers.GetAll(echoCtx)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expectedMeterLogsJson)+"\n", rec.Body.String())
}

func TestMeters_GetAll_ReturnZero(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockedMetersRepo := mock_repository.NewMockMeters(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/meters", nil)
	rec := httptest.NewRecorder()
	echoCtx := e.NewContext(req, rec)

	// given
	mockedMetersRepo.EXPECT().GetAll().Return(
		[]models.MeterLog{},
		nil,
	)

	metersHandlers := handlers.NewMeters(mockedMetersRepo)

	// when
	err := metersHandlers.GetAll(echoCtx)

	// then
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

func TestMeters_GetAll_ReturnError(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockedMetersRepo := mock_repository.NewMockMeters(mockCtl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/v1/meters", nil)
	rec := httptest.NewRecorder()
	echoCtx := e.NewContext(req, rec)

	mockedMetersRepo.EXPECT().GetAll().Return(
		[]models.MeterLog{},
		errors.New("TEST_ERROR"),
	)

	metersHandlers := handlers.NewMeters(mockedMetersRepo)

	// when
	err := metersHandlers.GetAll(echoCtx)

	// then
	assert.Errorf(t, err, "TEST_ERROR")
}
