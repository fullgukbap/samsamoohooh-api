package main

import (
	"log"
	"samsamoohooh-mini-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-mini-api/internal/infra/config"
	"samsamoohooh-mini-api/internal/infra/logger"

	"go.uber.org/zap"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Panicf("logger 생성에 실패하였습니다: %v\n", err)
	}
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Panicf("logger.Sync()를 실행하던 중 에러가 발생했습니다: %v\n", err)
		}
	}()

	config, err := config.New("./toml")
	if err != nil {
		logger.Panic(
			"config 생성에 실패하였습니다",
			zap.Error(err),
		)
	}

	db, err := database.NewDatabase(config)
	if err != nil {
		logger.Panic(
			"database 생성에 실패하였습니다",
			zap.Error(err),
		)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		logger.Panic(
			"database migration을 완료하지 못했습니다",
			zap.Error(err),
		)
	}

	_ = db
}
