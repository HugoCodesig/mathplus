package mathplus

import (
	mp "mathplus/internal/math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := mp.Add([]float64{2, 3, 5}...)
	expected := 10.0
	if result != expected {
		t.Errorf("Add() = %v; want %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := mp.Sub([]float64{10, 2, 3}...)
	expected := 5.0
	if result != expected {
		t.Errorf("Sub() = %v; want %v", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := mp.Mul([]float64{2, 3, 4}...)
	expected := 24.0
	if result != expected {
		t.Errorf("Mul() = %v; want %v", result, expected)
	}
}

func TestDiv(t *testing.T) {
	result, err := mp.Div(100, 2, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := 10.0
	if result != expected {
		t.Errorf("Div() = %v; want %v", result, expected)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := mp.Div(10, 0)
	if err == nil {
		t.Errorf("Expected error when dividing by zero, got nil")
	}
}

func TestAverage(t *testing.T) {
	result := mp.Average(2, 4, 6)
	expected := 4.0
	if result != expected {
		t.Errorf("Average() = %v; want %v", result, expected)
	}
}

func TestMedianOdd(t *testing.T) {
	result := mp.Median(3, 1, 5)
	expected := 3.0
	if result != expected {
		t.Errorf("Median() = %v; want %v", result, expected)
	}
}

func TestMode(t *testing.T) {
	result := mp.Mode(1, 2, 2, 3)
	if len(result) != 1 || result[0] != 2 {
		t.Errorf("Mode() = %v; want [2]", result)
	}
}

func TestIsPrime(t *testing.T) {
	if !mp.IsPrime(7) {
		t.Errorf("Expected 7 to be prime")
	}
	if mp.IsPrime(9) {
		t.Errorf("Expected 9 to be non-prime")
	}
}
