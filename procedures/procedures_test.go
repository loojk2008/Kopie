package procedures

import (
	"Kopie"
	"github.com/jinzhu/gorm"
	"testing"
)

func GetConfig() Kopie.Config {
	config := Kopie.Config{}
	config.Path = "../test_cfg.yaml"
	config.Load()
	return config
}

func TestProcedure_AutoMigrate(t *testing.T) {
	c := GetConfig()
	p := c.Procedures
	master := p[0].FormattedSpecs.GetMaster()
	slave := p[0].FormattedSpecs.GetSlave()
	slaveCon, err := gorm.Open(slave.Type, slave.Url)
	if err != nil {
		panic(err)
	}
	AutoMigrate([]string{"test"}, master.Url, slaveCon)
}
