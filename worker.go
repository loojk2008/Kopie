package Kopie

import (
	"errors"
	"github.com/cenkalti/backoff"
)

type Worker struct {
	Config    Config
	Procedure Procedure
	protocol  Protocol

	backOff *backoff.ExponentialBackOff
}

// Read it's procedure and determine what to do
func (w *Worker) Read() error {
	switch w.Procedure.Type {
	default:
		return errors.New("Could not detect protocol type.")
	case "pump":
		w.protocol = &Pump{w.Config, w.Procedure}
	case "testprotocol":
		w.protocol = &TestProtocol{w.Config, w.Procedure}
	}
	w.backOff = backoff.NewExponentialBackOff()
	return nil
}

func (w *Worker) Start() error {
	defer w.protocol.End()

	err := backoff.Retry(w.protocol.Initiate, w.backOff)
	if err != nil {
		return err
	}

	err = backoff.Retry(w.protocol.Start, w.backOff)
	if err != nil {
		return err
	}
	return nil
}
