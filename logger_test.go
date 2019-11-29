package logger_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/onskycloud/logger"
)

func TestContainer(t *testing.T) {
	logLevel := 0
	logLocation := ""
	logBucket := ""

	testLog(t, logLevel, logLocation, logBucket)

}
func testLog(t *testing.T, logLevel int, logLocation string, logBucket string) {
	logger, file, err := logger.Init(logLevel, logLocation, logBucket)
	if err != nil || file == nil || logger == nil {
		fmt.Printf("err:%+v\n", err)
		t.Fatal("log fail")
	}
	err = errors.New("some error")
	logger.Warn().Err(err).Msg("test warn")
	logger.Debug().Err(err).Msg("test debug")
	logger.Err(err).Msg("test error")
	t.Log("ok")
}
