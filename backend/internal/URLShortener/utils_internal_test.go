package URLShortener

import "testing"

type URLTest struct {
	str string
	res bool
}

var URLTests = []struct {
	str string
	res bool
}{
	{"https://job.ozon.ru", true},
	{"http://yandex.ru", true},
	{"https://yandex.ru", true},
	{"http://google.com", true},
	{"http://google", false},
	{"yandex.ru", false},
	{"http://192.168.1.1", true},
	{"someurl", false},
	{"http://yandex.ru:80", true},
	{"ftp://127.0.0.1:80", true},
}

func TestIsCorrectURL(t *testing.T) {
	for _, test := range URLTests {
		if res := isURL(test.str); res != test.res {
			t.Errorf("isURL(\"%v\") wanted: %v, got : %v", test.str, test.res, res)
		}
	}
}
