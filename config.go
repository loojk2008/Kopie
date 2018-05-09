package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
)

type Database struct {
	Name     string  `toml:"name"`
	Type     string  `toml:"type"`
	Host     string  `toml:"host"`
	Port     int     `toml:"port"`
	User     string  `toml:"user"`
	Password string  `toml:"password"`
	Sslmode  string  `toml:"sslmode"`
	Tables   []Table `toml:"tables"`
}

type Table struct {
	Name      string `toml:"name"`
	Replicate bool   `toml:"replicate"`
	alias     string
}

func (d *Database) GetTable(name string) (int, Table) {
	for i, table := range d.Tables {
		if table.Name == name || table.alias == name {
			return i, table
		}
	}
	return 0, Table{}
}

func (d *Database) PostgresUrl() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		d.Host,
		d.Port,
		d.User,
		d.Name,
		d.Password,
		d.Sslmode,
	)
}

func (d *Database) Url() string {
	switch d.Type {
	default:
		return ""
	case "postgres":
		return d.PostgresUrl()
	}
}

func (d *Database) Connect() *gorm.DB {
	db, err := gorm.Open(d.Type, d.Url())
	if err != nil {
		panic(err)
	}
	return db
}

type Procedure struct {
	Name string `toml:"name"`
	Type string `toml:"type"`
	Pump struct {
		Master      string   `toml:"master"`
		Tables      []string `toml:"tables"`
		Slave       string   `toml:"slave"`
		Automigrate bool     `toml:"automigrate"`
		Period 		int      `toml:"period"`

	} `toml:"pump"`
	Test struct {
		File    string `toml:"file"`
		Message string `toml:"message"`
	} `toml:"test"`
}

type Config struct {
	Label      string      `toml:"label"`
	Suffix     string      `toml:"suffix"`
	Databases  []Database  `toml:"databases"`
	Procedures []Procedure `toml:"protocols"`
}

func (c *Config) GetProcedure(name string) Procedure {
	for _, proc := range c.Procedures {
		if proc.Name == name {
			return proc
		}
	}
	return Procedure{}
}

func (c *Config) GetDatabase(name string) Database {
	for _, db := range c.Databases {
		if db.Name == name {
			return db
		}
	}
	return Database{}
}

func (c *Config) Read(path string) {
	tomlConf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if _, err := toml.Decode(string(tomlConf), &c); err != nil {
		panic(err)
	}
}
