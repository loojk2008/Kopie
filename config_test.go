package Kopie

import (
	"testing"
)

func Test_GetConfig(t *testing.T) {
	c := Config{Path: "test_cfg.yaml"}
	c.Load()

	//Testing if database 0 is correctly read
	database := c.Databases[0]

	if database.Name != "kopie_test" {
		t.Error("Incorrectly loaded database name.")
	}
	if database.Type != "postgres" {
		t.Error("Incorrectly loaded database type.")
	}

	if database.Url != "host=localhost port=5432 user=kopie dbname=kopie_test password=kopietestpw sslmode=disable" {
		t.Error("Incorrectly loaded database url.")
	}

	for _, table := range database.Tables.Tables {

		if table.dblink() != true {
			t.Errorf("%s dblink not correctly interpreted", table.name())
		}
		if table.replicate() != true {
			t.Errorf("%s replicate not correctly interpreted", table.name())
		}
	}

	//Testing if database 1 is correctly read
	database = c.Databases[1]

	if database.Name != "kopie_test2" {
		t.Error("Incorrectly loaded database name.")
	}
	if database.Type != "postgres" {
		t.Error("Incorrectly loaded database type.")
	}

	if database.Url != "host=localhost port=5432 user=kopie dbname=kopie_test2 password=kopietestpw sslmode=disable" {
		t.Error("Incorrectly loaded database url.")
	}

	for _, table := range database.Tables.Tables {

		if table.dblink() != false {
			t.Errorf("%s dblink not correctly interpreted", table.name())
		}
		if table.replicate() != false {
			t.Errorf("%s replicate not correctly interpreted", table.name())
		}
	}
	// Testing if procedurs are properly loaded
	if c.Procedures[0].Name != "mainProcedure" {
		t.Error("Incorrectly loaded first procedure name.")
	}

	if c.Procedures[0].Type != "replication" {
		t.Error("Incorrectly loaded first procedure type.")
	}

	if c.Procedures[0].FormattedSpecs.GetTables().Tables[0].name() != "mytable" {
		t.Error("Incorrectly loaded first procedure first and second table.")
	}

	if c.Procedures[1].Name != "secondProcedure" {
		t.Error("Incorrectly loaded first procedure name.")
	}

	if c.Procedures[1].Type != "replication" {
		t.Error("Incorrectly loaded first procedure type.")
	}

	if c.Procedures[1].FormattedSpecs.GetTables().Tables[0].name() != "mytable" {
		t.Error("Incorrectly loaded first procedure first and second table.")
	}
}
