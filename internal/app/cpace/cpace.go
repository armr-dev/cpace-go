package cpace

import (
	cpaceLib "filippo.io/cpace"
	"github.com/armr-dev/cpace-go/internal/utils"
	"time"
)

func NewContextInfo(idA string, idB string, ad []byte) *cpaceLib.ContextInfo {
	defer utils.TimeTrack(time.Now(), "NewContextInfo")
	return cpaceLib.NewContextInfo(idA, idB, ad)
}

func Start(password string, c *cpaceLib.ContextInfo) (msgA []byte, s *cpaceLib.State, err error) {
	defer utils.TimeTrack(time.Now(), "Start")
	return cpaceLib.Start(password, c)
}

func Exchange(password string, c *cpaceLib.ContextInfo, msgA []byte) (msgB []byte, key []byte, err error) {
	defer utils.TimeTrack(time.Now(), "Exchange")
	return cpaceLib.Exchange(password, c, msgA)
}
