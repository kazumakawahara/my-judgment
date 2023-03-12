package mjerr

import (
	"fmt"
	"runtime"
	"strings"
)

type stackTrace struct {
	fileName string
	funcName string
	line     int
}

func newStackTrace() stackTrace {
	pc, fileName, line, ok := runtime.Caller(2)
	if !ok {
		return stackTrace{}
	}

	funcName := shortFuncName(runtime.FuncForPC(pc))

	return stackTrace{
		fileName: fileName,
		funcName: funcName,
		line:     line,
	}
}

func (st stackTrace) format() string {
	return fmt.Sprintf("%s( %s:%d )", st.funcName, st.fileName, st.line)
}

func shortFuncName(f *runtime.Func) string {
	funcName := f.Name()
	i := strings.LastIndex(funcName, "/")
	funcName = funcName[i+1:]
	i = strings.Index(funcName, ".")

	return funcName[i+1:]
}
