package vm

import (
	"strconv"
	"strings"
	"time"
)

func (r *Value) CompareTo(to Value, opt CompareOption) (int, error) {
	switch to.Type {
	case ValueNull:
		if r.IsNil() {
			return 0, nil
		}
		return 1, nil
	case ValueBool:
		return r.CompareToBool(to.BoolValue(), opt)
	case ValueString:
		return r.CompareToString(to.StrValue(), opt)
	case ValueInt64:
		return r.CompareToInt64(to.IntValue(), opt)
	case ValueUint64:
		return r.CompareToUint64(to.UintValue(), opt)
	case ValueFloat64:
		return r.CompareToFloat64(to.FloatValue(), opt)
	case ValueDatetime:
		return r.CompareToDatetime(to.TimeUnixValue(), opt)
	case ValueInterval:
		return r.CompareToInterval(to.DurationValue(), opt)
	default:
		return 0, ErrUnknownValueType
	}
}

func (r *Value) CompareToBool(to bool, opt CompareOption) (int, error) {
	if r.Type == ValueBool {
		b := r.BoolValue()
		if b == to {
			return 0, nil
		}
		if b {
			return 1, nil
		}
		return -1, nil
	}

	if opt.Weak {
		switch r.Type {
		case ValueString:
			switch r.StrValue() {
			case "1", "t", "T", "true", "TRUE", "True":
				if to {
					return 0, nil
				}
				return -1, nil
			case "0", "f", "F", "false", "FALSE", "False":
				if to {
					return -1, nil
				}
				return 0, nil
			}
		case ValueInt64:
			if r.IntValue() == 0 {
				if to {
					return -1, nil
				}
				return 0, nil
			}
			if to {
				return 0, nil
			}
			return -1, nil
		case ValueUint64:
			if r.UintValue() == 0 {
				if to {
					return -1, nil
				}
				return 0, nil
			}
			if to {
				return 0, nil
			}
			return -1, nil
		}
	}
	return 0, NewTypeError(r, r.Type.String(), "bool")
}

func (r *Value) CompareToString(to string, opt CompareOption) (int, error) {
	if r.Type == ValueString {
		if opt.IgnoreCase {
			if strings.EqualFold(r.StrValue(), to) {
				return 0, nil
			}
			aS := strings.ToUpper(r.StrValue())
			toS := strings.ToUpper(to)

			if aS > toS {
				return 1, nil
			}
			return -1, nil
		}
		if r.StrValue() == to {
			return 0, nil
		}
		if r.StrValue() > to {
			return 1, nil
		}
		return -1, nil
	}

	if opt.Weak {
		switch r.Type {
		case ValueInt64:
			toI, err := strconv.ParseInt(to, 10, 64)
			if err != nil {
				return 0, NewTypeError(r, "string", "int")
			}
			if r.IntValue() > toI {
				return 1, nil
			}
			if r.IntValue() < toI {
				return -1, nil
			}
			return 0, nil
		case ValueUint64:
			toU, err := strconv.ParseUint(to, 10, 64)
			if err != nil {
				toI, err := strconv.ParseInt(to, 10, 64)
				if err != nil {
					return 0, NewTypeError(r, "string", "uint")
				}
				if toI < 0 {
					return 1, nil
				}
				toU = uint64(toI)
			}
			if r.UintValue() > toU {
				return 1, nil
			}
			if r.UintValue() < toU {
				return -1, nil
			}
			return 0, nil
		case ValueFloat64:
			toF, err := strconv.ParseFloat(to, 64)
			if err != nil {
				return 0, NewTypeError(r, "string", "uint")
			}
			if r.FloatValue() > toF {
				return 1, nil
			}
			return -1, nil
		}
	}
	return 0, NewTypeError(r, r.Type.String(), "string")
}

func compareToFloat64(a string, b float64, opt CompareOption, deferr error) (int, error) {
	aF, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, deferr
	}
	if aF > b {
		return 1, nil
	}
	return -1, nil
}

func (r *Value) CompareToInt64(to int64, opt CompareOption) (int, error) {
	switch r.Type {
	case ValueString:
		if opt.Weak {
			s := r.StrValue()
			if strings.HasSuffix(s, ".0") {
				s = strings.TrimSuffix(s, ".0")
			} else if strings.HasSuffix(s, ".00") {
				s = strings.TrimSuffix(s, ".00")
			}
			aI, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return compareToFloat64(s, float64(to), opt, NewTypeError(r, "string", "int"))
			}
			if aI > to {
				return 1, nil
			}
			if aI < to {
				return -1, nil
			}
			return 0, nil
		}
	case ValueInt64:
		if r.IntValue() > to {
			return 1, nil
		}
		if r.IntValue() < to {
			return -1, nil
		}
		return 0, nil
	case ValueUint64:
		if to < 0 {
			return 1, nil
		}
		u := uint64(to)

		if r.UintValue() > u {
			return 1, nil
		}
		if r.UintValue() < u {
			return -1, nil
		}
		return 0, nil
	case ValueFloat64:
		u := float64(to)

		if r.FloatValue() > u {
			return 1, nil
		}
		return -1, nil
	}
	return 0, NewTypeError(r, r.Type.String(), "int")
}

