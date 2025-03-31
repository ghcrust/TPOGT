package lcounter_test

import (
	"bytes"
	"lcounter"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"lcounter": lcounter.Main,
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

func TestCounter(t *testing.T) {
	t.Parallel()
	// in := new(bytes.Buffer)
	// fmt.Fprint(in, "line1\nline2\nline3")
	in := bytes.NewBufferString("line1\nline2\nline3")
	c, err := lcounter.NewCounter(
		lcounter.WithInput(in),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestWithInputFromArgs_SetToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := lcounter.NewCounter(
		lcounter.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	in := bytes.NewBufferString("line1\nline2\nline3")
	args := []string{}
	c, err := lcounter.NewCounter(
		lcounter.WithInputFromArgs(args),
		lcounter.WithInput(in),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Count()
	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
