package logutil

import (
	"context"
	"errors"
	"log"
)

func LogIfError(ctx context.Context, err error) {
	LogPrepIfError(ctx, err, 1)
}

func LogPrepIfError(ctx context.Context, err error, skip int) {
	handler := &ErrorHandler{
		wrapError: false,
		logError:  true,
	}
	_ = handler.HandlePrepIfError(ctx, err, skip+1)
}

func WrapIfError(ctx context.Context, err error) error {
	return WrapPrepIfError(ctx, err, 1)
}

func WrapPrepIfError(ctx context.Context, err error, skip int) error {
	handler := &ErrorHandler{
		wrapError: true,
		logError:  false,
	}
	return handler.HandlePrepIfError(ctx, err, skip+1)
}

func LogWrapIfError(ctx context.Context, err error) error {
	return LogWrapPrepIfError(ctx, err, 1)
}

func LogWrapPrepIfError(ctx context.Context, err error, skip int) error {
	handler := &ErrorHandler{
		wrapError: true,
		logError:  true,
	}
	return handler.HandlePrepIfError(ctx, err, skip+1)
}

type ErrorHandler struct {
	wrapError bool
	logError  bool
}

func (p *ErrorHandler) HandlePrepIfError(ctx context.Context, err error, skip int) error {
	if err == nil {
		return nil
	}
	switch errInstance := err.(type) {
	default:
		if p.logError {
			Log(ctx, func(ctx context.Context, format string, v ...interface{}) {
				log.Printf("Error " + format, v...)
			}, skip+1, err.Error())
		}
		var newErr error
		if p.wrapError {
			message := GetPrepMessage(ctx, skip+1, errInstance.Error())
			newErr = errors.New(message)
		}
		return newErr
	}
}
