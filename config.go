package Kopie

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"sort"
)

type Config struct {
	Path       string
	Databases  []Database
	Procedures []Procedure
}

func (c Config) GetDatabase(name string) (Database, error) {
	for _, db := range c.Databases {
		if db.Name == name {
			return db, nil
		}
	}
	return Database{}, errors.New("Database not present in config.")
}

func (c Config) GetProcedure(name string) (Procedure, error) {
	for _, p := range c.Procedures {
		if p.Name == name {
			return p, nil
		}
	}
	return Procedure{}, errors.New("Procedure not present in config.")
}

type Database struct {
	Name   string
	Type   string
	Url    string
	Tables Tables
}

func (d Database) GetTable(name string) Table {
	for _, table := range d.Tables.Tables {
		if table.name() == name {
			return table
		}
	}
	return nil
}

type Procedure struct {
	Name           string
	Type           string
	Automigrate    bool
	Specifications Specification `yaml:"specification"`
	FormattedSpecs SpecificationInterface
}

type Tables struct {
	Tables []Table `yaml:"table"`
}

func (t Tables) GetTables(names ...string) Tables {
	ret := Tables{}
	sort.Strings(names)
	for _, table := range t.Tables {
		res := sort.SearchStrings(names, table.name())
		if res != len(names) {
			ret.Tables = append(ret.Tables, table)
		}
	}
	return ret
}

type Table map[string]string

func (i Table) name() string {
	return i["name"]
}

func (i Table) replicate() bool {
	switch i["replicate"] {
	case "true":
		return true
	case "false":
		return false
	default:
		panic("Could not interpret replicate value.")
	}
}

func (i Table) dblink() bool {
	switch i["db_link"] {
	case "true":
		return true
	case "false":
		return false
	default:
		panic("Could not interpret db_link value.")
	}
}

type Specification map[string]string

func (s Specification) Master() string {
	return s["master"]
}

func (s Specification) Slave() string {
	return s["slave"]
}

func (s Specification) Tables() string {
	return s["tables"]
}

type Replication struct {
	Master Database
	Slave  Database
	Tables Tables
}

func (r *Replication) GetTables() Tables {
	return r.Tables
}

func (r *Replication) GetMaster() Database {
	return r.Master
}

func (r *Replication) GetSlave() Database {
	return r.Slave
}

type SpecificationInterface interface {
	GetTables() Tables
	GetMaster() Database
	GetSlave() Database
}

func (r *Replication) Init(config *Config, spec Specification) {
	var err error

	// find the Master db
	r.Master, err = config.GetDatabase(spec.Master())
	if err != nil {
		panic(err)
	}

	// find slave db
	r.Slave, err = config.GetDatabase(spec.Slave())
	if err != nil {
		panic(err)
	}

	// get the tables from master that need to be replicated to slave
	r.Tables = r.Master.Tables.GetTables(spec.Tables())
}

func (c *Config) Load() {
	yamlFile, err := ioutil.ReadFile(c.Path)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Lets convert the procedures to the appropriate interface
	for key, proc := range c.Procedures {
		replication := Replication{}
		switch proc.Type {
		case "replication":
			replication.Init(c, proc.Specifications)
			c.Procedures[key].FormattedSpecs = &replication
		}
	}
}
