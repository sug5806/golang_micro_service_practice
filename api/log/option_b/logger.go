package option_b

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
	"os"
)

var (
	Log *zap.Logger
)

func newWinFileSink(u *url.URL) (zap.Sink, error) {
	// Remove leading slash left by url.Parse()
	return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func init() {
	//_ = zap.RegisterSink("winfile", newWinFileSink)
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Log, err = logConfig.Build()

	if err != nil {
		Error(err.Error(), err)
		panic(err)
	}
}

func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Info(msg string, tags ...zap.Field) {
	Log.Info(msg, tags...)
	Log.Sync()

}

func Debug(msg string, tags ...zap.Field) {
	Log.Debug(msg, tags...)
	Log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	Log.Error(msg, tags...)
	Log.Sync()
}
