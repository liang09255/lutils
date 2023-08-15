package conv

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	stringOneNum    = "1"
	stringTrueFlag  = "true"
	stringFalseFlag = "false"

	intOneNum   = 1
	intTrueNum  = 1
	intFalseNum = 0

	float64OneNum   = 1.0
	float64TrueNum  = 1.0
	float64FalseNum = 0.0

	boolTrueFlag  = true
	boolFalseFlag = false
)

func TestToInt(t *testing.T) {
	assert.Equal(t, intOneNum, ToInt(stringOneNum))
	assert.Equal(t, intOneNum, ToInt(intOneNum))
	assert.Equal(t, intOneNum, ToInt(float64OneNum))
	assert.Equal(t, intTrueNum, ToInt(boolTrueFlag))
	assert.Equal(t, intFalseNum, ToInt(boolFalseFlag))
}

func TestToFloat64(t *testing.T) {
	assert.Equal(t, float64OneNum, ToFloat64(stringOneNum))
	assert.Equal(t, float64OneNum, ToFloat64(intOneNum))
	assert.Equal(t, float64OneNum, ToFloat64(float64OneNum))
	assert.Equal(t, float64TrueNum, ToFloat64(boolTrueFlag))
	assert.Equal(t, float64FalseNum, ToFloat64(boolFalseFlag))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, boolTrueFlag, ToBool(stringTrueFlag))
	assert.Equal(t, boolTrueFlag, ToBool(intTrueNum))
	assert.Equal(t, boolTrueFlag, ToBool(float64TrueNum))
	assert.Equal(t, boolFalseFlag, ToBool(stringFalseFlag))
	assert.Equal(t, boolFalseFlag, ToBool(intFalseNum))
	assert.Equal(t, boolFalseFlag, ToBool(float64FalseNum))
}

func TestToString(t *testing.T) {
	assert.Equal(t, stringOneNum, ToString(stringOneNum))
	assert.Equal(t, stringOneNum, ToString(intOneNum))
	assert.Equal(t, stringOneNum, ToString(float64OneNum))
	assert.Equal(t, stringTrueFlag, ToString(boolTrueFlag))
	assert.Equal(t, stringFalseFlag, ToString(boolFalseFlag))
}

func TestToJSON(t *testing.T) {
	assert.Equal(t, `{"a":"b"}`, ToJSON(map[string]string{"a": "b"}))
	assert.Equal(t, `{"a":"b"}`, ToJSON(struct {
		A string `json:"a"`
	}{A: "b"}))
}

func TestArray2Map(t *testing.T) {
	arr := []string{"a", "b", "c"}
	m := ArrayToMap(arr)
	log.Printf("%+v", m)
	assert.Equal(t, 3, len(m))
	_, okA := m["a"]
	_, okB := m["b"]
	_, okC := m["c"]
	_, okD := m["d"]
	assert.True(t, okA)
	assert.True(t, okB)
	assert.True(t, okC)
	assert.False(t, okD)
}

func TestArrayToBoolMap(t *testing.T) {
	arr := []string{"a", "b", "c"}
	m := ArrayToBoolMap(arr)
	assert.Equal(t, 3, len(m))
	assert.True(t, m["a"])
	assert.True(t, m["b"])
	assert.True(t, m["c"])
	assert.False(t, m["d"])
}
