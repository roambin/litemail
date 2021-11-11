package runtimeutil

import (
	"runtime"
)

/*
get current function name with path

e.g.
file: example.go
1 func A() {
2 	print(CurFuncName())
3 }
output: "example.A()"
*/
func CurFuncName() (funcName string) {
	funcName, _ = PrepFuncNameWithLine(1)
	return
}

/*
get current function name with path and current line

e.g.
file: example.go
1 func A() {
2 	print(CurFuncNameWithLine())
3 }
output: "example.A() 2"
*/
func CurFuncNameWithLine() (funcName string, line int) {
	return PrepFuncNameWithLine(1)
}

/*
get prep function name with path
when skip = 0, equal to CurFuncName()

e.g.
file: example.go
1 func A() {
2 	func() {
3 		print(PrepFuncName(1))
4 	}()
5 }
output: "example.A()"
*/
func PrepFuncName(skip int) (funcName string) {
	funcName, _ = PrepFuncNameWithLine(skip + 1)
	return
}

/*
get current function name with path and current line
when skip = 0, equal to CurFuncNameWithLine()

e.g.
file: example.go
1 func A() {
2 	func() {
3 		print(PrepFuncNameWithLine(1))
4 	}()
5 }
output: "example.A() 3"
*/
func PrepFuncNameWithLine(skip int) (funcName string, line int) {
	pc, _, line, _ := runtime.Caller(skip + 1)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
