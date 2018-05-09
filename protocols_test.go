package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var c = getConf()

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

func TestPump(t *testing.T) {
	pump := Pump{
		c,
		c.Procedures[0],
		Database{},
		Database{},
		&gorm.DB{},
		&gorm.DB{},
	}
	err := pump.Initiate()
	if err != nil {
		t.Error(err)
	}
	go func() {
		err = pump.Start()
		if err != nil {
			t.Error(err)
		}
	}()

	type Person struct {
		Name string
		Age  int
	}

	p := Person{}
	// Wait for data transfer
	time.Sleep(1 * time.Second)
	pump.slaveCon.Raw("select * from test_kopie where name = 'karel';").Scan(&p)
	assert.Equal(t, 18, p.Age)

	err = pump.End()
	if err != nil {
		t.Error(err)
	}
}
