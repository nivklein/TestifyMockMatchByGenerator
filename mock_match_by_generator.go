package TestifyMockMatchByGenerator

import "reflect"

func CreateMatcher[T any](refObj T) func(T) bool {
	refVal := reflect.ValueOf(refObj)
	refType := refVal.Type()

	if refVal.Kind() != reflect.Struct {
		panic("CreateMatcher only supports struct types")
	}

	return func(input T) bool {
		inputVal := reflect.ValueOf(input)
		inputType := inputVal.Type()

		if inputType != refType {
			return false
		}

		for i := 0; i < refType.NumField(); i++ {
			field := refType.Field(i)
			fieldType := field.Type

			if fieldType.Kind() != reflect.Bool &&
				fieldType.Kind() != reflect.Int &&
				fieldType.Kind() != reflect.String {
				continue
			}

			if refVal.Field(i).Interface() != inputVal.Field(i).Interface() {
				return false
			}
		}
		return true
	}
}
