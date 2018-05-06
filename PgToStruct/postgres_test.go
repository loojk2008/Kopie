package PgToStruct

import (
	"github.com/jinzhu/gorm"
	"testing"
)

var url = "host=localhost port=5432 user=kopie dbname=kopie_test password=kopietestpw sslmode=disable"

func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func GetConnection() *Postgres {
	psql, _ := NewPostgres(url)
	return psql
}

func tErr(t *testing.T, err error) {
	if err != nil {
		t.Error("Error in creating postgres instance")
	}
}

func TestNewPostgres(t *testing.T) {
	_, err := NewPostgres(url)
	tErr(t, err)
}

func TestPostgres_RetrieveTables(t *testing.T) {
	psql := GetConnection()
	tables, err := psql.retrieveAllTables()
	tErr(t, err)
	columns, err := tables.Columns()
	tErr(t, err)
	eq := testEq(columns, []string{"table_name"})
	if !eq {
		t.Error("Failed to retrieve tables")
	}
}

func TestPostgres_RetrieveFields(t *testing.T) {
	psql := GetConnection()
	fields, err := psql.RetrieveFields("test")
	tErr(t, err)
	if fields[0].Name != "name" {
		t.Error("Failed to retrieve first field name")
	}

	if fields[1].Name != "date" {
		t.Error("Failed to retrieve second field name")
	}

	if fields[0].Type != "character varying" {
		t.Error("Failed to retrieve first field type")
	}

	if fields[1].Type != "timestamp without time zone" {
		t.Error("Failed to retrieve second field type")
	}
}

func TestPostgres_RetrieveSelectedTables(t *testing.T) {
	psql := GetConnection()
	tables, err := psql.retrieveSelectedTables([]string{"test"})
	tErr(t, err)
	clmn, _ := tables.Columns()
	if clmn[0] != "table_name" {
		t.Error("Failed to retrieve selected table")
	}
}

func TestPostgres_TablesToStruct(t *testing.T) {
	psql := GetConnection()
	tables, err := psql.TablesToStruct([]string{"test"})
	tErr(t, err)
	if tables["test"].Name != "Test" {
		t.Error("Failed to retrieve table name")
	}
	if tables["test"].Fields[0].Name != "Name" {
		t.Error("Failed to retrieve first field name")
	}

	if tables["test"].Fields[1].Name != "Date" {
		t.Error("Failed to retrieve second field name")
	}
}

func TestPostgresfromCon(t *testing.T) {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		t.Error(err)
	}
	// if this errors the test fails.
	PostgresfromCon(*db)
}
