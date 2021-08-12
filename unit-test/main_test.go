package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//package nya harus sama dengan yg mau di test
//kalau mau test ingat awalnya Test dan panggil *testing.T
func TestPenjumlahan(t *testing.T) {
	numA, numB := 5, 5
	result := Penjumlahan(numA, numB)
	resultSeharusnya := 10
	t.Log(numA, "+", numB, "=", result)

	if result != resultSeharusnya {
		t.Error("SALAH! Result Seharusnya", resultSeharusnya)
	}
}

//Benchmark untuk lihat performa dari function tersebut
//kalau mau Benchmark ingat awalnya Test dan panggil *testing.B
func BenchmarkPenjumlahan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Penjumlahan(5, 5)
	}
}

//pakai library go get github.com/stretchr/testify
func TestPenjumlahan2(t *testing.T) {
	// assert equality
	numA, numB := 5, 5
	result := Penjumlahan(numA, numB)
	resultSeharusnya := 10
	//equal kalau sama => PASS
	//notEqual bagi yg tidak sama => PASS
	//assert.Equal(t, result, resultSeharusnya, "NotEqual SALAH! Result Seharusnya", resultSeharusnya)
	assert.NotEqual(t, result, resultSeharusnya, " Equal SALAH! Result Seharusnya", resultSeharusnya)
}

//buat run pakai => go test
//atau => go test main_test.go main.go
//kalau mau lihat detail => go test main_test.go main.go -v
//kalau cuma mau test satu function aja => go test -run=TestPenjumlahan -v
//buat run benchmark => go test -bench=.
//kalau cuma mau benchmark satu function aja => go test -bench=Penjumlahan ||  go test -bench=BenchmarkPenjumlahan
