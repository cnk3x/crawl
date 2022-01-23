package crawl

import (
	"github.com/goccy/go-yaml"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	d1, _ := time.ParseDuration("720h")
	t.Log(d1, Duration(d1))
	d, err := ParseDuration("30d")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d)

	var st struct {
		D1 Duration `json:"d2"`
	}
	st.D1 = Duration(time.Hour * 720)
	s, err := yaml.Marshal(st)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(s))

	s = []byte(`d2: "30d2h5s"
`)

	st.D1 = 0
	if err = yaml.Unmarshal(s, &st); err != nil {
		t.Fatal(err)
	}
	t.Log(st.D1)
}
