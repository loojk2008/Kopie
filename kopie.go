package Kopie

type Kopie struct {
	config Config
}

func (k *Kopie) Init(path string) {
	if path == "" {
		k.config.Path = "kopie_cfg.yaml"
	} else {
		k.config.Path = path
	}
	k.config.Load()
}
