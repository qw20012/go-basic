# go-error

This package can help define custom error and handle the error gracefully.

- Define custom error.
- Wrap the third party error.
- Wrap context informaton.
- Wrap additional trace information.
- Print statck trace message.

## Install
```
go get github.com/qw20012/go-error
```

## Usage

### Define custom error.

```
func devideByZero(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, err.New("Business.DevideByZero", "Given parameter divisor is zero in method devideByZero.")
	}
	return dividend / divisor, nil
}
```

### Wrap third pary error and relative context information.

```
func callThirdParty(id int, value int) (int, error) {
	r, e := thirdPartyMethod(id, value)
	if e != nil {
		params := make(map[string]any)
		params["id"] = id
		params["value"] = value
		err.New("ThirdParty.BadArgument", e, params)
	}
	return r, nil
}
```

### Wrap trace and relative context information.
```
func callDevideByZero(dividend int, divisor int) (int, error) {
	r, e := devideByZero(dividend, divisor)
	if e != nil {
		return 0, e.(*err.BqError).Wrap("Method callDevideByZero({dividend}, {divisor})").
		              WithParameter("dividend", dividend).WithParameter("divisor", divisor)
		 
	}
	return r, nil
}
```
### Check error kind and handle it.
```
func callAndHandleError() {
	_, e := callDevideByZero(2, 0)
	if e != nil {
		bqError := e.(*err.BqError)
		if bqError.Id() == "Business.DevideByZero" {
			fmt.Println("second is,: ", bqError)
		}
	}
}
```
## Contributing

PRs accepted.

## License

BSD-style © Barret Qin
