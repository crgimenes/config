package helper

import (
	"io"

	"github.com/nuveo/log"
)

// Closer close descriptor to use with defer.
func Closer(f io.Closer) {
	err := f.Close()
	if err != nil {
		log.Errorln(err)
	}
}
