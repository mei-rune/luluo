package luluo

func Div(leftValue, rightValue Value) (Value, error) {
	switch rightValue.Type {
	case ValueNull:
		return Null(), NewArithmeticError("/", leftValue.Type.String(), rightValue.Type.String())
	case ValueBool:
		return Null(), NewArithmeticError("/", leftValue.Type.String(), rightValue.Type.String())
	case ValueString:
		return Null(), NewArithmeticError("/", leftValue.Type.String(), rightValue.Type.String())
	case ValueInt64:
		return DivInt(leftValue, rightValue.IntValue())
	case ValueUint64:
		return DivUint(leftValue, rightValue.UintValue())
	case ValueFloat64:
		return DivFloat(leftValue, rightValue.FloatValue())
	// case ValueDatetime:
	//   return divDatetime(leftValue, IntToDatetime(rightValue.IntValue()))
	// case ValueInterval:
	//   return divInterval(leftValue, IntToInterval(rightValue.IntValue()))
	default:
		return Null(), NewArithmeticError("/", leftValue.Type.String(), rightValue.Type.String())
	}
}

func DivInt(left Value, right int64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("/", left.Type.String(), "int")
	case ValueBool:
		return Null(), NewArithmeticError("/", left.Type.String(), "int")
	case ValueString:
		return Null(), NewArithmeticError("/", left.Type.String(), "int")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) / float64(right)), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) / float64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() / float64(right)), nil
	default:
		return Null(), NewArithmeticError("/", left.Type.String(), "int")
	}
}

func DivUint(left Value, right uint64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("/", left.Type.String(), "uint")
	case ValueBool:
		return Null(), NewArithmeticError("/", left.Type.String(), "uint")
	case ValueString:
		return Null(), NewArithmeticError("/", left.Type.String(), "uint")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) / float64(right)), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) / float64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() / float64(right)), nil
	default:
		return Null(), NewArithmeticError("/", left.Type.String(), "uint")
	}
}

func DivFloat(left Value, right float64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("/", left.Type.String(), "float")
	case ValueBool:
		return Null(), NewArithmeticError("/", left.Type.String(), "float")
	case ValueString:
		return Null(), NewArithmeticError("/", left.Type.String(), "float")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) / float64(right)), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) / float64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() / right), nil
	default:
		return Null(), NewArithmeticError("/", left.Type.String(), "float")
	}
}

// func divDatetime(left Value, right time.Time) (Value, error) {
// 	if left.Type != ValueDatetime {
// 		return Null(), NewArithmeticError("/", left.Type.String(), "datetime")
// 	}

// 	t := IntToDatetime(left.IntValue())
// 	return IntervalToValue(t.Sub(right)), nil
// }

// func divInterval(left Value, right time.Duration) (Value, error) {
// 	switch left.Type {
// 	case ValueDatetime:
// 		t := IntToDatetime(left.IntValue())
// 		return DatetimeToValue(t.Add(-right)), nil
// 	case ValueInterval:
// 		t := IntToInterval(left.IntValue())
// 		return IntervalToValue(t -right), nil
// 	default:
// 		return Null(), NewArithmeticError("/", left.Type.String(), "datetime")
// 	}
// }
