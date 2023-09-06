package vm

import (
	"errors"
)

var ErrNotFound = errors.New("value not found")

var notFound = func(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func SetNotFound(cb func(error) bool) {
	notFound = cb
}

func IsErrNotFound(err error) bool {
	return notFound(err)
}

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

var errorWrapFunc func(err error, msg string) error

func SetErrorWrap(cb func(err error, msg string) error) {
	errorWrapFunc = cb
}

func errorWrap(err error, msg string) error {
	if errorWrapFunc != nil {
		return errorWrapFunc(err, msg)
	}

	return &wrapError{
		err: err,
		msg: msg + ": " + err.Error(),
	}
}

// func WrapAlias(g GetValuer, alias map[string]string) GetValuer {
// 	return GetValueFunc(func(tableAlias, name string) (Value, error) {
// 		tableName, ok := alias[tableAlias]
// 		if ok {
// 			return g.GetValue(tableName, name)
// 		}
// 		return g.GetValue(tableAlias, name)
// 	})
// }

var Funcs = map[string]func(ctx Context, values []Value) (Value, error){
	"round": Round,
}

func CallFunc(call func(Context, []Value) (Value, error), readValues func(Context) ([]Value, error)) func(ctx Context) (Value, error) {
	return func(ctx Context) (Value, error) {
		values, err := readValues(ctx)
		if err != nil {
			return Null(), err
		}
		return call(ctx, values)
	}
}

func Round(ctx Context, values []Value) (Value, error) {
	if len(values) == 0 {
		return Null(), newArgumentError("round", "round argument is missing")
	}
	if len(values) != 2 {
		return Null(), newArgumentError("round", "round argument isnot match")
	}
	decimaldigitValue := values[1]
	decimaldigits := -1
	switch decimaldigitValue.Type {
	// case ValueNull:
	//   return BoolToValue(false), nil
	// case ValueBool:
	// 	if value.BoolValue() {
	// 		return UintToValue(1), nil
	// 	}
	// 	return UintToValue(0), nil
	// case ValueString:
	//   	return ToDatetimeValue(value.Str)
	case ValueInt64:
		if decimaldigitValue.IntValue() < 0 {
			return Null(), newArgumentError("round", "round argument decimaldigits invalid")
		}
		decimaldigits = int(decimaldigitValue.IntValue())
	case ValueUint64:
		decimaldigits = int(decimaldigitValue.UintValue())
	// case ValueFloat64:
	//  	return UintToValue(uint64(value.FloatValue())), nil
	// case ValueDatetime:
	// 	return value, nil
	// case ValueInterval:
	// 	return Null(), NewArithmeticError("convert", value.Type.String(), "datetime")
	default:
		return Null(), newConvertError(nil, decimaldigitValue, "datetime")
	}

	xValue := values[0]
	switch xValue.Type {
	// case ValueInt64:
	// case ValueUint64:
	case ValueFloat64:
		f64 := xValue.FloatValue()
		for i := 0; i < decimaldigits; i++ {
			f64 = f64 * 10
		}
		f64 = float64(int64(f64 + 0.5))

		for i := 0; i < decimaldigits; i++ {
			f64 = f64 / 10
		}
		return FloatToValue(f64), nil
	// case ValueDatetime:
	//  return value, nil
	// case ValueInterval:
	//  return Null(), NewArithmeticError("convert", value.Type.String(), "datetime")
	default:
		return Null(), newConvertError(nil, xValue, "datetime")
	}
}

func newArgumentError(name string, msg string) error {
	return errors.New(msg)
}
