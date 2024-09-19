package logger

import "go.uber.org/zap"

type Logger struct {
	zapLogger *zap.Logger
	*zap.SugaredLogger
}

func (l *Logger) Sync() error {
	return l.zapLogger.Sync()
}

func NewLogger() (*Logger, error) {
	// TODO: file에 저장하게끔 설정
	// TODO: lumberjack 이용해서 파일 관리하면 될듯 (https://github.com/natefinch/lumberjack)
	// TODO: configuration을 설정하여 보기 좋게 변경
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	l := Logger{
		zapLogger:     zapLogger,
		SugaredLogger: zapLogger.Sugar(),
	}

	return &l, nil
}
