package logger

import (
	"fmt"
	"github.com/juricaKenda/golang-rest-example/src/server"
)

type impl struct {
}

func New() server.ILogger {
	i := &impl{}
	i.Trace("creating logger....")
	return i
}

func (i *impl) Error(msg string) {
	fmt.Printf("[ERR] %s\n", msg)
}

func (i *impl) Trace(msg string) {
	fmt.Printf("[TRACE] %s\n", msg)
}
