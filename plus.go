package vm

import (
	"time"

	"errors"
)

func NewArithmeticError(op, left, right string) error {
	return errors.New("cloudn't '" + left + "' " + op + " '" + right + "'")
}

func PlusFunc(left, right func(Context) (Value, error)) func(Context) (Value, error) {
	return func(ctx Context) (Value, error) {
		leftValue, err := left(ctx)
		if err != nil {
			return Null(), err
		}
		rightValue, err := right(ctx)
		if err != nil {
			return Null(), err
		}

		return Plus(leftValue, rightValue)
	}
}

func Plus(leftValue, rightValue Value) (Value, error) {
	switch rightValue.Type {
	case ValueNull:
		return Null(), NewArithmeticError("+", leftValue.Type.String(), rightValue.Type.String())
	case ValueBool:
		return Null(), NewArithmeticError("+", leftValue.Type.String(), rightValue.Type.String())
	case ValueString:
		return Null(), NewArithmeticError("+", leftValue.Type.String(), rightValue.Type.String())
	case ValueInt64:
		return plusInt(leftValue, rightValue.IntValue())
	case ValueUint64:
		return plusUint(leftValue, rightValue.UintValue())
	case ValueFloat64:
		return plusFloat(leftValue, rightValue.FloatValue())
	case ValueDatetime:
		return plusDatetime(leftValue, IntToDatetime(rightValue.IntValue()))
	case ValueInterval:
		return plusInterval(leftValue, IntToInterval(rightValue.IntValue()))
	default:
		return Null(), NewArithmeticError("+", leftValue.Type.String(), rightValue.Type.String())
	}
}

func plusInt(left Value, right int64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("+", left.Type.String(), "int")
	case ValueBool:
		return Null(), NewArithmeticError("+", left.Type.String(), "int")
	case ValueString:
		return Null(), NewArithmeticError("+", left.Type.String(), "int")
	case ValueInt64:
		return IntToValue(left.IntValue() + right), nil
	case ValueUint64:
		if right < 0 {
			u64 := uint64(-right)
			if left.UintValue() < u64 {
				return IntToValue(right + int64(left.UintValue())), nil
			}
			return UintToValue(left.UintValue() - u64), nil
		}
		return UintToValue(left.UintValue() + uint64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() + float64(right)), nil
	default:
		return Null(), NewArithmeticError("+", left.Type.String(), "int")
	}
}

func plusUint(left Value, right uint64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("+", left.Type.String(), "uint")
	case ValueBool:
		return Null(), NewArithmeticError("+", left.Type.String(), "uint")
	case ValueString:
		return Null(), NewArithmeticError("+", left.Type.String(), "uint")
	case ValueInt64:
		if left.IntValue() < 0 {
			u64 := uint64(-left.IntValue())
			if u64 > right {
				return IntToValue(left.IntValue() + int64(right)), nil
			}
			return UintToValue(right - u64), nil
		}
		return IntToValue(left.IntValue() + int64(right)), nil
	case ValueUint64:
		return UintToValue(left.UintValue() + right), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() + float64(right)), nil
	default:
		return Null(), NewArithmeticError("+", left.Type.String(), "uint")
	}
}

func plusFloat(left Value, right float64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("+", left.Type.String(), "float")
	case ValueBool:
		return Null(), NewArithmeticError("+", left.Type.String(), "float")
	case ValueString:
		return Null(), NewArithmeticError("+", left.Type.String(), "float")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) + right), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) + right), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() + float64(right)), nil
	default:
		return Null(), NewArithmeticError("+", left.Type.String(), "float")
	}
}

func plusDatetime(left Value, right time.Time) (Value, error) {
	if left.Type != ValueInterval {
		return Null(), NewArithmeticError("+", left.Type.String(), "datetime")
	}
	return DatetimeToValue(right.Add(left.DurationValue())), nil
}

func plusInterval(left Value, right time.Duration) (Value, error) {
	if left.Type != ValueDatetime {
		return Null(), NewArithmeticError("+", left.Type.String(), "datetime")
	}

	t := left.DatetimeValue()
	return DatetimeToValue(t.Add(right)), nil
}
