package luluo

func (v Value) Round(decimaldigits int) (Value, error) {
	switch v.Type {
	case ValueInt64:
		return v, nil
	case ValueUint64:
		return v, nil
	case ValueFloat64:
		f64 := v.FloatValue()
		for i := 0; i < decimaldigits; i++ {
			f64 = f64 * 10
		}
		f64 = float64(int64(f64 + 0.5))

		for i := 0; i < decimaldigits; i++ {
			f64 = f64 / 10
		}
		return FloatToValue(f64), nil
	default:
		return Null(), NewArithmeticError("round", v.Type.String(), "float")
	}
}
