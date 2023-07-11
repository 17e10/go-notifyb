# go-notifyb

[![GoDev][godev-image]][godev-url]

go-notifyb パッケージは通知を表す error を提供します.

## Usage

```go
const MyEOF notifyb.Notify = "my eof"

func Foo() error {
	return MyEOF
}
```

## License

This software is released under the MIT License, see LICENSE.

## Author

17e10

[godev-image]: https://pkg.go.dev/badge/github.com/17e10/go-notifyb
[godev-url]: https://pkg.go.dev/github.com/17e10/go-notifyb
