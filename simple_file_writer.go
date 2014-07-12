package uguis

import "fmt"

const serviceNameSimpleFileWriter = "simpleFileWriter"

// simpleFileWriter represents a simple file writer.
type simpleFileWriter struct {
	reqC       chan file
	resC       chan string
	closedReqC chan struct{}
	app        *Application
	lgr        Logger
}

// Write writes a file.
func (w *simpleFileWriter) Write(f file) {
	// Send a request to the write goroutine.
	w.reqC <- f
}

// Close closes the file writer.
func (w *simpleFileWriter) Close() error {
	// Close the request channel.
	close(w.reqC)

	// Wait until the write goroutine is closed.
	<-w.closedReqC

	return nil
}

// write writes a file.
func (w *simpleFileWriter) write() {
	for f := range w.reqC {
		//TODO
		fmt.Println(f)
	}

	// Send a closed signal.
	w.closedReqC <- struct{}{}
}

// NewSimpleFileWriter creates and returns a simple file writer.
func NewSimpleFileWriter(
	app *Application,
	lgr Logger,
	opts *SimpleFileWriterOptions,
) FileWriter {
	// Initialize options.
	if opts == nil {
		opts = &SimpleFileWriterOptions{}
	}
	opts.setDefaults()

	w := &simpleFileWriter{
		reqC:       make(chan file, opts.ReqCBfSize),
		resC:       make(chan string, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a write goroutine.
	go w.write()

	return w
}
