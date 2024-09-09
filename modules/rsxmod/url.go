package rsxmod

import (
	"context"
	"fmt"
	"net/url"

	"github.com/risor-io/risor/object"
)

func Require(funcName string, count int, args []object.Object) *object.Error {
	nArgs := len(args)
	if nArgs != count {
		if count == 1 {
			return object.Errorf(
				fmt.Sprintf("type error: %s() takes exactly 1 argument (%d given)",
					funcName, nArgs))
		}
		return object.Errorf(
			fmt.Sprintf("type error: %s() takes exactly %d arguments (%d given)",
				funcName, count, nArgs))
	}
	return nil
}

func QueryEscape(ctx context.Context, args ...object.Object) object.Object {
	if err := Require("query_escape", 1, args); err != nil {
		return err
	}
	q, err := object.AsString(args[0])
	if err != nil {
		return err
	}

	return object.NewString(url.QueryEscape(q))
}

func PathEscape(ctx context.Context, args ...object.Object) object.Object {
	if err := Require("path_escape", 1, args); err != nil {
		return err
	}
	q, err := object.AsString(args[0])
	if err != nil {
		return err
	}

	return object.NewString(url.PathEscape(q))
}

func Module() *object.Module {
	return object.NewBuiltinsModule("rsxmod",
		map[string]object.Object{
			"query_escape": object.NewBuiltin("query_escape", QueryEscape),
		},
		QueryEscape,
	)
}
