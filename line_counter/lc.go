package lcounter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type Counter struct {
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
	return lines
}

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Count())
}
