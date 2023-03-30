package utils

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	logrus.Info(fmt.Sprintf("%s took %s", name, elapsed))
}

func ParseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}
