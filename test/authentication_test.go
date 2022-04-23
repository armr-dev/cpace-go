package test

import (
	cpaceLib "filippo.io/cpace"
	"github.com/armr-dev/cpace-go/internal/testConfig"
	"testing"
)

func BenchmarkAuthentication(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Client
		cClient := cpaceLib.NewContextInfo(testConfig.IdA, testConfig.IdB, nil)
		msgA, state, _ := cpaceLib.Start(testConfig.Password, cClient)

		// Server
		cServer := cpaceLib.NewContextInfo(testConfig.IdA, testConfig.IdB, nil)
		msgB, _, _ := cpaceLib.Exchange(testConfig.Password, cServer, msgA)

		// Client
		_, _ = state.Finish(msgB)
	}
}
