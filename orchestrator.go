package main

import (
	"github.com/cenkalti/backoff"
)

type Orchestrator struct {
	Conf     Config
	backOffs map[string]*backoff.ExponentialBackOff
}

func (o *Orchestrator) StartWorker(proc Procedure) {
	o.backOffs[proc.Name] = backoff.NewExponentialBackOff()
	w := Worker{
		o.Conf,
		proc,
		nil,
		nil,
	}

	err := w.Read()
	if err != nil {
		panic(err)
	}

	err = backoff.Retry(w.Start, o.backOffs[proc.Name])
	if err != nil {
		panic(err)
	}
}

func (o *Orchestrator) Start() {
	for _, proc := range o.Conf.Procedures {
		go o.StartWorker(proc)
	}
	select {}
}
