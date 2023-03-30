package infrastructure

import (
	"time"

	"github.com/jpillora/backoff"
	"github.com/krobus00/krokit/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	DB           *gorm.DB
	StopTickerCh chan bool
)

type DBConfig struct {
	DSN                string
	MaxIdleConns       int
	MaxOpenConns       int
	ConnMaxLifetime    time.Duration
	PingInternal       time.Duration
	LogLevel           string
	ReconnectFactor    float64
	RetryAttempts      float64
	ReconnectMinJitter time.Duration
	ReconnectMaxJitter time.Duration
}

func InitializeDBConn(config *DBConfig) {
	conn, err := openDBConn(config)
	if err != nil {
		logrus.WithField("databaseDSN", config.DSN).Fatal("failed to connect  database: ", err.Error())
	}

	DB = conn
	StopTickerCh = make(chan bool)

	go checkConnection(config, time.NewTicker(config.PingInternal))

	switch config.LogLevel {
	case "error":
		DB.Logger = DB.Logger.LogMode(gormLogger.Error)
	case "warn":
		DB.Logger = DB.Logger.LogMode(gormLogger.Warn)
	default:
		DB.Logger = DB.Logger.LogMode(gormLogger.Info)
	}

	logrus.Info("Connection to database Server success...")
}

func checkConnection(config *DBConfig, ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := DB.DB(); err != nil {
				reconnectDBConn(config)
			}
		}
	}
}

func reconnectDBConn(config *DBConfig) {
	b := backoff.Backoff{
		Factor: config.ReconnectFactor,
		Jitter: true,
		Min:    config.ReconnectMinJitter,
		Max:    config.ReconnectMaxJitter,
	}

	dbRetryAttempts := config.RetryAttempts

	for b.Attempt() < dbRetryAttempts {
		conn, err := openDBConn(config)
		if err != nil {
			logrus.WithField("databaseDSN", config.DSN).Error("failed to connect database: ", err.Error())
		}

		if conn != nil {
			DB = conn
			break
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= dbRetryAttempts {
		logrus.Fatal("maximum retry to connect database")
	}
	b.Reset()
}

func openDBConn(config *DBConfig) (*gorm.DB, error) {
	psqlDialector := postgres.Open(config.DSN)
	db, err := gorm.Open(psqlDialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	utils.ContinueOrFatal(err)

	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(config.ConnMaxLifetime)

	return db, nil
}
