package logutil

import (
	"context"
	"fmt"
	"github.com/roambin/litemail/utils/runtimeutil"
)

func GetMessage(ctx context.Context, format string, v ...interface{}) string {
	return GetPrepMessage(ctx, 1, format, v...)
}

func GetPrepMessage(ctx context.Context, skip int, format string, v ...interface{}) string {
	message := fmt.Sprintf(format, v...)
	funcName, line := runtimeutil.PrepFuncNameWithLine(skip + 1)
	message = fmt.Sprintf("%s:%d\n%s", funcName, line, message)
	return message
}
