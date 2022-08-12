package ref

import (
	"reflect"
	"strconv"
)

func Parse(value any) (reflect.Type, reflect.Kind, reflect.Value) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	k := t.Kind()

	if k == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
		k = t.Kind()
	}
	return t, k, v
}

// Get the exact value of given reflect.Value.
func GetValue[T any](value reflect.Value) T {
	return value.Interface().(T)
}

// Identify whether the any type is nil.
func IsNil(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return true
	}

	kind := v.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && v.IsNil() {
		return true
	}

	if kind == reflect.Ptr {
		return IsNil(reflect.Indirect(v).Interface())
	}

	/*
		switch reflect.TypeOf(value).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(value).IsNil()
		}*/

	return false
}

// Identify whether the any type is zero.
func IsZero(value any) bool {
	if IsNil(value) {
		return true
	}

	v := reflect.ValueOf(value)

	t := v.Type()
	switch t.Kind() {
	case reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Slice:
		s := reflect.MakeSlice(t, 0, 0)
		return reflect.DeepEqual(v.Interface(), s.Interface())
	case reflect.Ptr:
		//return reflect.Indirect(v).IsZero()
		if v.IsNil() {
			return true
		}
		deref := v.Elem().Interface()
		return IsZero(deref)
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(t).Interface())
	}
}

// Make sure any type is created. Create by reflect if it is not there.
func NewIfEmpty[T any](value T) T {
	if !IsNil(value) {
		return value
	}

	newValue := ReflectNew(value)
	//creatd := newValue.Interface().(T)
	creatd := GetValue[T](newValue)
	return creatd
}

// Todo: Would complete this function in future.
func ReflectNew(value any) reflect.Value {
	ty := reflect.TypeOf(value)
	switch ty.Kind() {
	case reflect.Interface:
		return reflect.MakeMap(ty).Elem()
	case reflect.Ptr:
		return reflect.New(ty.Elem()).Elem()
	case reflect.Map:
		return reflect.MakeMap(ty)
	case reflect.Struct:
		return reflect.New(ty.Elem()).Elem()
	case reflect.Slice:
		return reflect.MakeSlice(ty, 0, 0)
	}

	return reflect.ValueOf(value)
}

// Help set struct basic field by given value
func SetStructBasicField(field reflect.Value, value any) {
	strValue := value.(string)

	switch field.Type().Kind() {
	case reflect.String:
		field.SetString(strValue)
	case reflect.Bool:
		if realValue, err := strconv.ParseBool(strValue); err == nil {
			field.SetBool(realValue)
		}
	case reflect.Float32:
		if realValue, err := strconv.ParseFloat(strValue, 32); err == nil {
			field.SetFloat(realValue)
		}
	case reflect.Float64:
		if realValue, err := strconv.ParseFloat(strValue, 64); err == nil {
			field.SetFloat(realValue)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if realValue, err := strconv.Atoi(strValue); err == nil {
			field.SetInt(int64(realValue))
		}
	}
}

// Help set struct basic pointer field by given value
func SetStructBasicPtrField(field reflect.Value, value any) {
	strValue := value.(string)

	switch field.Type().Elem().Kind() {
	case reflect.String:
		field.Set(reflect.ValueOf(&strValue))

	case reflect.Int:
		if realValue, err := strconv.Atoi(strValue); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}

	case reflect.Int8:
		if realValue, err := strconv.ParseInt(strValue, 10, 8); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Int16:
		if realValue, err := strconv.ParseInt(strValue, 10, 16); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Int32:
		if realValue, err := strconv.ParseInt(strValue, 10, 32); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Int64:
		if realValue, err := strconv.ParseInt(strValue, 10, 64); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Bool:
		if realValue, err := strconv.ParseBool(strValue); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Float32:
		if realValue, err := strconv.ParseFloat(strValue, 32); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	case reflect.Float64:
		if realValue, err := strconv.ParseFloat(strValue, 64); err == nil {
			field.Set(reflect.ValueOf(&realValue))
		}
	}
}

// Help append slice with given basic element value
func AppendSliceBasicElem(aSlice reflect.Value, value reflect.Value) reflect.Value {
	switch aSlice.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if realValue, err := strconv.Atoi(value.Interface().(string)); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(realValue))
		}

	case reflect.String:
		if realValue, ok := value.Interface().(string); ok {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(realValue))
		}
	case reflect.Bool:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseBool(strValue); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(realValue))
		}
	case reflect.Float32:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseFloat(strValue, 32); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(realValue))
		}
	case reflect.Float64:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseFloat(strValue, 64); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(realValue))
		}
	}
	return aSlice
}

// Help append slice with given basic pointer element value
func AppendSliceBasicPtrElem(aSlice reflect.Value, value reflect.Value) reflect.Value {
	switch aSlice.Type().Elem().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if realValue, err := strconv.Atoi(value.Interface().(string)); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(&realValue))
		}
	case reflect.String:
		if realValue, ok := value.Interface().(string); ok {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(&realValue))
		}
	case reflect.Bool:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseBool(strValue); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(&realValue))
		}
	case reflect.Float32:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseFloat(strValue, 32); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(&realValue))
		}
	case reflect.Float64:
		strValue := value.Interface().(string)
		if realValue, err := strconv.ParseFloat(strValue, 64); err == nil {
			aSlice = reflect.Append(aSlice, reflect.ValueOf(&realValue))
		}
	}
	return aSlice
}
