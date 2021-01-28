package simplylog_test

import (
	"io"
	"os"
	"testing"

	"github.com/jdockerty/simplylog"
	"github.com/stretchr/testify/assert"
)

type MockFile os.File

func (MockFile) Open() (*os.File, error) {
	return &os.File{}, nil
}

func TestNewAreDefaults(t *testing.T) {
	assert := assert.New(t)
	logger := simplylog.New()

	assert.IsType(os.Stderr, logger.Out)
	assert.False(logger.Verbose)
}

func TestSetOutputChanged(t *testing.T) {
	assert := assert.New(t)
	var m MockFile
	file, err := m.Open()
	assert.Nil(err)

	logger := simplylog.New()
	logger.SetOutput(file)

	assert.Implements((*io.Writer)(nil), logger.Out, "Out must implement the io.Writer interface.")
}