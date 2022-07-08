package math

func Penjumlahan(angka ...int) int {
	var result int
	for _, a := range angka {
		result += a
	}
	return result
}

func Pengurangan(angka ...int) int {
	var result int
	for _, a := range angka {
		result -= a
	}
	return result
}
