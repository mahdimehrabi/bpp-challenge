package infrastractures

import (
	"log"
	"os"
)

type PasargadLogger struct {
	LG *log.Logger
}

func (l *PasargadLogger) Error(err string) {
	l.LG.Print(err)
}

func NewLogger() PasargadLogger {
	lg := log.New(os.Stdout, "pasargad ", log.LstdFlags)
	return PasargadLogger{LG: lg}
}
