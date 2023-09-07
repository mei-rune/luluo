package luluo

func (v Value) Uminus() (Value, error) {
	switch v.Type {
	// case ValueNull:
	//   return Null(), NewArithmeticError("-", value.Type.String(), "")
	// case ValueBool:
	//   return Null(), NewArithmeticError("-", value.Type.String(), "")
	// case ValueString:
	//   return Null(), NewArithmeticError("-", value.Type.String(), "")
	case ValueInt64:
		return IntToValue(-v.IntValue()), nil
	case ValueUint64:
		return IntToValue(-int64(v.UintValue())), nil
	case ValueFloat64:
		return FloatToValue(-v.FloatValue()), nil
	case ValueInterval:
		return IntervalToValue(-v.DurationValue()), nil
	default:
		return Null(), NewArithmeticError("-", v.Type.String(), "")
	}
}

func  Uminus(v Value) (Value, error) {
	return v.Uminus()
}