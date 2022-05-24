package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/randlab-team/meters/repository"
)

type Meters struct {
	metersRepo repository.Meters
}

func NewMeters(metersRepo repository.Meters) *Meters {
	return &Meters{metersRepo: metersRepo}
}

func (m *Meters) GetAll(c echo.Context) error {
	meterLogs, err := m.metersRepo.GetAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch all meter logs")

		return errors.Wrap(err, "failed to fetch all meter logs")
	}

	if len(meterLogs) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, meterLogs)
}
