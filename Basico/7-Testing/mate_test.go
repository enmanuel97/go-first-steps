package testing

import "testing"

/*
func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 10 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 10)
	}
}*/

func TestSuma(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 10, 35},
	}

	for _, item := range table {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{2, 7, 7},
		{25, 10, 25},
	}

	for _, item := range table {
		mayor := GetMax(item.a, item.b)
		if mayor != item.c {
			t.Errorf("El numero mayor debe de ser: %d", item.c)
		}
	}
}

func TestFibo(t *testing.T) {
	table := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		fibo := Fibonacci(item.n)
		if fibo != item.r {
			t.Errorf("Serie Fibonacci incorrecta, se esperaba: %d y tiene %d", item.r, fibo)
		}
	}
}
