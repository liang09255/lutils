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

## conv
This package can convert the type of a variable to another type.

```go
go get -u  github.com/liang09255/lutils/conv
```

Example:

```go
str := "1.0"
conv.ToInt(str)     // 1 (int)
conv.ToUint(str)    // 1 (uint)
conv.ToFloat64(str) // 1 (float64)
str = "true"
conv.ToBool(str) 	// true (bool)
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

