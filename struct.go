package dfl

import (
	"fmt"
	"strings"
)

type Graph struct {
	records  []Record
	arrows   []Arrow
	template Template
}

func (g Graph) String() string {
	builder := strings.Builder{}
	for _, r := range g.records {
		builder.WriteString(fmt.Sprintf("%s\n", r.String()))
	}
	arrow := strings.Builder{}
	for _, r := range g.arrows {
		arrow.WriteString(fmt.Sprintf("%s\n", r.String()))
	}
	return strings.Join([]string{
		g.template.head(),
		builder.String(),
		arrow.String(),
		g.template.tail(),
	}, "\n",
	)
}

type Record struct {
	name   string
	values []Value
}

func (r Record) String() string {
	strs := make([]string, 0)
	for _, v := range r.values {
		strs = append(strs, v.String())
	}

	values := strings.Join(strs, "|")
	return fmt.Sprintf(`%s [label="%s|%s"]`,r.name, r.name, values)
}

type Value struct {
	name string
	port string
}

func (v Value) String() string {
	return "<" + v.port + "> " + v.name
}

type Arrow struct {
	src  string
	dest string
}

func (v Arrow) String() string {
	return fmt.Sprintf("%s -> %s", v.src, v.dest)
}

type Template interface {
	head() string
	tail() string
}

type SimpleTempate struct {
}

func (SimpleTempate) head() string {
	return `digraph structs {
 node [shape=record]
 rankdir=TB
`

}

func (SimpleTempate) tail() string {
	return ` }`
}
