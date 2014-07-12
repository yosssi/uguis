package uguis

// File chnage types
const (
	fileChangeTypeCreate = iota
	fileChangeTypeDelete
)

const filenameFormat = "uguis-%d"

// file represents a file.
type file struct {
	path       string
	data       []byte
	changeType uint8
}

// newFile creates and returns a file.
func newFile(path string, data []byte, changeType uint8) file {
	return file{
		path:       path,
		data:       data,
		changeType: changeType,
	}
}
