package main

import (
	"strings"
)

type textParts struct {
	text  string
	style string
}

type ArtBuilder struct {
	parts []textParts
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{}
}

func (b *ArtBuilder) AddText(text string) *ArtBuilder {

	b.parts = append(b.parts, textParts{text, "normal"})
	return b
}

func (b *ArtBuilder) SetStyle(style string) *ArtBuilder {

	if style != "normal" && style != "italic" && style != "bold" && style != "outline" {

		panic("invalid style")
	}

	if len(b.parts) > 0 {
		b.parts[len(b.parts)-1].style = style
	}
	return b
}

func (b *ArtBuilder) Build() string {

	lines := make([]string, 8)

	for _, pat := range b.parts {
		for i, s := range render(pat.text, pat.style) {

			lines[i] += s
		}
	}
	return strings.Join(lines, "\n") + "\n"
}

func render(text, style string) []string {

	lines := make([]string, 8)

	if style == "bold" {
		text += text
	}
	if style == "italic" {

		for i := range lines {

			lines[i] = strings.Repeat(" ", 7-i) + text
		}
		return lines
	}

	if style == "outline" {

		border := "+" + strings.Repeat("-", len(text)) + "+"

		lines[0], lines[7], lines[1] = border, border, "|"+text+"|"

		for i := 2; i < 7; i++ {

			lines[i] = "|" + strings.Repeat(" ", len(text)) + "|"
		}
		return lines
	}

	for i := range lines {
		lines[i] = text
	}
	return lines
}
