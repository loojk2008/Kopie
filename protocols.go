package Kopie

import (
	"Kopie/PgToStruct"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Protocol interface {
	Initiate() error
	Start()	   error
	End()	error
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

func (t *TestProtocol) Start()	error {
	err := writeFile(t.Test.File, t.Test.Message)
	if err != nil {
		return err
	}
	return nil
}

func (t *TestProtocol) End() error {
	err := deleteFile(t.Test.File)
	if err != nil{
		return err
	}
	return nil
}

type Pump struct {
	Config
	Procedure
}

func (p *Pump) Initiate() error {return nil}
func (p *Pump) Start()	error {return nil}
func (p *Pump) End()	error {return nil}



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
