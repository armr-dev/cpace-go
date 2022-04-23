package test

import (
	"github.com/armr-dev/cpace-go/internal/app/client"
	"github.com/armr-dev/cpace-go/internal/testConfig"
	"testing"
)

func BenchmarkAuthenticationClientServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client.Registration(testConfig.IdA, testConfig.Password)
		client.Authentication(testConfig.IdA, testConfig.Password)
	}
}
