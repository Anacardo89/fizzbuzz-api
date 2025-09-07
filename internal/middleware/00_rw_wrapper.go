package middleware

import "net/http"

type RWWrapper struct {
	http.ResponseWriter
	status int
	size   int
}

func newRWWrapper(w http.ResponseWriter) *RWWrapper {
	return &RWWrapper{ResponseWriter: w}
}

func getRWWrapper(w http.ResponseWriter) *RWWrapper {
	if rw, ok := w.(*RWWrapper); ok {
		return rw
	}
	return newRWWrapper(w)
}

func (rw *RWWrapper) Status() int {
	return rw.status
}
func (rw *RWWrapper) Size() int {
	return rw.size
}

// Implements http.ResponseWriter
func (rw *RWWrapper) Header() http.Header {
	return rw.ResponseWriter.Header()
}

func (rw *RWWrapper) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

func (rw *RWWrapper) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
