package applogger

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func Init(dev bool) {
	var lg *zap.Logger
	if dev {
		lg, _ = zap.NewDevelopment()
	} else {
		lg, _ = zap.NewProduction()
	}
	logger = lg.Sugar()
}

func Close() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}

func Error(msg string, err error, fields ...zap.Field) {
	params := make([]interface{}, len(fields)+1)
	for i, fld := range fields {
		params[i] = fld
	}
	params[len(params)-1] = zap.Error(err)
	logger.Errorw(
		msg,
		params...,
	)
}
