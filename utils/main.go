package utils

import (
	"fmt"
	"reflect"
)

func GetDataWithType(typeData string, t interface{}) interface{} {
	var result interface{}
	var (
		I        string = "int"
		IArray   string = "[]int"
		I8       string = "int8"
		I8Array  string = "[]int8"
		I16      string = "int16"
		I16Array string = "[]int16"
		I32      string = "int32"
		I32Array string = "[]int32"
		I64      string = "int64"
		// I64Array  string = "[]int64"
		UI        string = "uint"
		UIArray   string = "[]uint"
		UI8       string = "uint8"
		UI8Array  string = "[]uint8"
		UI16      string = "uint16"
		UI16Array string = "[]uint16"
		UI32      string = "uint32"
		UI32Array string = "[]uint32"
		UI64      string = "uint64"
		UI64Array string = "[]uint64"
		F32       string = "float32"
		F32Array  string = "[]float32"
		F64       string = "float64"
		F64Array  string = "[]float64"
		B         string = "bool"
		BArray    string = "[]bool"
		R         string = "rune"
		RArray    string = "[]rune"
		C64       string = "complex64"
		C64Array  string = "[]complex64"
		C128      string = "complex128"
		C128Array string = "[]complex128"
		Str       string = "string"
		StrArray  string = "[]string"
	)
	value := reflect.ValueOf(t)
	switch typeData {
	case I:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(int))
	case IArray:
		rs := value.Interface().([]int)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case I8:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(int8))
	case I8Array:
		rs := value.Interface().([]int8)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case I16:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(int16))
	case I16Array:
		rs := value.Interface().([]int16)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case I32:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(int32))
	case I32Array:
		rs := value.Interface().([]int32)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case I64:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(int64))
	case UI64Array:
		rs := value.Interface().([]int64)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case UI:
		result = fmt.Sprintf("\033[94m%v\033[0m", value.Interface().(uint))
	case UIArray:
		rs := value.Interface().([]uint)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case UI8:
		result = value.Interface().(uint8)
	case UI8Array:
		rs := value.Interface().([]uint8)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case UI16:
		result = value.Interface().(uint16)
	case UI16Array:
		rs := value.Interface().([]uint16)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case UI32:
		result = value.Interface().(uint32)
	case UI32Array:
		rs := value.Interface().([]uint32)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case UI64:
		result = value.Interface().(uint64)
	case UI64Array:
		rs := value.Interface().([]uint64)
		res := []string{}
		for I, item := range rs {
			tr := "\033[94m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf("\033[94m%v\033[0m,", item))
		}
		result = res
	case F32:
		result = value.Interface().(float32)
	case F32Array:
		rs := value.Interface().([]float32)
		res := []string{}
		for I, item := range rs {
			tr := "\033[93m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case F64:
		result = value.Interface().(float64)
	case F64Array:
		rs := value.Interface().([]float64)
		res := []string{}
		for I, item := range rs {
			tr := "\033[93m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case B:
		result = fmt.Sprintf("\033[35m%v\033[0m", value.Interface().(bool))
	case BArray:
		rs := value.Interface().([]bool)
		res := []string{}
		for I, item := range rs {
			tr := "\033[35m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case R:
		result = value.Interface().(rune)
	case RArray:
		rs := value.Interface().([]rune)
		res := []string{}
		for I, item := range rs {
			tr := "\033[35m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case C64:
		result = value.Interface().(complex64)
	case C64Array:
		rs := value.Interface().([]complex64)
		res := []string{}
		for I, item := range rs {
			tr := "\033[35m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case C128:
		result = value.Interface().(complex128)
	case C128Array:
		rs := value.Interface().([]complex128)
		res := []string{}
		for I, item := range rs {
			tr := "\033[35m%v\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	case Str:
		result = fmt.Sprintf("\033[92m'%v'\033[0m", value.Interface().(string))
	case StrArray:
		rs := value.Interface().([]complex64)
		res := []string{}
		for I, item := range rs {
			tr := "\033[92m'%v'\033[0m"
			if I != len(rs)-1 {
				tr += ","
			}
			res = append(res, fmt.Sprintf(tr, item))
		}
		result = res
	default:
		result = "\033[94m [ Object ] \033[0m"
	}
	return result
}
