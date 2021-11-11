package runtimeutil

import (
	"testing"
)

func TestCurrentFuncName(t *testing.T) {
	funcName := CurFuncName()
	println(funcName)
}

func TestCurFuncNameWithLine(t *testing.T) {
	funcName, line := CurFuncNameWithLine()
	println(funcName, line)
}

func TestPrepFuncName(t *testing.T) {
	funcName := func() string {
		return PrepFuncName(1)
	}()
	println(funcName)
}

func TestPrepFuncNameWithLine(t *testing.T) {
	funcName, line := func() (string, int) {
		return PrepFuncNameWithLine(1)
	}()
	println(funcName, line)
}
