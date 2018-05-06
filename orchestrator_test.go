package Kopie

import (
	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/assert"
	"testing"
)

var backoffs = make(map[string]*backoff.ExponentialBackOff)

var orchest = Orchestrator{
	c,
	backoffs,
}

func StartWorker() {
	orchest.StartWorker(c.Procedures[1])
}

func TestOrchestrator_StartWorker(t *testing.T) {
	assert.NotPanics(t, StartWorker)
}
