package tools

import (
	"fmt"
)

// MarkDown 是为了在代码中将string方便的转为MD格式的string
type MarkDown struct {
	Content string
}

const (
	newLine = "\r\n"
)

// NewLine 换行新起一行
func (m *MarkDown) NewLine() *MarkDown {
	m.Content += newLine
	return m
}

// NewLines 新起多行
func (m *MarkDown) NewLines(count int) *MarkDown {
	if count <= 0 {
		return m
	}
	mk := ""
	for i := 0; i < count; i++ {
		mk += newLine
	}
	m.Content += mk
	return m
}

// Text 普通文本
func (m *MarkDown) Text(str string) *MarkDown {
	m.Content += str
	return m
}

// Title 标题，level 表示了#的数量
func (m *MarkDown) Title(str string, level int) *MarkDown {
	mk := ""
	for i := 0; i < level; i++ {
		mk += "#"
	}
	m.Content += mk + " " + str
	return m
}

// Blockquotes 引用
func (m *MarkDown) Blockquotes(str string) *MarkDown {
	m.Content += "> " + str
	return m
}

// List 普通的列表
func (m *MarkDown) List(str string) *MarkDown {
	m.Content += "-   " + str
	return m
}

// ListOrder 前缀为1. 2. 3. 的列表
func (m *MarkDown) ListOrder(strs []string) *MarkDown {
	var t string
	for i := 0; i < len(strs); i++ {
		t += string(i+1) + "   " + strs[i] + newLine
	}
	m.Content += t
	return m
}

// CodeOneLine 单行的代码文本
func (m *MarkDown) CodeOneLine(str string) *MarkDown {
	m.Content += "`" + str + "`"
	return m
}

// Code 多行代码文本 codelang对应代码语言
func (m *MarkDown) Code(str string, codelang string) *MarkDown {
	m.Content += newLine + "```" + codelang + newLine + str + "```" + newLine
	return m
}

// DividingLine 下划线
func (m *MarkDown) DividingLine() *MarkDown {
	m.Content += newLine + "---------" + newLine
	return m
}

// Link A标签 连接
func (m *MarkDown) Link(name string, url string) *MarkDown {
	m.Content += "[" + name + "](" + url + ")"
	return m
}

// Img 图片连接
func (m *MarkDown) Img(name string, url string) *MarkDown {
	m.Content += "![" + name + "](" + url + ")"
	return m
}

// Bold 加粗
func (m *MarkDown) Bold(str string) *MarkDown {
	m.Content += "**" + str + "**"
	return m
}

// Table 简易的表格，将kv转化为表格
func (m *MarkDown) Table(kv map[string]interface{}) *MarkDown {
	m.Content += newLine + "| key | value |" + newLine + "| :------------ | :------------ |"

	for key, value := range kv {
		m.Content += newLine + fmt.Sprintf("| %s | %s |", key, value)
	}
	return m
}
