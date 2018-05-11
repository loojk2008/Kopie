package main

import "os"

type Kopie struct {
	ConfigPath string
	config Config
}

func NewKopie(path string) Kopie {
	K := Kopie{ConfigPath:path}
	K.config = Config{}
	K.config.Read(K.ConfigPath)
	return K
}

func (k *Kopie) Start() {
	orc := Orchestrator{k.config, nil}
	orc.Start()
}

func main(){
	k := NewKopie(os.Getenv("KOPIE_CONFIG_PATH"))
	k.Start()
}
