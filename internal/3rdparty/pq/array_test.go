package pq

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestBoolArrayValue(t *testing.T) {
	result, err := BoolArray(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = BoolArray([]bool{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = BoolArray([]bool{false, true, false}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{f,t,f}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkBoolArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([]bool, 10)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(2) == 0
	}
	a := BoolArray(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}

func TestByteaArrayValue(t *testing.T) {
	result, err := ByteaArray(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = ByteaArray([][]byte{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = ByteaArray([][]byte{{'\xDE', '\xAD', '\xBE', '\xEF'}, {'\xFE', '\xFF'}, {}}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{"\\xdeadbeef","\\xfeff","\\x"}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkByteaArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([][]byte, 10)
	for i := 0; i < len(x); i++ {
		x[i] = make([]byte, len(x))
		for j := 0; j < len(x); j++ {
			x[i][j] = byte(rand.Int())
		}
	}
	a := ByteaArray(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}

func TestFloat64ArrayValue(t *testing.T) {
	result, err := Float64Array(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = Float64Array([]float64{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = Float64Array([]float64{1.2, 3.4, 5.6}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{1.2,3.4,5.6}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkFloat64ArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([]float64, 10)
	for i := 0; i < len(x); i++ {
		x[i] = rand.NormFloat64()
	}
	a := Float64Array(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}

func TestInt64ArrayValue(t *testing.T) {
	result, err := Int64Array(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = Int64Array([]int64{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = Int64Array([]int64{1, 2, 3}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{1,2,3}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkInt64ArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([]int64, 10)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int63()
	}
	a := Int64Array(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}

func TestFloat32ArrayValue(t *testing.T) {
	result, err := Float32Array(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = Float32Array([]float32{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = Float32Array([]float32{1.2, 3.4, 5.6}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{1.2,3.4,5.6}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkFloat32ArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([]float32, 10)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Float32()
	}
	a := Float32Array(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}
func TestInt32ArrayValue(t *testing.T) {
	result, err := Int32Array(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = Int32Array([]int32{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = Int32Array([]int32{1, 2, 3}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{1,2,3}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkInt32ArrayValue(b *testing.B) {
	rand.Seed(1)
	x := make([]int32, 10)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int31()
	}
	a := Int32Array(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}

func TestStringArrayValue(t *testing.T) {
	result, err := StringArray(nil).Value()

	if err != nil {
		t.Fatalf("Expected no error for nil, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil, got %q", result)
	}

	result, err = StringArray([]string{}).Value()

	if err != nil {
		t.Fatalf("Expected no error for empty, got %v", err)
	}
	if expected := `{}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty, got %q", result)
	}

	result, err = StringArray([]string{`a`, `\b`, `c"`, `d,e`}).Value()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if expected := `{"a","\\b","c\"","d,e"}`; !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func BenchmarkStringArrayValue(b *testing.B) {
	x := make([]string, 10)
	for i := 0; i < len(x); i++ {
		x[i] = strings.Repeat(`abc"def\ghi`, 5)
	}
	a := StringArray(x)

	for i := 0; i < b.N; i++ {
		a.Value()
	}
}
