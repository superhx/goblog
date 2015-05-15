package blog

import (
	"fmt"
	"os"
	"time"
)

//Article ...
type Article struct {
	Title  string
	Date   *JSONTime
	Update *JSONTime
	Tags   []string
	Origin *JSONFileInfo
}

//JSONTime ...
type JSONTime struct {
	time.Time
}

//MarshalJSON ...
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))
	return []byte(stamp), nil
}

//UnmarshalJSON ...
func (t *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	t.Time, err = time.Parse("2006/01/02|15:04:05", string(b))
	return
}

//JSONFileInfo ...
type JSONFileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime *JSONTime
	IsDir   bool
}
