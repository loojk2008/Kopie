package procedures

import (
	"Kopie/PgToStruct"
	"fmt"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(tables []string, originUrl string, remote *gorm.DB) {
	psql, err := PgToStruct.NewPostgres(originUrl)
	if err != nil {
		panic(err)
	}
	structs, err := psql.TablesToStruct(tables)
	if err != nil {
		panic(err)
	}

	for _, model := range structs {
		fmt.Println("Generate sql useing %s", model)
	}
}
