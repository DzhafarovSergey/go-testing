package reverse

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestReverce(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{input: "hello world", want: "dlrow olleh"},
		{"go", "og"},
		{"abba", "abba"},
		{"", ""},
		{"😉🙂", "🙂😉"},
		{"😉", "😉"},
		{"a", "a"},
	}
	for _, c := range cases {
		got := Reverse(c.input)
		assertString(t, got, c.want)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Reverse("hello world")
	}
}
