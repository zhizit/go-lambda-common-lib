package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 指定されたキーでログを設定し、zap.Loggerを返します
// log.Fatal を使用すると、即座にプログラムが終了し、defer で登録したリソース解放処理が実行されない場合があるので使用しないでください
// エラー発生時には早期リターンで処理を中断し、クリーンアップ処理を確実に行ってください

// ///////////////// 使い方 //////////////////
// logger, err := createLogger("sample-key")
//
//	if err != nil {
//		return err
//	}
//
// defer logger.Sync()
//
// logger.Info("Info log!")
// logger.Error("Error log!")
// /////////////////////////////////////////
func CreateLogger(key string) (*zap.Logger, error) {
	conf := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			TimeKey:     "time",
			MessageKey:  "msg",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
	}

	logger, err := conf.Build()
	if err != nil {
		return nil, err
	}

	return logger.With(zap.String("key", key)), nil
}
