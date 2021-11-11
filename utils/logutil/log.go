package logutil

import (
	"context"
	"log"
)

/*
these functions are same to log.XXX but will fill current function name and line

e.g.
file: example.go
1 func A() {
2 	Info(context.Background(), "%s", "message")
3 }
output: "XXX example/A:2 message"
*/

func Info(ctx context.Context, format string, v ...interface{}) {
	Log(ctx, func(ctx context.Context, format string, v ...interface{}) {
		log.Printf(format, v...)
	}, 1, format, v...)
}

func Log(ctx context.Context,
	logFunc func(ctx context.Context, format string, v ...interface{}),
	skip int,
	format string,
	v ...interface{}) string {
	message := GetPrepMessage(ctx, skip+1, format, v...)
	logFunc(ctx, message)
	return message
}
