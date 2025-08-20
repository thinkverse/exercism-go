package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	reader io.Reader
	nbyte  int64
	nops   int
	mutex  sync.RWMutex
}

type writeCounter struct {
	writer io.Writer
	nbyte  int64
	nops   int
	mutex  sync.RWMutex
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readwriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{nbyte: 0, nops: 0, writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{nbyte: 0, nops: 0, reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readwriteCounter{
		NewReadCounter(readwriter),
		NewWriteCounter(readwriter),
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	n, err := rc.reader.Read(p)

	if err == nil {
		rc.nbyte += int64(n)
		rc.nops++
	}

	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mutex.RLock()
	defer rc.mutex.RUnlock()

	return rc.nbyte, rc.nops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	n, err := wc.writer.Write(p)

	if err == nil {
		wc.nbyte += int64(n)
		wc.nops++
	}

	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mutex.RLock()
	defer wc.mutex.RUnlock()

	return wc.nbyte, wc.nops
}
