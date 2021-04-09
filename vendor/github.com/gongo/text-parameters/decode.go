package parameters

import (
	"io"
	"reflect"
	"strconv"
)

type Decoder struct {
	reader io.Reader
}

func NewDecorder(r io.Reader) *Decoder {
	return &Decoder{reader: r}
}

func (d *Decoder) Decode(v interface{}) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return &CodingStructPointerError{rv}
	}

	rv = rv.Elem()

	params, err := Unmarshal(d.reader)
	if err != nil {
		return err
	}

	for i := 0; i < rv.NumField(); i++ {
		if err := d.store(rv.Type().Field(i), rv.Field(i), params); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) store(rt reflect.StructField, rv reflect.Value, params *TextParameters) error {
	if rt.PkgPath != "" { // unexported field
		return nil
	}

	keyName := rt.Name
	if rt.Tag.Get("parameters") != "" {
		keyName = rt.Tag.Get("parameters")
	}

	body := params.Get(keyName)

	if body == "" {
		return nil
	}

	decodeError := &DecodeFieldTypeError{t: rt, body: body}

	switch rt.Type.Kind() {
	case reflect.String:
		rv.SetString(body)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(body, 10, 64)
		if err != nil {
			return decodeError
		}
		rv.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, err := strconv.ParseUint(body, 10, 64)
		if err != nil {
			return decodeError
		}
		rv.SetUint(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(body, rv.Type().Bits())
		if err != nil {
			return decodeError
		}
		rv.SetFloat(n)
	}

	return nil
}
