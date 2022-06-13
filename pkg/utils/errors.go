package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

func LogError(err error) {
	wrappedError := tracerr.Wrap(err)
	traceback := tracerr.StackTrace(wrappedError)
	for _, traceItem := range traceback[:3] {
		logrus.Error(struct {
			FuncName   string
			LineNumber int
			FilePath   string
			Error      string
		}{
			FuncName:   traceItem.Func,
			LineNumber: traceItem.Line,
			FilePath:   traceItem.Path,
			Error:      err.Error(),
		})
	}
}
