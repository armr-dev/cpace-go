package testConfig

import cpaceLib "filippo.io/cpace"

var IdA = "client"
var IdB = "server"
var Password = "password"
var C = cpaceLib.NewContextInfo(IdA, IdB, nil)
var MsgA, State, _ = cpaceLib.Start(Password, C)
var MsgB, _, _ = cpaceLib.Exchange(Password, C, MsgA)
