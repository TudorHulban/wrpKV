package badgerwrap

import (
	"log"
	"os"
	"sync"
)

// Custom logger with info and debug. Although not necessary besides l, keeping info for being available. To add log level.
type customLog struct {
	l         *log.Logger
	spoolTo   string
	spoolJSON bool // if we want JSON format
}

var once sync.Once

func NewLogger(pAppName, pSpoolTo string, pSpoolJSON bool) (logger, error) {
	result := new(customLog)
	var errCreate error

	once.Do(func() {
		result, errCreate = createFileLogger(pAppName, pSpoolTo, pSpoolJSON)
	})
	return result, errCreate
}

func createFileLogger(pAppName, pSpoolTo string, pSpoolJSON bool) (*customLog, error) {
	osFile, errFileOpen := os.OpenFile(pSpoolTo, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if errFileOpen != nil {
		return nil, errFileOpen
	}

	result := new(customLog)
	result.l = log.New(osFile, pAppName+" ", log.Llongfile)
	result.spoolTo = pSpoolTo
	result.spoolJSON = pSpoolJSON

	return result, nil
}

// not using pointer receiver as type contains pointer already.
func (l customLog) Debugf(format string, args ...interface{}) {
	l.l.Printf(format, args...)
}
