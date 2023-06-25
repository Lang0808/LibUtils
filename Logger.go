package libutils

import (
	"os"

	"go.uber.org/zap"

	"fmt"

	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger(projectName string) {

	outputPath := fmt.Sprintf("/tmp/logs/%v/%v.log", projectName, projectName)
	errorOutputPath := fmt.Sprintf("/tmp/logs/%v/%v_error.log", projectName, projectName)
	if !IsExist(outputPath) {
		err := CreateFile(outputPath)
		if err != nil {
			panic("Cannot initialize file for info logger")
		}
	}
	if !IsExist(errorOutputPath) {
		err := CreateFile(errorOutputPath)
		if err != nil {
			panic("Cannot initialize file for error logger")
		}
	}
	fileInfo, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("open info file error")
	}
	fmt.Printf("fileInfo: %v\n", fileInfo)

	fileError, err := os.OpenFile(errorOutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("open error file error")
	}

	encoderCfg := zapcore.EncoderConfig{
		MessageKey: "Message",
		TimeKey:    "Time",
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.Lock(fileInfo), lowPriority),
		zapcore.NewCore(encoder, zapcore.Lock(fileError), highPriority),
	)

	logger = zap.New(core)
}

func Error(package_name string, function_name string, message string) {
	defer logger.Sync()
	msg := "[%v; %v] %v"
	msg = fmt.Sprintf(msg, package_name, function_name, message)
	logger.Error(msg)
}

func Info(package_name string, function_name string, message string) {
	defer logger.Sync()
	msg := "[%v; %v] %v"
	msg = fmt.Sprintf(msg, package_name, function_name, message)
	logger.Info(msg)
}
