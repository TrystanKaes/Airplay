package parameters

import (
	"io"
	"reflect"
	"strconv"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

func (e *Encoder) Encode(v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return &CodingStructPointerError{rv}
	}

	rv = rv.Elem()
	params := &TextParameters{}

	for i := 0; i < rv.NumField(); i++ {
		e.encode(rv.Type().Field(i), rv.Field(i), params)
	}

	encoded := Marshal(params)
	_, err := io.WriteString(e.writer, encoded)

	return err
}

func (e *Encoder) encode(rt reflect.StructField, rv reflect.Value, params *TextParameters) {
	if rt.PkgPath != "" { // unexported field
		return
	}

	keyName := rt.Name
	if rt.Tag.Get("parameters") != "" {
		keyName = rt.Tag.Get("parameters")
	}

	switch rt.Type.Kind() {
	case reflect.String:
		params.Set(keyName, rv.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		params.Set(keyName, strconv.FormatInt(rv.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		params.Set(keyName, strconv.FormatUint(rv.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		value := strconv.FormatFloat(rv.Float(), 'g', -1, rv.Type().Bits())
		params.Set(keyName, value)
	}
}
