package Kopie

import "testing"

func TestKopie_Init(t *testing.T) {
	kopie := Kopie{}
	kopie.Init("test_cfg.yaml")
	if kopie.config.Databases[0].Name != "kopie_test" {
		t.Error("Incorrectly loaded kopie config")
	}
}
