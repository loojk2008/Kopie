package Kopie

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var c =  getConf()

var testProtocol = TestProtocol{
	c,
	c.Procedures[1],
}

func TestTestProtocol_Initiate(t *testing.T) {
	err := testProtocol.Initiate()
	if err != nil {
		t.Error(err)
	}
}

func TestTestProtocol_Start(t *testing.T) {
	err := testProtocol.Start()
	if err != nil {
		t.Error(err)
	}

	msg, err := readFile(testProtocol.Procedure.Test.File)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, msg, testProtocol.Procedure.Test.Message)
}

func TestTestProtocol_End(t *testing.T) {
	err := testProtocol.End()
	if err != nil {
		t.Error(err)
	}
}