package uguis

import "os"

const serviceNameSimpleFileWriter = "simpleFileWriter"

// simpleFileWriter represents a simple file writer.
type simpleFileWriter struct {
	reqC       chan FileWriterRequest
	resC       chan FileWriterResponse
	closedReqC chan struct{}
	app        *Application
	lgr        Logger
}

// Write writes a file.
func (w *simpleFileWriter) Write(req FileWriterRequest) {
	// Send a request to the write goroutine.
	w.reqC <- req
}

// Close closes the file writer.
func (w *simpleFileWriter) Close() error {
	// Close the request channel.
	close(w.reqC)

	// Wait until the write goroutine is closed.
	<-w.closedReqC

	return nil
}

// ResC returns a response channel.
func (w *simpleFileWriter) ResC() <-chan FileWriterResponse {
	return w.resC
}

// write writes a file.
func (w *simpleFileWriter) write() {
	for req := range w.reqC {
		switch req.file.changeType {
		case fileChangeTypeCreate:
			path := req.file.path

			f, err := os.Create(path)
			if err != nil {
				w.logError(err)
				continue
			}

			_, err = f.Write(req.file.data)
			f.Close()
			if err != nil {
				w.logError(err)
				continue
			}

			w.resC <- NewFileWriterResponse(req.tweet, path)
		case fileChangeTypeDelete:
			if err := os.Remove(req.file.path); err != nil {
				w.logError(err)
				continue
			}
		}
	}

	// Send a closed signal.
	w.closedReqC <- struct{}{}
}

func (w *simpleFileWriter) logError(err error) {
	w.lgr.Print(NewLog(
		LogLevelERROR,
		w.app.Hostname,
		serviceNameSimpleFileWriter,
		err.Error(),
	))
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
		reqC:       make(chan FileWriterRequest, opts.ReqCBfSize),
		resC:       make(chan FileWriterResponse, opts.ResCBfSize),
		closedReqC: make(chan struct{}),
		app:        app,
		lgr:        lgr,
	}

	// Launch a write goroutine.
	go w.write()

	return w
}
