package dynamic_struct

import (
	"encoding/json"
	"reflect"
)

type Field struct {
	Name  string
	Value any
}

func ConstructStruct(currentFields map[string]any) any {
	var fields []reflect.StructField
	for k, v := range currentFields {
		fields = append(fields, reflect.StructField{
			Name: k,
			Type: reflect.TypeOf(v),
		})
	}
	v := reflect.New(reflect.StructOf(fields)).Elem()
	for i := 0; i < len(currentFields); i++ {
		v.Field(i).Set(reflect.ValueOf(currentFields[fields[i].Name]))
	}

	s := v.Addr().Elem().Interface()
	return s
}

func ReconstructStruct(obj any, fields ...Field) any {
	currencyFields := map[string]any{}
	currentType := reflect.TypeOf(obj)
	currentValue := reflect.ValueOf(obj)
	for i := 0; i < currentType.NumField(); i++ {
		field := currentType.Field(i)
		currencyFields[field.Name] = currentValue.Field(i).Interface()
	}
	for _, v := range fields {
		currencyFields[v.Name] = v.Value
	}
	return ConstructStruct(currencyFields)
}

func ToConcreteObject[T interface{}](obj any) *T {
	//TODO unsafe
	// (*T)(unsafe.Pointer(&obj) падает на out of memory при касте)
	d, _ := json.Marshal(obj)
	data := new(T)
	json.Unmarshal(d, data)
	return data
}
