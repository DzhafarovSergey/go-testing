package hello

import "testing"

func TestHello(t *testing.T) {
	//Проверка пустой строки
	s := Hello("")
	s_success := "Hello, world"

	if s != s_success {
		t.Errorf("Hello(\"\") = %s; want %s", s, s_success)
	}
}
