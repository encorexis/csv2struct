package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CSV struct {
	Name    string
	Headers []*Header
	Records []Record
	Buf     [][]string
}

func New(file *os.File) *CSV {
	c := &CSV{
		Name: TrimExt(file.Name()),
	}
	c.MakeBuf(file)
	c.MakeHeaders()
	c.MakeRecords()
	return c
}

func (c *CSV) MakeBuf(reader io.Reader) {
	r := csv.NewReader(reader)
	buf, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	c.Buf = buf
}

func (c *CSV) MakeHeaders() {
	if len(c.Buf) == 0 {
		log.Fatal("no data")
	}
	for _, name := range c.Buf[0] {
		header := &Header{
			Name: name,
		}
		c.Headers = append(c.Headers, header)
	}
}

func (c *CSV) MakeRecords() {
	if len(c.Buf) < 2 {
		log.Fatal("no records")
	}
	for recordIndex := 1; recordIndex < len(c.Buf); recordIndex++ {
		record := Record{
			Data: make(map[*Header]string),
		}
		for headerIndex := range c.Headers {
			record.Data[c.Headers[headerIndex]] = c.Buf[recordIndex][headerIndex]
		}
		c.Records = append(c.Records, record)
	}
}

func TrimExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
