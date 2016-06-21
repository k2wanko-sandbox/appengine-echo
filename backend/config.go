package backend

import "time"

var config AppConfig

type AppConfig struct {
	Loc *time.Location
}

func loadConfig() (err error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return
	}
	time.Local, config.Loc = loc, loc
	return
}
