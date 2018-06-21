package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

const (
	BUFSIZE = math.MaxInt32
)

type Buffer struct {
	data []string
	name string
}

func initBuffer(path string) (*Buffer, error) {
	buf := new(Buffer)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory", path)
	}

	if err := buf.open(path); err != nil {
		return nil, err
	}
	return buf, nil
}

func (buf *Buffer) open(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buf.name = f.Name()

	if err := buf.read(f); err != nil {
		return err
	}

	return nil
}

func (buf *Buffer) read(f *os.File) error {
	scanner := bufio.NewScanner(f)
	var str []string
	for scanner.Scan() {
		text := ""
		for _, v := range []byte(scanner.Text()) {
			if v == byte('\t') {
				text += "  "
				continue
			}
			text += string(v)
		}
		str = append(str, text)
	}
	if err := scanner.Err(); err != nil {
		return errors.New(fmt.Sprintf("scanner err:", err))
	}
	buf.data = str

	return nil
}
