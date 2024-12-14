package localization

import "testing"

func TestGet(t *testing.T) {
	schema := newSchema()
	schema.load(file)

	tcases := []struct {
		lang string
		key  string
		want string
	}{
		{"en", "not_found", "Not found"},
		{"fa", "not_found", "پیدا نشد"},
	}

	for _, tcase := range tcases {
		if got := Get(tcase.lang, tcase.key); got != tcase.want {
			t.Errorf("Get() = %v, want %v", got, tcase.want)
		}
	}

	return
}
