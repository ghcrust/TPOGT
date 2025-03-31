package lcounter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	files  []io.Reader
	reader io.Reader
	writer io.Writer
}

type option func(*Counter) error

// functional option constructors
func WithInput(r io.Reader) option {
	return func(c *Counter) error {
		if r == nil {
			return errors.New("nil reader")
		}
		c.reader = r
		return nil
	}
}

func WithOutput(w io.Writer) option {
	return func(c *Counter) error {
		if w == nil {
			return errors.New("nil writer")
		}
		c.writer = w
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		if len(args) < 1 {
			return nil
		}
		c.files = make([]io.Reader, len(args))
		for i, f := range args {
			file, err := os.Open(f)
			//defer file.Close()
			if err != nil {
				return err
			}
			c.files[i] = file
			/*
				r, _ := io.ReadAll(file)
				reader := bytes.NewReader(r)
				/ reader := new(bytes.Buffer)
				// reader.WriteFrom(file)
				c.reader = reader
			*/
		}
		c.reader = io.MultiReader(c.files...)
		return nil
	}
}

/* alternative: methodical options
func (c *Counter) WithInput(r io.Reader) *Counter {
	c.reader = r
	return c
}
func (c *Counter) WithOutput(w io.Writer) *Counter {
	c.writer = w
	return c
}
lc.NC().WI(r).WO(w)
*/

func NewCounter(opts ...option) (*Counter, error) {
	c := new(Counter)
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Counter) Count() int {
	lines := 0
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		lines++
	}
	for _, file := range c.files {
		file.(io.Closer).Close()
	}
	return lines
}

func Main() {
	/*
		var c *Counter
		var err error
		if len(os.Args) > 1 {
			c, err = NewCounter(WithInputFromArgs(os.Args[1:]))
		} else {
			c, err = NewCounter(WithInput(os.Stdin))
		}
	*/
	c, err := NewCounter(WithInput(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(c.Count())
	}
}
