package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, config *Config) (Store, func(...string)) {
	t.Helper()

	s, err := New(config)
	if err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		defer s.Close()
		st := s.(*store)
		if len(tables) > 0 {
			if _, err := st.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
	}
}
