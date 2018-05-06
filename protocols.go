package Kopie

import (
	"Kopie/PgToStruct"
	"fmt"
	"github.com/adam-hanna/arrayOperations"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/hashstructure"
	"strings"
)

type Protocol interface {
	Initiate() error
	Start() error
	End() error
}

type TestProtocol struct {
	Config
	Procedure
}

func (t *TestProtocol) Initiate() error {
	err := createFile(t.Test.File)
	if err != nil {
		return err
	}
	return nil
}

func (t *TestProtocol) Start() error {
	err := writeFile(t.Test.File, t.Test.Message)
	if err != nil {
		return err
	}
	return nil
}

func (t *TestProtocol) End() error {
	err := deleteFile(t.Test.File)
	if err != nil {
		return err
	}
	return nil
}

type Pump struct {
	Config
	Procedure

	master Database
	slave  Database

	masterCon *gorm.DB
	slaveCon  *gorm.DB
}

func CreateTable(params *PgToStruct.TemplateParams) string {
	columns := ""
	for _, column := range params.Fields {
		columns = columns + strings.ToLower(column.Name) + strings.ToUpper(column.Type) + ","
	}
	columns = strings.TrimSuffix(columns, ",")
	return columns
}

func (p *Pump) Initiate() error {
	p.master = p.Config.GetDatabase(p.Procedure.Pump.Master)
	p.slave = p.Config.GetDatabase(p.Procedure.Pump.Slave)

	// Connect to the databases and setup required extensions
	p.masterCon = p.master.Connect()
	p.slaveCon = p.slave.Connect()

	// Get the tables that need to be replicated.
	masterPg, err := PgToStruct.PostgresfromCon(*p.masterCon)
	if err != nil {
		return err
	}
	slavePg, err := PgToStruct.PostgresfromCon(*p.slaveCon)
	if err != nil {
		return err
	}

	// Get mappings of the current tables
	toBeReplicated, err := masterPg.TablesToStruct(p.Procedure.Pump.Tables)
	ExistingTables, err := slavePg.TablesToStruct(p.Procedure.Pump.Tables)

	// Create the tables to be replicated on slave if they do not exist. On name collisions, create a new name
	// procedurally.
	for tableName, table := range toBeReplicated {
		if _, ok := ExistingTables[tableName]; !ok {
			//	if table does not exist, we can simply copy it: easy peasy
			// first we create a proper expression for the columns:
			columns := CreateTable(table)

			// And now we create the table
			p.slaveCon.Exec(`CREATE TABLE IF NOT EXIST %s (%s)`, tableName, columns)

		} else {
			//	If the table does exist it is more complicated. If the tables are equal, it is no problem:
			d, ok := arrayOperations.Difference(table.Fields, ExistingTables[tableName].Fields)
			if !ok {
				//	No difference, the tables are equal
				break
			} else {
				//	Tables are different. First we create a hash of the slice, and append that to the table name.
				hash, err := hashstructure.Hash(d, nil)
				if err != nil {
					return (err)
				}
				name := tableName + string(hash)
				columns := CreateTable(table)

				// And now we create the table. If we already performed this operation last time, it will fail;
				// but we still store the name for use later.
				p.slaveCon.Exec(`CREATE TABLE IF NOT EXIST %s (%s)`, name, columns)
				// Storing the alias for later
				i, _ := p.master.GetTable(tableName)
				p.master.Tables[i].alias = name
			}
		}
	}

	p.masterCon.Exec("CREATE EXTENSION IF NOT EXISTS postgres_fdw")
	p.slaveCon.Exec("CREATE EXTENSION IF NOT EXISTS postgres_fdw")

	//Data is pushed from master to slave
	p.masterCon.Exec(`CREATE SERVER IF NOT EXISTS %s
						 FOREIGN DATA WRAPPER postgres_fdw
						 OPTIONS (host '%s', port '%d', dbname '%s');`,
		p.slave.Name, p.slave.Host, p.slave.Port, p.slave.Name)

	p.masterCon.Exec(`CREATE USER MAPPING IF NOT EXISTS FOR %s
 						  SERVER %s
						  OPTIONS (user '%s', password '%s');`,
		p.slave.User, p.slave.Name, p.slave.User, p.slave.Password)

	p.masterCon.Exec(`IMPORT FOREIGN SCHEMA public
    					  FROM SERVER %s INTO public;`, p.slave.Name, p.slave.Name, p.Name+p.slave.Name)
	return nil
}

func (p *Pump) Start() error { return nil }

func (p *Pump) End() error {
	p.masterCon.Close()
	p.slaveCon.Close()
	return nil
}

func GetModels(tables []string, db gorm.DB) (map[string]*PgToStruct.TemplateParams, error) {
	psql, err := PgToStruct.PostgresfromCon(db)
	if err != nil {
		return nil, err
	}
	structs, err := psql.TablesToStruct(tables)
	if err != nil {
		return nil, err
	}
	return structs, nil
}

func AutoMigrate(structs, remote *gorm.DB) {
	fmt.Println("Automigration has not been implemented yet. :( It is a difficult problem.")
}
