package reflects

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Fields(i interface{}) ([]reflect.StructField, error) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("unsupported type %s, only support struct type", t.Kind()))
	}

	var fields []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, field(t.Field(i))...)
	}
	return fields, nil
}

func field(f reflect.StructField) (fields []reflect.StructField) {
	t := f.Type

	if t == reflect.TypeOf(time.Time{}) {
		fields = append(fields, f)
		return
	}

	if t.Kind() != reflect.Struct {
		fields = append(fields, f)
		return
	}

	for i := 0; i < t.NumField(); i++ {
		ff := t.Field(i)
		fields = append(fields, field(ff)...)
	}

	return
}

func Values(i interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New(fmt.Sprintf("unsupported type %s, only support struct type", v.Kind()))
	}

	var values []interface{}
	for i := 0; i < v.NumField(); i++ {
		values = append(values, value(v.Field(i))...)
	}

	return values, nil
}

func value(v reflect.Value) (values []interface{}) {
	if v.Type() == reflect.TypeOf(time.Time{}) {
		values = append(values, v.Interface())
		return
	}

	if v.Kind() != reflect.Struct {
		values = append(values, v.Interface())
		return
	}

	for i := 0; i < v.NumField(); i++ {
		vv := v.Field(i)
		values = append(values, value(vv)...)
	}

	return
}

func FieldValue(i interface{}) (map[string]interface{}, error) {
	fields, err := Fields(i)
	if err != nil {
		return nil, err
	}

	values, err := Values(i)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{}, len(fields))

	for i, f := range fields {
		m[f.Name] = values[i]
	}

	return m, nil
}
