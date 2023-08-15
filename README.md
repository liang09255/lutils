# lutils

Some commonly used tools in development

## dfo
The name "dfo" is derived from the homophonic sound of "default".

This package can safely load pointer values and replace empty pointers with default values.

```go
go get -u  github.com/liang09255/lutils/dfo
```

Example:

```go
nums1, nums2 := 1, 1
addTwoNum := func(a *int, b *int) int {
    // This is a service which can add two numbers
    return dfo.IntPtr(a) + dfo.IntPtr(b)
}
addTwoNum(nil, nil)       // 0 dfo turn nil to 0
addTwoNum(&nums1, nil)    // 1
addTwoNum(nil, &nums2)    // 1
addTwoNum(&nums1, &nums2) // 2
```

Without dfo, we will implement the function like this, dfo can make the code look cleaner

```go
addTwoNum := func(a *int, b *int) int {
	num1, num2 := 0, 0
	if a != nil {
		num1 = *a 
	}
	if b != nil {
		num2 = *b
	}
	return num1 + num2
}
```



## conv

This package can convert the type of a variable to another type.

We can use `conv.SetLogger` to set logger, conv will automatically collect errors into logger

```go
go get -u  github.com/liang09255/lutils/conv
```

Example:

```go
str := "1.0"
i1 := conv.ToInt(str) // 1 (int)
ui1 := conv.ToUint(str) // 1 (uint)
f1 := conv.ToFloat64(str) // 1 (float64)

str = "true"
b := conv.ToBool(str) // true (bool)

m := map[string]string{"a": "b"}
jsonContent := conv.ToJSON(m) // {"a":"b"}

arr := []string{"a", "b", "c"}
m1 := conv.ArrayToMap(arr) // map[a:{} b:{} c:{}]
_, ok1 := m1["a"] // ok1 => true
_, ok2 := m1["d"] // ok2 => false
m2 := conv.ArrayToBoolMap(arr)
ok3 := m2["a"] // true
ok4 := m2["d"] // false
```

```go
// use logger
conv.SetLogger(Logger)
// 	logger need to implement interface Logger
// 	type Logger interface {
// 		Info(args ...interface{})
//		Warn(args ...interface{})
//		Error(args ...interface{})
//	}
```

