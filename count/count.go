package count

import (
	"bufio"
	"errors"
	"flag"
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

func (c *Counter) CloseFiles() {
	for _, file := range c.files {
		file.(io.Closer).Close()
	}
}

func (c *Counter) CountBytes() int {
	bytes := 0
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes++
	}
	c.CloseFiles()
	return bytes
}

func (c *Counter) CountLines() int {
	lines := 0
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		lines++
	}
	c.CloseFiles()
	return lines
}

func (c *Counter) CountWords() int {
	words := 0
	scanner := bufio.NewScanner(c.reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	c.CloseFiles()
	return words
}

func Main() int {
	var c *Counter
	var err error
	lineMode := flag.Bool("lines", false, "Count lines instead of words")
	byteMode := flag.Bool("bytes", false, "Count bytes instead of words")
	flag.Parse()
	if len(flag.Args()) > 1 {
		c, err = NewCounter(WithInputFromArgs(flag.Args()))
	} else {
		c, err = NewCounter(WithInput(os.Stdin))
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	switch {
	case *lineMode && *byteMode:
		fmt.Fprintln(os.Stderr, "-bytes and -lines are exclusive flags")
		return 1
	case *lineMode:
		fmt.Println(c.CountLines())
	case *byteMode:
		fmt.Println(c.CountBytes())
	default:
		fmt.Println(c.CountWords())
	}
	return 0
}
