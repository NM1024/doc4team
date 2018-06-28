package tools

import (
	"fmt"
)

type MarkDown struct {
	Content string
}

const (
	newLine = "\r\n"
)

func (m *MarkDown) NewLine() *MarkDown {
	m.Content += newLine
	return m
}
func (m *MarkDown) NewLines(count int) *MarkDown {
	mk := ""
	for i := 0; i < count; i++ {
		mk += newLine
	}
	m.Content += mk
	return m
}

func (m *MarkDown) Text(str string) *MarkDown {
	m.Content += str
	return m
}
func (m *MarkDown) Title(str string, level int) *MarkDown {
	mk := ""
	for i := 0; i < level; i++ {
		mk += "#"
	}
	m.Content += mk + " " + str
	return m
}

func (m *MarkDown) Blockquotes(str string) *MarkDown {
	m.Content += "> " + str
	return m
}

func (m *MarkDown) List(str string) *MarkDown {
	m.Content += "-   " + str
	return m
}

func (m *MarkDown) ListOrder(strs []string) *MarkDown {
	var t string
	for i := 0; i < len(strs); i++ {
		t += string(i+1) + "   " + strs[i] + newLine
	}
	m.Content += t
	return m
}
func (m *MarkDown) CodeOneLine(str string) *MarkDown {
	m.Content += "`" + str + "`"
	return m
}
func (m *MarkDown) Code(str string, codelang string) *MarkDown {
	m.Content += newLine + "```" + codelang + newLine + str + "```" + newLine
	return m
}
func (m *MarkDown) DividingLine() *MarkDown {
	m.Content += newLine + "*****" + newLine
	return m
}

func (m *MarkDown) Link(name string, url string) *MarkDown {
	m.Content += "[" + name + "](" + url + ")"
	return m
}

func (m *MarkDown) Img(name string, url string) *MarkDown {
	m.Content += "![" + name + "](" + url + ")"
	return m
}

func (m *MarkDown) Bold(str string) *MarkDown {
	m.Content += "**" + str + "**"
	return m
}

func (m *MarkDown) Table(kv map[string]interface{}) *MarkDown {
	m.Content += newLine + "| key | value |" + newLine + "| :------------ | :------------ |"

	for key, value := range kv {
		m.Content += newLine + fmt.Sprintf("| %s | %s |", key, value)
	}
	return m
}
