package test

import (
	cpaceLib "filippo.io/cpace"
	"github.com/armr-dev/cpace-go/internal/testConfig"
	"testing"
)

func BenchmarkNewContextInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpaceLib.NewContextInfo(testConfig.IdA, testConfig.IdB, nil)
	}
}

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpaceLib.Start(testConfig.Password, testConfig.C)
	}
}

func BenchmarkExchange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cpaceLib.Exchange(testConfig.Password, testConfig.C, testConfig.MsgA)
	}
}

func BenchmarkFinish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testConfig.State.Finish(testConfig.MsgB)
	}
}
