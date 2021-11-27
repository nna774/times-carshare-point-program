package http

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nna774/times-carshare-point-program/tcp"
)

const (
	lib = iota
	myImpl
)

var mode = myImpl

func Get(uri string) (*http.Response, error) {
	if mode == lib {
		return http.Get(uri)
	}
	u, err := url.Parse(uri) // parse mendoi
	if err != nil {
		return nil, err
	}
	conn, err := tcp.NewTCPConnection(u.Host, 80) // kimeuchi
	if err != nil {
		return nil, fmt.Errorf("tcp conn creation failed: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			panic(err) // believe never
		}
	}()
	conn.Write([]byte(fmt.Sprintf("GET %v HTTP/1.1\r\nHost: %v\r\n\r\n", u.Path, u.Host)))
	reader := bufio.NewReader(conn)
	sstatus, p, err := reader.ReadLine()
	status := string(sstatus)
	if err != nil {
		return nil, fmt.Errorf("ReadLine failed while parse status line(%v)", err)
	}
	if p {
		return nil, fmt.Errorf("status line too long(prefix: %v)", status)
	}
	ss := strings.Split(status, " ")
	if len(ss) < 2 {
		return nil, fmt.Errorf("got bad response(%v)", status)
	}
	fmt.Printf("###http### status line: %v\n", status)
	code, err := strconv.Atoi(ss[1])
	if err != nil {
		return nil, fmt.Errorf("bad status code(%v): %v", code, err)
	}
	headers := http.Header{} // Header.get shitai
	for {
		bline, _, err := reader.ReadLine()
		line := string(bline)
		if err != nil {
			return nil, fmt.Errorf("ReadLine failed while parse header(%v)", err)
		}
		fmt.Printf("###http### header line: %v\n", line)
		if line == "" {
			break
		}

		ls := strings.SplitN(line, ":", 2)
		headers[ls[0]] = append(headers[ls[1]], ls[1][1:])
	}
	slength := headers.Get("Content-Length") // koko
	length, err := strconv.Atoi(slength)
	fmt.Printf("###http### content-length %v\n", length)
	buf := &bytes.Buffer{}
	if err == nil { // Content-Length atta; read body
		read, _ := io.CopyN(buf, reader, int64(length))
		fmt.Printf("###http### content read %v\n", read)
	}
	return &http.Response{
		Header:     headers,
		StatusCode: code,
		Body:       io.NopCloser(buf),
	}, nil
}
