package interfaces

import (
	"fmt"
	"strings"
)

type StringSeparator interface {
	Separate(value string) string
}

type CommaSeparator struct {
}

func (c *CommaSeparator) Separate(value string) string {
	return strings.Join(strings.Split(value, " "), ",")
}

func (c *CommaSeparator) Divider() {
}

type DotSeparator struct {
}

func (d *DotSeparator) Separate(value string) string {
	return strings.Join(strings.Split(value, " "), ".")
}

type PipeSeparator struct {
	value string
}

func (p *PipeSeparator) Separate(value string) error {
	splitValue := strings.Split(value, " ")
	if len(splitValue) < 2 {
		return fmt.Errorf("value must have at least two chars separated by a blank space")
	}
	p.value = strings.Join(splitValue, "|")
	return nil
}

func HandleString(value string, separator StringSeparator) {
	fmt.Printf("before: %q, after: %q\n", value, separator.Separate(value))
}
