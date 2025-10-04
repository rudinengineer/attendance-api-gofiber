package utils

import (
	"absensi-api/internal/config"
	"time"
)

func TimeNow() (time.Time, error) {
	configuration := config.Load()
	loc, err := time.LoadLocation(configuration.Server.Timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().In(loc), nil
}
