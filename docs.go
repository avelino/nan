/*
Package nan - Zero allocation Nullable structures in one library with handy conversion functions,
marshallers and unmarshallers.

Features:
- short name "nan"

Supported types:
- bool
- float32
- float64
- int
- int8
- int16
- int32
- int64
- string
- time.Time
- more types will be added at necessary

Supported marshallers:
- Standart JSON
- jsoniter
- easyjson
- Scylla and Cassandra (gocql, gocqlx)

Usage

Simply create struct field or variable with one of exported types and use it without any changes in external API.

JSON input/output will be converted to null or non null values. Scylla and Cassandra will
be save this variables correctly.

	var data struct {
		Code nan.NullString `json:"code"`
	}

	b, err := jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":null}
	fmt.Println(string(b))

	data.Code = nan.StringToNullString("1")
	// Or
	// data.Code = nan.NullString{String: "1", Valid: true}

	b, err = jsoniter.Marshal(data)
	if err != nil {
		panic(err)
	}

	// {"code":"1"}
	fmt.Println(string(b))
*/
package nan
