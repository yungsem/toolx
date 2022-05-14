package strs

import "strings"

type builder struct {
	sb strings.Builder
}

func Builder() *builder {
	var sb strings.Builder
	return &builder{sb}
}

func (b *builder) Append(s string) *builder {
	b.sb.WriteString(s)
	return b
}

func (b *builder) String() string {
	return b.sb.String()
}
