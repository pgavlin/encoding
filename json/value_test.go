package json

import (
	"reflect"
	"testing"
)

type ValueOptionals struct {
	Vr Value[int] `json:"vr"`
	Vo Value[int] `json:"vo,omitempty"`
}

var valueOptionalsExpected = `{
 "vr": null
}`

func TestValueOmitEmpty(t *testing.T) {
	var o ValueOptionals

	got, err := MarshalIndent(&o, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	if got := string(got); got != valueOptionalsExpected {
		t.Errorf(" got: %s\nwant: %s\n", got, valueOptionalsExpected)
	}
}

var valueUnmarshalExpected = ValueOptionals{
	Vr: NullOf[int](),
}

func TestValueUnmarshal(t *testing.T) {
	var got ValueOptionals

	if err := Unmarshal([]byte(valueOptionalsExpected), &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, valueUnmarshalExpected) {
		t.Errorf(" got: %v\nwant: %v\n", got, valueUnmarshalExpected)
	}
}

func TestValueBasics(t *testing.T) {
	var undefined Value[int]
	if !undefined.IsUndefined() {
		t.Errorf("zero value is not undefined")
	}
	if undefined.IsNull() {
		t.Errorf("zero value is null")
	}
	if undefined.Value() != nil {
		t.Errorf("zero value has a value")
	}
	if undefined.ValueOrZero() != 0 {
		t.Errorf("zero value has a value")
	}

	null := NullOf[int]()
	if null.IsUndefined() {
		t.Errorf("null value is undefined")
	}
	if !null.IsNull() {
		t.Errorf("null value is not null")
	}
	if null.Value() != nil {
		t.Errorf("null value has a value")
	}
	if null.ValueOrZero() != 0 {
		t.Errorf("null value has a value")
	}

	value := ValueOf(42)
	if value.IsUndefined() {
		t.Errorf("defined value is undefined")
	}
	if value.IsNull() {
		t.Errorf("defined value is null")
	}
	if value.Value() == nil {
		t.Errorf("defined value has no value")
	}
	if value.ValueOrZero() != 42 {
		t.Errorf("defined value has no value")
	}

	intV := 42
	maybe := MaybeOf(&intV)
	if maybe.IsUndefined() {
		t.Errorf("defined maybe is undefined")
	}
	if maybe.IsNull() {
		t.Errorf("defined maybe is null")
	}
	if maybe.Value() == nil {
		t.Errorf("defined maybe has no value")
	}
	if maybe.ValueOrZero() != 42 {
		t.Errorf("defined maybe has no value")
	}

	maybeNot := MaybeOf[int](nil)
	if maybeNot.IsUndefined() {
		t.Errorf("maybeNot value is undefined")
	}
	if !maybeNot.IsNull() {
		t.Errorf("maybeNot value is not null")
	}
	if maybeNot.Value() != nil {
		t.Errorf("maybeNot value has a value")
	}
	if maybeNot.ValueOrZero() != 0 {
		t.Errorf("maybeNot value has a value")
	}
}
