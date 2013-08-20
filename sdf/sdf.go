package sdf

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

type sdfParser struct {
	entry    *Entry
	curModel int
	line     []byte
	modified map[string]string
	seqres   map[byte][]string

	missing map[byte][]missingResidue
	r465    bool

	processed map[seen]bool
	lastSeen  seen
}

func ReadSDF(fp string) (*Entry, error) {
	var reader io.Reader
	var err error

	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader = f

	// If the file is gzipped, use the gzip decompressor.
	if path.Ext(fp) == ".gz" {
		reader, err = gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
	}

	return Read(reader, fp)
}
