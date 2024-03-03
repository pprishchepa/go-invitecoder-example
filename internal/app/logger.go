package app

import (
	"io"
	"log"
	"os"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/rs/zerolog"
)

func newLogger(conf config.Config) (zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(conf.Log.Level)
	if err != nil {
		return zerolog.Logger{}, err
	}

	writer := io.Writer(os.Stdout)
	if conf.Log.Pretty {
		writer = zerolog.NewConsoleWriter()
	}

	logger := zerolog.New(writer).Level(level)

	log.SetFlags(0)
	log.SetOutput(logger)

	return logger.With().Timestamp().Logger(), nil
}
