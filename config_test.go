package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var configpath = "kopie.toml"

func getConf() Config {
	c := Config{}
	c.Read(configpath)
	return c
}

func TestConfig_Read(t *testing.T) {
	c := getConf()
	assert.Equal(t, "Kopie default configuration.", c.Label)

	// Testing if we loaded the first database correctly
	assert.Equal(t, "kopie_test", c.Databases[0].Name)
	assert.Equal(t, "postgres", c.Databases[0].Type)
	assert.Equal(t, "localhost", c.Databases[0].Host)
	assert.Equal(t, 5432, c.Databases[0].Port)
	assert.Equal(t, "kopietestpw", c.Databases[0].Password)
	assert.Equal(t, "disable", c.Databases[0].Sslmode)

	assert.Equal(t, "test", c.Databases[0].Tables[0].Name)
	assert.Equal(t, true, c.Databases[0].Tables[0].Replicate)

	assert.Equal(t, "test2", c.Databases[0].Tables[1].Name)
	assert.Equal(t, true, c.Databases[0].Tables[1].Replicate)

	// Testing if we loaded the second database correctly
	assert.Equal(t, "kopie_test2", c.Databases[1].Name)
	assert.Equal(t, "postgres", c.Databases[1].Type)
	assert.Equal(t, "localhost", c.Databases[1].Host)
	assert.Equal(t, 5432, c.Databases[1].Port)
	assert.Equal(t, "kopietestpw", c.Databases[1].Password)
	assert.Equal(t, "disable", c.Databases[1].Sslmode)

	assert.Equal(t, "test", c.Databases[1].Tables[0].Name)
	assert.Equal(t, true, c.Databases[1].Tables[0].Replicate)

	assert.Equal(t, "test2", c.Databases[1].Tables[1].Name)
	assert.Equal(t, true, c.Databases[1].Tables[1].Replicate)

	// Testing if we loaded the first procedure correctly
	assert.Equal(t, "mainProcedure", c.Procedures[0].Name)
	assert.Equal(t, "pump", c.Procedures[0].Type)
	assert.Equal(t, "kopie_test", c.Procedures[0].Pump.Master)
	assert.Equal(t, "kopie_test2", c.Procedures[0].Pump.Slave)
	assert.Equal(t, true, c.Procedures[0].Pump.Automigrate)
	assert.Equal(t, "test", c.Procedures[0].Pump.Tables[0])
	assert.Equal(t, "test2", c.Procedures[0].Pump.Tables[1])

	// Testing if we loaded the second procedure correctly
	assert.Equal(t, "testProcedure", c.Procedures[1].Name)
	assert.Equal(t, "testprotocol", c.Procedures[1].Type)
	assert.Equal(t, "testprotocol.test", c.Procedures[1].Test.File)
	assert.Equal(t, "Testing", c.Procedures[1].Test.Message)

}

func TestDatabase_PostgresUrl(t *testing.T) {
	c := getConf()

	assert.Equal(t, "host=localhost port=5432 user=kopie dbname=kopie_test password=kopietestpw sslmode=disable", c.Databases[0].PostgresUrl())
	assert.Equal(t, "host=localhost port=5432 user=kopie dbname=kopie_test2 password=kopietestpw sslmode=disable", c.Databases[1].PostgresUrl())
}

func TestDatabase_Url(t *testing.T) {
	c := getConf()

	assert.Equal(t, "host=localhost port=5432 user=kopie dbname=kopie_test password=kopietestpw sslmode=disable", c.Databases[0].Url())
	assert.Equal(t, "host=localhost port=5432 user=kopie dbname=kopie_test2 password=kopietestpw sslmode=disable", c.Databases[1].Url())
}

func TestDatabase_Connect(t *testing.T) {
	c := getConf()
	// If this throws an error the connection failed
	c.Databases[0].Connect()
}
