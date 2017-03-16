package logging

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
)

func TestNewLoggerShouldInstantiateNewLogger(t *testing.T) {
	logger := NewLogger()

	assert.Equal(t, 3, logger.LogLevel, fmt.Sprintf("Expected default LogLevel to be %d", LOG_LEVEL_INFO))
	assert.Equal(t, os.Stderr, logger.DebugFile, "Expected default DebugFile to be Stderr", LOG_LEVEL_INFO)
	assert.Equal(t, os.Stderr, logger.ErrorFile, "Expected default ErrorFile to be Stderr", LOG_LEVEL_INFO)
	assert.Equal(t, os.Stdout, logger.InfoFile, "Expected default InfoFile to be Stderr", LOG_LEVEL_INFO)
	assert.Equal(t, os.Stderr, logger.WarnFile, "Expected default WarnFile to be Stderr", LOG_LEVEL_INFO)
}

func TestNewLoggerShouldHaveCallableLogMethods(t *testing.T) {
	logger := NewLogger()
	logger.LogLevel = LOG_LEVEL_DEBUG
	logger.Debug("foo")
	logger.Debugf("debug: %s", "foo")
	logger.Error(1)
	logger.Errorf("error: %d", 1)
	logger.Info("hello")
	logger.Infof("info: %s", "hello")
	logger.Warn(true)
	logger.Warnf("warn: %t", true)
}

func TestNewLoggerShouldBeSilenceable(t *testing.T) {
	logger := NewLogger()
	logger.LogLevel = LOG_LEVEL_NONE
	logger.Debug("foo")
	logger.Debugf("debug: %s", "foo")
	logger.Error(1)
	logger.Errorf("error: %d", 1)
	logger.Info("hello")
	logger.Infof("info: %s", "hello")
	logger.Warn(true)
	logger.Warnf("warn: %t", true)
}
