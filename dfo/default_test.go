package dfo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntPtr(t *testing.T) {
	num := 1
	num1 := IntPtr(&num)
	assert.Equal(t, num, num1)
	num2 := IntPtr(nil)
	assert.Equal(t, 0, num2)
}

func TestInt8Ptr(t *testing.T) {
	num := int8(1)
	num1 := Int8Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Int8Ptr(nil)
	assert.Equal(t, int8(0), num2)
}

func TestInt16Ptr(t *testing.T) {
	num := int16(1)
	num1 := Int16Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Int16Ptr(nil)
	assert.Equal(t, int16(0), num2)
}

func TestInt32Ptr(t *testing.T) {
	num := int32(1)
	num1 := Int32Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Int32Ptr(nil)
	assert.Equal(t, int32(0), num2)
}

func TestInt64Ptr(t *testing.T) {
	num := int64(1)
	num1 := Int64Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Int64Ptr(nil)
	assert.Equal(t, int64(0), num2)
}

func TestUintPtr(t *testing.T) {
	num := uint(1)
	num1 := UintPtr(&num)
	assert.Equal(t, num, num1)
	num2 := UintPtr(nil)
	assert.Equal(t, uint(0), num2)
}

func TestUint8Ptr(t *testing.T) {
	num := uint8(1)
	num1 := Uint8Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Uint8Ptr(nil)
	assert.Equal(t, uint8(0), num2)
}

func TestUint16Ptr(t *testing.T) {
	num := uint16(1)
	num1 := Uint16Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Uint16Ptr(nil)
	assert.Equal(t, uint16(0), num2)
}

func TestUint32Ptr(t *testing.T) {
	num := uint32(1)
	num1 := Uint32Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Uint32Ptr(nil)
	assert.Equal(t, uint32(0), num2)
}

func TestUint64Ptr(t *testing.T) {
	num := uint64(1)
	num1 := Uint64Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Uint64Ptr(nil)
	assert.Equal(t, uint64(0), num2)
}

func TestFloat32Ptr(t *testing.T) {
	num := float32(1)
	num1 := Float32Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Float32Ptr(nil)
	assert.Equal(t, float32(0), num2)
}

func TestFloat64Ptr(t *testing.T) {
	num := float64(1)
	num1 := Float64Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Float64Ptr(nil)
	assert.Equal(t, float64(0), num2)
}

func TestBoolPtr(t *testing.T) {
	num := true
	num1 := BoolPtr(&num)
	assert.Equal(t, num, num1)
	num2 := BoolPtr(nil)
	assert.Equal(t, false, num2)
}

func TestStringPtr(t *testing.T) {
	num := "1"
	num1 := StringPtr(&num)
	assert.Equal(t, num, num1)
	num2 := StringPtr(nil)
	assert.Equal(t, "", num2)
}

func TestComplex64Ptr(t *testing.T) {
	num := complex64(1)
	num1 := Complex64Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Complex64Ptr(nil)
	assert.Equal(t, complex64(0), num2)
}

func TestComplex128Ptr(t *testing.T) {
	num := complex128(1)
	num1 := Complex128Ptr(&num)
	assert.Equal(t, num, num1)
	num2 := Complex128Ptr(nil)
	assert.Equal(t, complex128(0), num2)
}
