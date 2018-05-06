package Kopie

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var worker = Worker{
	c,
	c.Procedures[1],
	nil,
	nil,
}


func TestWorker_Read(t *testing.T) {
	assert.NoError(t, worker.Read())
	assert.IsType(t, &TestProtocol{}, worker.protocol)
	assert.NotEmpty(t, worker.protocol)
	assert.NotEmpty(t, worker.backOff)
}

func TestWorker_Start(t *testing.T) {
	assert.NoError(t, worker.Read())
	assert.NoError(t, worker.Start())
}