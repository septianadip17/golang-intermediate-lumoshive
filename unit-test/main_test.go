package main

import "testing"

// TestAdd menggunakan t.Error untuk menggagalkan tes
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Error("Expected 5, but got", result)
	}
}

// TestAddUsingErrorf menggunakan t.Errorf untuk menggagalkan tes dengan pesan berformat
func TestAddUsingErrorf(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// TestAddUsingFail menggunakan t.Fail untuk menggagalkan tes
func TestAddUsingFail(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Log("Expected 5, but got", result)
		t.Fail()
	}
}

// TestAddUsingFailNow menggunakan t.FailNow untuk menggagalkan tes dan menghentikan eksekusi
func TestAddUsingFailNow(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.FailNow()
	}
}

// TestAddUsingFatal menggunakan t.Fatal untuk menggagalkan tes dengan pesan
func TestAddUsingFatal(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Fatal("Expected 5, but got", result)
	}
}

// TestAddUsingFatalf menggunakan t.Fatalf untuk menggagalkan tes dengan pesan berformat
func TestAddUsingFatalf(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Fatalf("Add(2, 3) = %d; want %d", result, expected)
	}
}
