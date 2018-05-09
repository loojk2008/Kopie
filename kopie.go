package main

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
