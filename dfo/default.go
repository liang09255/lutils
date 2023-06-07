package dfo

func IntPtr(p *int) int {
	return ptrDefault(p, 0)
}

func Int8Ptr(p *int8) int8 {
	return ptrDefault(p, 0)
}

func Int16Ptr(p *int16) int16 {
	return ptrDefault(p, 0)
}

func Int32Ptr(p *int32) int32 {
	return ptrDefault[int32](p, 0)
}

func Int64Ptr(p *int64) int64 {
	return ptrDefault(p, 0)
}

func UintPtr(p *uint) uint {
	return ptrDefault(p, 0)
}

func Uint8Ptr(p *uint8) uint8 {
	return ptrDefault(p, 0)
}

func Uint16Ptr(p *uint16) uint16 {
	return ptrDefault(p, 0)
}

func Uint32Ptr(p *uint32) uint32 {
	return ptrDefault(p, 0)
}

func Uint64Ptr(p *uint64) uint64 {
	return ptrDefault(p, 0)
}

func Float32Ptr(p *float32) float32 {
	return ptrDefault(p, 0)
}

func Float64Ptr(p *float64) float64 {
	return ptrDefault(p, 0)
}

func BoolPtr(p *bool) bool {
	return ptrDefault(p, false)
}

func StringPtr(p *string) string {
	return ptrDefault(p, "")
}

func Complex64Ptr(p *complex64) complex64 {
	return ptrDefault(p, 0)
}

func Complex128Ptr(p *complex128) complex128 {
	return ptrDefault(p, 0)
}

func ptrDefault[T any](s *T, def T) T {
	if s == nil {
		return def
	}
	return *s
}
