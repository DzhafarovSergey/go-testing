package hello

import (
	"errors"
	//"fmt"
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if (got != nil) != want {
		if want {
			t.Fatalf("ожидалась ошибка, но её нет")
		}
		t.Fatalf("не ожидалась ошибка, но получена: %v", got)
	}
}

func TestHello(t *testing.T) {
	//Проверка пустой строки
	cases := []struct {
		name      string
		want      string
		test_name string
		err       error
	}{
		{"Go", "Hello, Go", "basic", nil},
		{"World", "Hello, World", "basic", nil},
		{"", "", "empty", errors.New("name is empty")},
		{"Гофер", "Hello, Гофер", "unicode", nil},
	}
	for _, c := range cases {
		t.Run(c.test_name, func(t *testing.T) {
			got, err := Hello(c.name)
			w := false
			if c.name == "" {
				w = true
			}
			assertError(t, err, w)
			assertString(t, got, c.want)
		})
	}
}
