package datatype

import "net/http"

// StubResponseWriter could be used for testing purposes
type StubResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

// NewStubResponseWriter creates new StubResponseWriter
func NewStubResponseWriter() *StubResponseWriter {
	return &StubResponseWriter{
		header: http.Header{},
	}
}

// Header returns writer header
func (rw *StubResponseWriter) Header() http.Header {
	return rw.header
}

// Write puts buffer to writer body. Returns
func (rw *StubResponseWriter) Write(buf []byte) (int, error) {
	rw.body = buf
	// implement it as per your requirement
	return 0, nil
}

// WriteHeader puts statusCode to writer
func (rw *StubResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}