func (r *Value) CompareToUint64(to uint64, opt CompareOption) (int, error) {
	switch r.Type {
	case ValueString:
		if opt.Weak {
			s := r.StrValue()
			if strings.HasSuffix(s, ".0") {
				s = strings.TrimSuffix(s, ".0")
			} else if strings.HasSuffix(s, ".00") {
				s = strings.TrimSuffix(s, ".00")
			}
			aU, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				a, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return compareToFloat64(s, float64(to), opt, NewTypeError(r, "string", "uint"))
				}
				if a < 0 {
					return -1, nil
				}
				aU = uint64(a)
			}
			if aU > to {
				return 1, nil
			}
			if aU < to {
				return -1, nil
			}
			return 0, nil
		}
	case ValueInt64:
		if r.IntValue() < 0 {
			return -1, nil
		}
		u := uint64(r.IntValue())
		if u > to {
			return 1, nil
		}
		if u < to {
			return -1, nil
		}
		return 0, nil
	case ValueUint64:
		if r.UintValue() > to {
			return 1, nil
		}
		if r.UintValue() < to {
			return -1, nil
		}
		return 0, nil
	case ValueFloat64:
		u := float64(to)

		if r.FloatValue() > u {
			return 1, nil
		}
		return -1, nil
	}
	return 0, NewTypeError(r, r.Type.String(), "uint")
}

func (r *Value) CompareToFloat64(to float64, opt CompareOption) (int, error) {
	switch r.Type {
	case ValueString:
		if opt.Weak {
			return compareToFloat64(r.StrValue(), to, opt, NewTypeError(r, "string", "int"))
		}
	case ValueInt64:
		u := float64(r.IntValue())
		if u > to {
			return 1, nil
		}
		return -1, nil
	case ValueUint64:
		u := float64(r.UintValue())
		if u > to {
			return 1, nil
		}
		return -1, nil
	case ValueFloat64:
		if r.FloatValue() > to {
			return 1, nil
		}
		return -1, nil
	}
	return 0, NewTypeError(r, r.Type.String(), "float")
}

func (r *Value) CompareToDatetime(to int64, opt CompareOption) (int, error) {
	var value int64
	switch r.Type {
	case ValueDatetime:
		value = r.IntValue()
	case ValueString:
		if !opt.Weak {
			return 0, NewTypeError(r, r.Type.String(), "datetime")
		}
		t, err := ToDatetime(r.StrValue())
		if err != nil {
			return 0, NewTypeError(r, r.Type.String(), "datetime")
		}
		value = DatetimeToInt(t)
	case ValueInt64:
		if !opt.Weak {
			return 0, NewTypeError(r, r.Type.String(), "datetime")
		}
		value = r.IntValue()
	case ValueUint64:
		if !opt.Weak {
			return 0, NewTypeError(r, r.Type.String(), "datetime")
		}
		value = int64(r.UintValue())
	case ValueFloat64:
		if !opt.Weak {
			return 0, NewTypeError(r, r.Type.String(), "datetime")
		}
		value = int64(r.FloatValue())
	default:
		return 0, NewTypeError(r, r.Type.String(), "datetime")
	}

	if value > to {
		return 1, nil
	}
	if value < to {
		return -1, nil
	}
	return 0, nil
}

func (r *Value) CompareToInterval(to time.Duration, opt CompareOption) (int, error) {
	var value time.Duration
	switch r.Type {
	case ValueInterval:
		value = r.DurationValue()
	case ValueString:
		if !opt.Weak {
			return 0, NewTypeError(r, r.Type.String(), "interval")
		}
		t, err := time.ParseDuration(r.StrValue())
		if err != nil {
			return 0, NewTypeError(r, r.Type.String(), "interval")
		}
		value = t
	// case ValueInt64:
	//  if !opt.Weak {
	//    return 0, NewTypeError(r, r.Type.String(), "interval")
	//  }
	//  value = r.IntValue()
	// case ValueUint64:
	//  if !opt.Weak {
	//    return 0, NewTypeError(r, r.Type.String(), "interval")
	//  }
	//  value = int64(r.UintValue())
	// case ValueFloat64:
	//  if !opt.Weak {
	//    return 0, NewTypeError(r, r.Type.String(), "interval")
	//  }
	//  value = int64(r.FloatValue())
	default:
		return 0, NewTypeError(r, r.Type.String(), "interval")
	}

	if value > to {
		return 1, nil
	}
	if value < to {
		return -1, nil
	}
	return 0, nil
}
