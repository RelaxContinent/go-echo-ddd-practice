package framework

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger   *zap.Logger
	logLevel log.Lvl
}

func NewLogger() (echo.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &zapLogger{
		logger:   logger,
		logLevel: log.INFO,
	}, nil
}

func (z *zapLogger) Output() io.Writer {
	return os.Stdout
}

func (z *zapLogger) SetOutput(_ io.Writer) {}

func (z *zapLogger) Prefix() string {
	return ""
}

func (z *zapLogger) SetPrefix(_ string) {}

func (z *zapLogger) Level() log.Lvl {
	return z.logLevel
}

func (z *zapLogger) SetLevel(lvl log.Lvl) {
	z.logLevel = lvl
}

func (z *zapLogger) SetHeader(_ string) {}

func (z *zapLogger) Print(i ...interface{}) {
	z.logger.Sugar().Info(i...)
}

func (z *zapLogger) Printf(format string, args ...interface{}) {
	z.logger.Sugar().Infof(format, args...)
}

func (z *zapLogger) Printj(j log.JSON) {
	z.logger.Sugar().Infow("print message", "data", j)
}

func (z *zapLogger) Debug(i ...interface{}) {
	z.logger.Sugar().Debug(i...)
}

func (z *zapLogger) Debugf(format string, args ...interface{}) {
	z.logger.Sugar().Debugf(format, args...)
}

func (z *zapLogger) Debugj(j log.JSON) {
	z.logger.Sugar().Debugw("debug message", "data", j)
}

func (z *zapLogger) Info(i ...interface{}) {
	z.logger.Sugar().Info(i...)
}

func (z *zapLogger) Infof(format string, args ...interface{}) {
	z.logger.Sugar().Infof(format, args...)
}

func (z *zapLogger) Infoj(j log.JSON) {
	z.logger.Sugar().Infow("info message", "data", j)
}

func (z *zapLogger) Warn(i ...interface{}) {
	z.logger.Sugar().Warn(i...)
}

func (z *zapLogger) Warnf(format string, args ...interface{}) {
	z.logger.Sugar().Warnf(format, args...)
}

func (z *zapLogger) Warnj(j log.JSON) {
	z.logger.Sugar().Warnw("warn message", "data", j)
}

func (z *zapLogger) Error(i ...interface{}) {
	z.logger.Sugar().Error(i...)
}

func (z *zapLogger) Errorf(format string, args ...interface{}) {
	z.logger.Sugar().Errorf(format, args...)
}

func (z *zapLogger) Errorj(j log.JSON) {
	z.logger.Sugar().Errorw("error message", "data", j)
}

func (z *zapLogger) Fatal(i ...interface{}) {
	z.logger.Sugar().Fatal(i...)
}

func (z *zapLogger) Fatalf(format string, args ...interface{}) {
	z.logger.Sugar().Fatalf(format, args...)
}

func (z *zapLogger) Fatalj(j log.JSON) {
	z.logger.Sugar().Fatalw("fatal message", "data", j)
}

func (z *zapLogger) Panic(i ...interface{}) {
	z.logger.Sugar().Panic(i...)
}

func (z *zapLogger) Panicf(format string, args ...interface{}) {
	z.logger.Sugar().Panicf(format, args...)
}

func (z *zapLogger) Panicj(j log.JSON) {
	z.logger.Sugar().Panicw("panic message", "data", j)
}
