package conv

import (
	"encoding/json"
	"strconv"
)

// TODO add log

func ToInt(v interface{}) int {
	switch t := v.(type) {
	case int:
		return t
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	case float32:
		return int(t)
	case float64:
		return int(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToInt error:", err)
			return 0
		}
		return int(value)
	}
	return 0
}

func ToInt8(v interface{}) int8 {
	switch t := v.(type) {
	case int:
		return int8(t)
	case int8:
		return t
	case int16:
		return int8(t)
	case int32:
		return int8(t)
	case int64:
		return int8(t)
	case uint:
		return int8(t)
	case uint8:
		return int8(t)
	case uint16:
		return int8(t)
	case uint32:
		return int8(t)
	case uint64:
		return int8(t)
	case float32:
		return int8(t)
	case float64:
		return int8(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToInt8 error:", err)
			return 0
		}
		return int8(value)
	}
	return 0
}

func ToInt16(v interface{}) int16 {
	switch t := v.(type) {
	case int:
		return int16(t)
	case int8:
		return int16(t)
	case int16:
		return t
	case int32:
		return int16(t)
	case int64:
		return int16(t)
	case uint:
		return int16(t)
	case uint8:
		return int16(t)
	case uint16:
		return int16(t)
	case uint32:
		return int16(t)
	case uint64:
		return int16(t)
	case float32:
		return int16(t)
	case float64:
		return int16(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToInt16 error:", err)
			return 0
		}
		return int16(value)
	}
	return 0
}

func ToInt32(v interface{}) int32 {
	switch t := v.(type) {
	case int:
		return int32(t)
	case int8:
		return int32(t)
	case int16:
		return int32(t)
	case int32:
		return t
	case int64:
		return int32(t)
	case uint:
		return int32(t)
	case uint8:
		return int32(t)
	case uint16:
		return int32(t)
	case uint32:
		return int32(t)
	case uint64:
		return int32(t)
	case float32:
		return int32(t)
	case float64:
		return int32(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToInt32 error:", err)
			return 0
		}
		return int32(value)
	}
	return 0
}

func ToInt64(v interface{}) int64 {
	switch t := v.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToInt64 error:", err)
			return 0
		}
		return value
	}
	return 0
}

func ToUint(v interface{}) uint {
	switch t := v.(type) {
	case int:
		return uint(t)
	case int8:
		return uint(t)
	case int16:
		return uint(t)
	case int32:
		return uint(t)
	case int64:
		return uint(t)
	case uint:
		return t
	case uint8:
		return uint(t)
	case uint16:
		return uint(t)
	case uint32:
		return uint(t)
	case uint64:
		return uint(t)
	case float32:
		return uint(t)
	case float64:
		return uint(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToUint error:", err)
			return 0
		}
		return uint(value)
	}
	return 0
}

func ToUint8(v interface{}) uint8 {
	switch t := v.(type) {
	case int:
		return uint8(t)
	case int8:
		return uint8(t)
	case int16:
		return uint8(t)
	case int32:
		return uint8(t)
	case int64:
		return uint8(t)
	case uint:
		return uint8(t)
	case uint8:
		return t
	case uint16:
		return uint8(t)
	case uint32:
		return uint8(t)
	case uint64:
		return uint8(t)
	case float32:
		return uint8(t)
	case float64:
		return uint8(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToUint8 error:", err)
			return 0
		}
		return uint8(value)
	}
	return 0
}

func ToUint16(v interface{}) uint16 {
	switch t := v.(type) {
	case int:
		return uint16(t)
	case int8:
		return uint16(t)
	case int16:
		return uint16(t)
	case int32:
		return uint16(t)
	case int64:
		return uint16(t)
	case uint:
		return uint16(t)
	case uint8:
		return uint16(t)
	case uint16:
		return t
	case uint32:
		return uint16(t)
	case uint64:
		return uint16(t)
	case float32:
		return uint16(t)
	case float64:
		return uint16(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToUint16 error:", err)
			return 0
		}
		return uint16(value)
	}
	return 0
}

func ToUint32(v interface{}) uint32 {
	switch t := v.(type) {
	case int:
		return uint32(t)
	case int8:
		return uint32(t)
	case int16:
		return uint32(t)
	case int32:
		return uint32(t)
	case int64:
		return uint32(t)
	case uint:
		return uint32(t)
	case uint8:
		return uint32(t)
	case uint16:
		return uint32(t)
	case uint32:
		return t
	case uint64:
		return uint32(t)
	case float32:
		return uint32(t)
	case float64:
		return uint32(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToUint32 error:", err)
			return 0
		}
		return uint32(value)
	}
	return 0
}

func ToUint64(v interface{}) uint64 {
	switch t := v.(type) {
	case int:
		return uint64(t)
	case int8:
		return uint64(t)
	case int16:
		return uint64(t)
	case int32:
		return uint64(t)
	case int64:
		return uint64(t)
	case uint:
		return uint64(t)
	case uint8:
		return uint64(t)
	case uint16:
		return uint64(t)
	case uint32:
		return uint64(t)
	case uint64:
		return t
	case float32:
		return uint64(t)
	case float64:
		return uint64(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			convLogger.Error("conv.ToUint64 error:", err)
			return 0
		}
		return value
	}
	return 0
}

func ToFloat32(v interface{}) float32 {
	switch t := v.(type) {
	case int:
		return float32(t)
	case int8:
		return float32(t)
	case int16:
		return float32(t)
	case int32:
		return float32(t)
	case int64:
		return float32(t)
	case uint:
		return float32(t)
	case uint8:
		return float32(t)
	case uint16:
		return float32(t)
	case uint32:
		return float32(t)
	case uint64:
		return float32(t)
	case float32:
		return t
	case float64:
		return float32(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseFloat(t, 64)
		if err != nil {
			convLogger.Error("conv.ToFloat32 error:", err)
			return 0
		}
		return float32(value)
	}
	return 0
}

func ToFloat64(v interface{}) float64 {
	switch t := v.(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return t
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		value, err := strconv.ParseFloat(t, 64)
		if err != nil {
			convLogger.Error("conv.ToFloat64 error:", err)
			return 0
		}
		return value
	}
	return 0
}

// ToString
// Attention: it will convert bool to true/false in lowercase
func ToString(v interface{}) string {
	switch t := v.(type) {
	case int:
		return strconv.FormatInt(int64(t), 10)
	case int8:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint8:
		return strconv.FormatUint(uint64(t), 10)
	case uint16:
		return strconv.FormatUint(uint64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case bool:
		if t {
			return "true"
		}
		return "false"
	case string:
		return t
	}
	return ""
}

func ToBool(v interface{}) bool {
	switch t := v.(type) {
	case int:
		return t != 0
	case int8:
		return t != 0
	case int16:
		return t != 0
	case int32:
		return t != 0
	case int64:
		return t != 0
	case uint:
		return t != 0
	case uint8:
		return t != 0
	case uint16:
		return t != 0
	case uint32:
		return t != 0
	case uint64:
		return t != 0
	case float32:
		return t != 0
	case float64:
		return t != 0
	case bool:
		return t
	case string:
		value, err := strconv.ParseBool(t)
		if err != nil {
			convLogger.Error("conv.ToBool error:", err)
			return false
		}
		return value
	}
	return false
}

func ToJSON(v interface{}) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		convLogger.Error("conv.ToJSON error:", err)
		return ""
	}
	return string(marshal)
}
