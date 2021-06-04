package rake

import (
	"encoding/json"
	"io"
)

func DebugWriter(writer io.Writer) *debugWriter {
	return &debugWriter{
		writer: writer,
	}
}

type debugWriter struct {
	writer io.Writer
}

func (s *debugWriter) Load(configPtr interface{}) {
	io.WriteString(s.writer, "Config : ")
	objectString, err := json.MarshalIndent(configPtr, "", "  ")
	check(err)
	_, err = s.writer.Write([]byte(objectString))
	check(err)
	io.WriteString(s.writer, "\n")
}
