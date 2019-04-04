package base

import "github.com/hyperledger/fabric/core/chaincode/shim"

// Logger is a logger interface mainly used for mocking Chaincode logging during testing
type Logger interface {
	SetLevel(level shim.LoggingLevel)
	IsEnabledFor(level shim.LoggingLevel) bool
	Debug(args ...interface{})
	Info(args ...interface{})
	Notice(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Critical(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
}
