# go-basic

This package provides some basic utility functions.

## Usage

### IsNil

Identify whether the any type is nil.
```
	var emptyArray [1]int
	if IsNil(emptyArray) {
		t.Fatalf("GetOrCreate with emtpy slice failed")
	}
```
### NewIfEmpty

Make sure any type is created. Create by reflect if it is not there.
```
	var emptyAnyMap map[string]any
	fromEmptyAnyMap := NewIfEmpty(emptyAnyMap)
	if fromEmptyAnyMap == nil {
		t.Fatalf("GetOrCreate with emtpy map failed")
	}
	fromEmptyAnyMap["key"] = 1
```

## Contributing

PRs accepted.

## License

BSD-style © Barret Qin
