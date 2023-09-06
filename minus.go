package luluo

import (
	"time"
)

func Minus(leftValue, rightValue Value) (Value, error) {
	switch rightValue.Type {
	case ValueNull:
		return Null(), NewArithmeticError("-", leftValue.Type.String(), rightValue.Type.String())
	case ValueBool:
		return Null(), NewArithmeticError("-", leftValue.Type.String(), rightValue.Type.String())
	case ValueString:
		return Null(), NewArithmeticError("-", leftValue.Type.String(), rightValue.Type.String())
	case ValueInt64:
		return minusInt(leftValue, rightValue.IntValue())
	case ValueUint64:
		return minusUint(leftValue, rightValue.UintValue())
	case ValueFloat64:
		return minusFloat(leftValue, rightValue.FloatValue())
	case ValueDatetime:
		return minusDatetime(leftValue, IntToDatetime(rightValue.IntValue()))
	case ValueInterval:
		return minusInterval(leftValue, IntToInterval(rightValue.IntValue()))
	default:
		return Null(), NewArithmeticError("-", leftValue.Type.String(), rightValue.Type.String())
	}
}

func minusInt(left Value, right int64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("-", left.Type.String(), "int")
	case ValueBool:
		return Null(), NewArithmeticError("-", left.Type.String(), "int")
	case ValueString:
		return Null(), NewArithmeticError("-", left.Type.String(), "int")
	case ValueInt64:
		return IntToValue(left.IntValue() - right), nil
	case ValueUint64:
		if right < 0 {
			u64 := uint64(-right)
			return UintToValue(left.UintValue() + u64), nil
		}
		return UintToValue(left.UintValue() + uint64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() - float64(right)), nil
	default:
		return Null(), NewArithmeticError("-", left.Type.String(), "int")
	}
}

func minusUint(left Value, right uint64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("-", left.Type.String(), "uint")
	case ValueBool:
		return Null(), NewArithmeticError("-", left.Type.String(), "uint")
	case ValueString:
		return Null(), NewArithmeticError("-", left.Type.String(), "uint")
	case ValueInt64:
		if left.IntValue() < 0 {
			return IntToValue(left.IntValue() - int64(right)), nil
		}
		return IntToValue(left.IntValue() - int64(right)), nil
	case ValueUint64:
		if left.UintValue() > right {
			return UintToValue(left.UintValue() - right), nil
		}
		return IntToValue(-int64(right - left.UintValue())), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() - float64(right)), nil
	default:
		return Null(), NewArithmeticError("-", left.Type.String(), "uint")
	}
}

func minusFloat(left Value, right float64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("-", left.Type.String(), "float")
	case ValueBool:
		return Null(), NewArithmeticError("-", left.Type.String(), "float")
	case ValueString:
		return Null(), NewArithmeticError("-", left.Type.String(), "float")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) - right), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) - right), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() - right), nil
	default:
		return Null(), NewArithmeticError("-", left.Type.String(), "float")
	}
}

func minusDatetime(left Value, right time.Time) (Value, error) {
	if left.Type != ValueDatetime {
		return Null(), NewArithmeticError("-", left.Type.String(), "datetime")
	}

	t := left.DatetimeValue()
	return IntervalToValue(t.Sub(right)), nil
}

func minusInterval(left Value, right time.Duration) (Value, error) {
	switch left.Type {
	case ValueDatetime:
		t := left.DatetimeValue()
		return DatetimeToValue(t.Add(-right)), nil
	case ValueInterval:
		t := left.DurationValue()
		return IntervalToValue(t - right), nil
	default:
		return Null(), NewArithmeticError("-", left.Type.String(), "datetime")
	}
}
