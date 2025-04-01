package count_test

import (
	"bytes"
	"count"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"count": func() {
			count.Main()
		},
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
		/*
			Setup: func(env *testscript.Env) error {
				src := "testdata/three_lines.txt"
				dst := filepath.Join(env.WorkDir, "three_lines.txt")

				data, err := os.ReadFile(src)
				if err != nil {
					t.Fatal(err)
				}
				return os.WriteFile(dst, data, 0644)
			},
		*/
	})
}

func TestCountLines(t *testing.T) {
	t.Parallel()
	// in := new(bytes.Buffer)
	// fmt.Fprint(in, "line1\nline2\nline3")
	in := bytes.NewBufferString("line1\nline2\nline3")
	c, err := count.NewCounter(
		count.WithInput(in),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.CountLines()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.CountLines()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestIgnoreEmptyargs(t *testing.T) {
	t.Parallel()
	in := bytes.NewBufferString("line1\nline2\nline3")
	args := []string{}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
		count.WithInput(in),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.CountLines()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestCountWords(t *testing.T) {
	t.Parallel()
	buf := bytes.NewBufferString("two words\nthird_word\nfourth and_fifth_word")
	counter, err := count.NewCounter(count.WithInput(buf))
	if err != nil {
		t.Fatal(err)
	}
	want := 5
	got := counter.CountWords()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
