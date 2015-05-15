package blog

import (
	"fmt"
	"os"
	"time"
)

//Article ...
type Article struct {
	Title         string    `json:"title"`
	Date          *JSONTime `json:"date"`
	Tags          []string  `json:"tags"`
	*JSONFileInfo `json:"fileInfo"`
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
	if err != nil {
		t.Time, err = time.Parse(time.RFC3339, string(b))
	}
	return
}

//JSONFileInfo ...
type JSONFileInfo struct {
	FName    string      `json:"name"`
	FSize    int64       `json:"size"`
	FMode    os.FileMode `json:"mode"`
	FModTime *JSONTime   `json:"modTime"`
	FIsDir   bool        `json:"isDir"`
}

//Name ...
func (f JSONFileInfo) Name() string {
	return f.FName
}

//Size ...
func (f JSONFileInfo) Size() int64 {
	return f.FSize
}

//Mode ...
func (f JSONFileInfo) Mode() os.FileMode {
	return f.FMode
}

//ModTime ...
func (f JSONFileInfo) ModTime() *JSONTime {
	return f.FModTime
}

//IsDir ...
func (f JSONFileInfo) IsDir() bool {
	return f.FIsDir
}
