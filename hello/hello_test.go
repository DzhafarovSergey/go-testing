package hello

import (
	"errors"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	//Проверка пустой строки
	cases := []struct {
		name string
		want string
		err  error
	}{
		{"Anton", "Hello, Anton!", nil},
		{"Sergey", "Hello, Sergey!", nil},
		{"", "", errors.New("name is empty")},
	}

	for _, c := range cases {
		name_test := fmt.Sprintf("%s_%s", c.name, c.want)
		t.Run(name_test, func(t *testing.T) {
			got, err := Hello(c.name)
			if c.err != nil {
				if err == nil || err.Error() != c.err.Error() {
					t.Errorf("got %s, want %s", err, c.err)
				}
			}

			if got != c.want {
				t.Errorf("got %s, want %s", got, c.want)
			}

		})
	}
}
