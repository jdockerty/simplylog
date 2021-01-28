package simplylog_test

import (
	"io"
	"testing"

	"github.com/jdockerty/simplylog"
	"github.com/stretchr/testify/assert"
)


func TestNewAreDefaults(t *testing.T) {
	assert := assert.New(t)
	logger := simplylog.New()

	assert.Implements((*io.Writer)(nil), logger.Out, "Out must implement the io.Writer interface.")
	
}

func TestSetOutputChanged(t *testing.T) {

}