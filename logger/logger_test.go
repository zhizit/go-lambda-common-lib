package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCreateLogger(t *testing.T) {
	// ロガーを作成
	logger, err := CreateLogger("test-key")

	// エラーチェック
	assert.NoError(t, err)
	assert.NotNil(t, logger)

	// メモリバッファにログをキャプチャするためのエンコーダを設定
	var buf bytes.Buffer
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		LevelKey:    "level",
		TimeKey:     "time",
		MessageKey:  "msg",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})
	core := zapcore.NewCore(encoder, zapcore.AddSync(&buf), zapcore.InfoLevel)

	// 新しいロガーを作成
	testLogger := zap.New(core)

	// ログメッセージを記録
	testLogger.Info("Test message", zap.String("key", "test-key"))

	// 出力されたログ内容を取得
	logOutput := buf.String()

	// ログ内容が含まれているか確認
	assert.Contains(t, logOutput, `"msg":"Test message"`)
	assert.Contains(t, logOutput, `"key":"test-key"`)
	assert.Contains(t, logOutput, `"level":"info"`)
}
