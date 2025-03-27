package match

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type match struct {
	reader io.Reader
	writer io.Writer
}

func NewMatcher() *match {
	return &match{
		reader: os.Stdin,
		writer: os.Stdout,
	}
}

func (m *match) WithReader(r io.Reader) *match {
	m.reader = r
	return m
}

func (m *match) WithWriter(w io.Writer) *match {
	m.writer = w
	return m
}

func (m *match) Match(s string) {
	scanner := bufio.NewScanner(m.reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {
			fmt.Fprintln(m.writer, scanner.Text())
		}
	}
}
