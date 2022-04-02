package concurrency_test

import (
	"concurrency"
	"testing"
)

func Test_SignalToQueue(t *testing.T) {
	concurrency.SignalToQueue()
}

func Test_Broadcast(t *testing.T) {
	concurrency.Broadcast()
}
