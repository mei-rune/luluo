package luluo

func Mult(leftValue, rightValue Value) (Value, error) {
	switch rightValue.Type {
	case ValueNull:
		return Null(), NewArithmeticError("*", leftValue.Type.String(), rightValue.Type.String())
	case ValueBool:
		return Null(), NewArithmeticError("*", leftValue.Type.String(), rightValue.Type.String())
	case ValueString:
		return Null(), NewArithmeticError("*", leftValue.Type.String(), rightValue.Type.String())
	case ValueInt64:
		return multInt(leftValue, rightValue.IntValue())
	case ValueUint64:
		return multUint(leftValue, rightValue.UintValue())
	case ValueFloat64:
		return multFloat(leftValue, rightValue.FloatValue())
	// case ValueDatetime:
	//   return multDatetime(leftValue, IntToDatetime(rightValue.IntValue()))
	// case ValueInterval:
	//   return multInterval(leftValue, IntToInterval(rightValue.IntValue()))
	default:
		return Null(), NewArithmeticError("*", leftValue.Type.String(), rightValue.Type.String())
	}
}

func multInt(left Value, right int64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("*", left.Type.String(), "int")
	case ValueBool:
		return Null(), NewArithmeticError("*", left.Type.String(), "int")
	case ValueString:
		return Null(), NewArithmeticError("*", left.Type.String(), "int")
	case ValueInt64:
		return IntToValue(left.IntValue() * right), nil
	case ValueUint64:
		if right < 0 {
			return IntToValue(int64(left.UintValue()) * right), nil
		}
		return UintToValue(left.UintValue() * uint64(right)), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() * float64(right)), nil
	default:
		return Null(), NewArithmeticError("*", left.Type.String(), "int")
	}
}

func multUint(left Value, right uint64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("*", left.Type.String(), "uint")
	case ValueBool:
		return Null(), NewArithmeticError("*", left.Type.String(), "uint")
	case ValueString:
		return Null(), NewArithmeticError("*", left.Type.String(), "uint")
	case ValueInt64:
		if left.IntValue() < 0 {
			return IntToValue(left.IntValue() * int64(right)), nil
		}
		return UintToValue(uint64(left.IntValue()) * right), nil
	case ValueUint64:
		return UintToValue(left.UintValue() * right), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() * float64(right)), nil
	default:
		return Null(), NewArithmeticError("*", left.Type.String(), "uint")
	}
}

func multFloat(left Value, right float64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("*", left.Type.String(), "float")
	case ValueBool:
		return Null(), NewArithmeticError("*", left.Type.String(), "float")
	case ValueString:
		return Null(), NewArithmeticError("*", left.Type.String(), "float")
	case ValueInt64:
		return FloatToValue(float64(left.IntValue()) * right), nil
	case ValueUint64:
		return FloatToValue(float64(left.UintValue()) * right), nil
	case ValueFloat64:
		return FloatToValue(left.FloatValue() * right), nil
	default:
		return Null(), NewArithmeticError("*", left.Type.String(), "float")
	}
}

// func multInterval(left Value, right time.Duration) (Value, error) {
// 	switch left.Type {
// 	case ValueDatetime:
// 		t := IntToDatetime(left.IntValue())
// 		return DatetimeToValue(t.Add(-right)), nil
// 	case ValueInterval:
// 		t := IntToInterval(left.IntValue())
// 		return IntervalToValue(t -right), nil
// 	default:
// 		return Null(), NewArithmeticError("*", left.Type.String(), "datetime")
// 	}
// }
