package log

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// LineFormatter 格式化
type LineFormatter struct {
	TimestampFormat  string
	CallerPrettyfier func(*runtime.Frame) (function string, file string)
}

// Format 格式化
func (f *LineFormatter) Format(entry *Entry) ([]byte, error) {

	data := make(Fields)
	for k, v := range entry.Data {
		data[k] = v
	}

	b := &bytes.Buffer{}
	b.Reset()

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	levelText := strings.ToUpper(entry.Level.String())
	levelText = levelText[0:4]

	b.WriteString(fmt.Sprint("[", entry.Time.Format(timestampFormat), "] "))
	b.WriteString(fmt.Sprint("[", levelText, "] "))

	if entry.HasCaller() {

		var funcVal, fileVal string
		if f.CallerPrettyfier != nil {
			funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
		} else {
			funcVal = entry.Caller.Function
			fileVal = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		}

		b.WriteString(funcVal)
		b.WriteString(" ")
		b.WriteString(fileVal)
		b.WriteString(" ")
	}

	b.WriteString(strings.TrimSuffix(entry.Message, "\n"))

	for k, v := range data {

		appendKeyValue(b, k, v)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func needsQuoting(text string) bool {

	if len(text) == 0 {
		return true
	}

	if text[0] == '"' {
		return false
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	appendValue(b, value)
}

func appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
