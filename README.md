# null

![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
![version](https://img.shields.io/badge/Go-v1.19-blue)

A zero dependency library to handle SQL nullable values.

## Contents

- [null](#null)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Null](#null-1)
    - [null.Bool](#nullbool)
    - [null.Byte](#nullbyte)
    - [null.Float64](#nullfloat64)
    - [null.Int16](#nullint16)
    - [null.Int32](#nullint32)
    - [null.Int64](#nullint64)
    - [null.String](#nullstring)
    - [null.Time](#nulltime)
  - [License](#license)

## Installation

1. You can use the below command to install `null`.

```bash
$ go get -u github.com/mobyfox/null
```

2. Import it in your code

```go
import "github.com/mobyfox/null"
```

## Null

### null.Bool

Nullable bool.

Marshals to JSON null if data coming from SQL is null. False will not produce a null Bool.

```go
type Example struct {
    Bool null.Bool `json:"bool"`
}
```

### null.Byte

Nullable byte.

Marshals to JSON null if data coming from SQL is null. Blank input will not produce a null Byte.

```go
type Example struct {
    Byte null.Byte `json:"byte"`
}
```

### null.Float64

Nullable float64.

Marshals to JSON null if data coming from SQL is null. Zero input will not produce a null Float64.

```go
type Example struct {
    Float64 null.Float64 `json:"float64"`
}
```

### null.Int16

Nullable int16.

Marshals to JSON null if data coming from SQL is null. Zero input will not produce a null Int16.

```go
type Example struct {
    Int16 null.Int16 `json:"int16"`
}
```

### null.Int32

Nullable int32.

Marshals to JSON null if data coming from SQL is null. Zero input will not produce a null Int32.

```go
type Example struct {
    Int32 null.Int32 `json:"int32"`
}
```

### null.Int64

Nullable int64.

Marshals to JSON null if data coming from SQL is null. Zero input will not produce a null Int64.

```go
type Example struct {
    Int64 null.Int64 `json:"int64"`
}
```

### null.String

Nullable string.

Marshals to JSON null if data coming from SQL is null. Blank input will not produce a null String.

```go
type Example struct {
    String null.String `json:"string"`
}
```

### null.Time

Nullable time.

Marshals to JSON null if data coming from SQL is null. Blank input will not produce a null Time.

```go
type Example struct {
    Time null.Time `json:"time"`
}
```

## License

This project utilizes the [Apache License 2.0](https://choosealicense.com/licenses/apache-2.0/) please do read the permissions, conditions and limitations before using this library.