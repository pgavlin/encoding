package json

import (
	"reflect"
	"unsafe"
)

type jsonValue interface {
	valueType() reflect.Type
}

type valueHeader struct {
	value   unsafe.Pointer
	defined bool
}

type Value[T any] struct {
	value   *T
	defined bool
}

func NullOf[T any]() Value[T] {
	return Value[T]{defined: true}
}

func ValueOf[T any](value T) Value[T] {
	return Value[T]{defined: true, value: &value}
}

func MaybeOf[T any](value *T) Value[T] {
	return Value[T]{defined: true, value: value}
}

func (v Value[T]) IsUndefined() bool {
	return !v.defined
}

func (v Value[T]) IsNull() bool {
	return v.defined && v.value == nil
}

func (v Value[T]) Value() *T {
	return v.value
}

func (v Value[T]) ValueOrZero() (value T) {
	if v.value == nil {
		return
	}
	return *v.value
}

func (v Value[T]) valueType() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func (v Value[T]) MarshalJSON() ([]byte, error) {
	return Marshal(v.value)
}

func (v *Value[T]) UnmarshalJSON(b []byte) error {
	v.defined = true
	return Unmarshal(b, &v.value)
}
