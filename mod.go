package vm

func ModFunc(left, right func(Context) (Value, error)) func(Context) (Value, error) {
	return func(ctx Context) (Value, error) {
		leftValue, err := left(ctx)
		if err != nil {
			return Null(), err
		}
		rightValue, err := right(ctx)
		if err != nil {
			return Null(), err
		}

		switch rightValue.Type {
		case ValueNull:
			return Null(), NewArithmeticError("mod", leftValue.Type.String(), rightValue.Type.String())
		case ValueBool:
			return Null(), NewArithmeticError("mod", leftValue.Type.String(), rightValue.Type.String())
		case ValueString:
			return Null(), NewArithmeticError("mod", leftValue.Type.String(), rightValue.Type.String())
		case ValueInt64:
			return modInt(leftValue, rightValue.IntValue())
		case ValueUint64:
			return modUint(leftValue, rightValue.UintValue())
		// case ValueFloat64:
		//   return modFloat(leftValue, rightValue.FloatValue())
		// case ValueDatetime:
		//   return modDatetime(leftValue, IntToDatetime(rightValue.IntValue()))
		// case ValueInterval:
		//   return modInterval(leftValue, IntToInterval(rightValue.IntValue()))
		default:
			return Null(), NewArithmeticError("mod", leftValue.Type.String(), rightValue.Type.String())
		}
	}
}

func modInt(left Value, right int64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("mod", left.Type.String(), "int")
	case ValueBool:
		return Null(), NewArithmeticError("mod", left.Type.String(), "int")
	case ValueString:
		return Null(), NewArithmeticError("mod", left.Type.String(), "int")
	case ValueInt64:
		return IntToValue(left.IntValue() % right), nil
	case ValueUint64:
		if right < 0 {
			return IntToValue(-int64(left.UintValue() % uint64(-right))), nil
		}
		return UintToValue(left.UintValue() % uint64(right)), nil
		// case ValueFloat64:
		// 	return FloatToValue(left.FloatValue() * float64(right)), nil
	}
	return Null(), NewArithmeticError("mod", left.Type.String(), "int")
}

func modUint(left Value, right uint64) (Value, error) {
	switch left.Type {
	case ValueNull:
		return Null(), NewArithmeticError("mod", left.Type.String(), "uint")
	case ValueBool:
		return Null(), NewArithmeticError("mod", left.Type.String(), "uint")
	case ValueString:
		return Null(), NewArithmeticError("mod", left.Type.String(), "uint")
	case ValueInt64:
		if left.IntValue() < 0 {
			return IntToValue(-int64(uint64(left.IntValue()) % right)), nil
		}
		return UintToValue(uint64(left.IntValue()) % right), nil
	case ValueUint64:
		return UintToValue(left.UintValue() % right), nil
	// case ValueFloat64:
	// 	return FloatToValue(left.FloatValue() * float64(right)), nil
	default:
		return Null(), NewArithmeticError("mod", left.Type.String(), "uint")
	}
}

// func modFloat(left Value, right float64) (Value, error) {
// 	switch left.Type {
// 	case ValueNull:
// 		return Null(), NewArithmeticError("mod", left.Type.String(), "float")
// 	case ValueBool:
// 		return Null(), NewArithmeticError("mod", left.Type.String(), "float")
// 	case ValueString:
// 		return Null(), NewArithmeticError("mod", left.Type.String(), "float")
// 	case ValueInt64:
// 		return FloatToValue(float64(left.IntValue()) * right), nil
// 	case ValueUint64:
// 		return FloatToValue(float64(left.UintValue()) * right), nil
// 	case ValueFloat64:
// 		return FloatToValue(left.FloatValue() * right), nil
// 	default:
// 		return Null(), NewArithmeticError("mod", left.Type.String(), "float")
// 	}
// }

// func modInterval(left Value, right time.Duration) (Value, error) {
// 	switch left.Type {
// 	case ValueDatetime:
// 		t := IntToDatetime(left.IntValue())
// 		return DatetimeToValue(t.Add(-right)), nil
// 	case ValueInterval:
// 		t := IntToInterval(left.IntValue())
// 		return IntervalToValue(t -right), nil
// 	default:
// 		return Null(), NewArithmeticError("mod", left.Type.String(), "datetime")
// 	}
// }
