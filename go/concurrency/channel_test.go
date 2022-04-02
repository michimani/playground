package concurrency_test

import (
	"concurrency"
	"testing"
)

func Test_Channel(t *testing.T) {
	concurrency.Channel()
}

func Test_BufferedChannel(t *testing.T) {
	concurrency.BufferedChannel()
}
