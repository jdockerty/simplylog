package simplylog_test

import (
	"github.com/jdockerty/simplylog"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	logger *simplylog.Logger = simplylog.New()
)

func TestParsedTypeIsCorrect(t *testing.T) {
	assert := assert.New(t)

	logger.Format.SetType("JSON")
	assert.Equal("json", logger.Format.Type, "expected 'json', got %s", logger.Format.Type)

	logger.Format.SetType("Json")
	assert.Equal("json", logger.Format.Type)

	logger.Format.SetType("TEXT")
	assert.Equal("text", logger.Format.Type)
}

func TestParsedTypeTextWhenInvalid(t *testing.T) {
	logger.Format.SetType("Definitely not a type")
	assert.Equal(t, "text", logger.Format.Type)
}

func TestTimestampDefault(t *testing.T) {
	assert.Equal(t, "15:04:05 02/01/2006", logger.Format.Timestamp)
	t.Log("Timestamps defaults are equal.")
}

func TestTimestampChangedToUserDefined(t *testing.T) {
	logger.Format.SetTimestamp(time.Kitchen)
	assert.Equal(t, "3:04PM", logger.Format.Timestamp)

}
