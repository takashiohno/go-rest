package logs

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogger *zap.Logger
func init() {

    // configration
    config := zap.NewProductionConfig()
    // iso8601 format
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    enc := zapcore.NewJSONEncoder(config.EncoderConfig)
    sink := zapcore.AddSync(
        // log lotate by lumberjack
        &lumberjack.Logger{
            Filename: "zaplogger.log",
            MaxSize: 500,   // megabytes
        },
    )
    zapLogger := zap.New(
        zapcore.NewCore(enc, sink, config.Level),
    )
    defer zapLogger.Sync()

    // global setting
    undo := zap.ReplaceGlobals(zapLogger)
    defer undo()
    ZapLogger = zap.L()
}
