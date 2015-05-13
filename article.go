package blog

import "time"

//Article ...
type Article struct {
	Title  string
	Date   CustomTime
	Update CustomTime
	Tags   []string
}

//CustomTime ...
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

//UnmarshalJSON ...
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.Parse(ctLayout, string(b))
	return
}

//MarshalJSON ...
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

var nilTime = (time.Time{}).UnixNano()

// IsSet ...
func (ct *CustomTime) IsSet() bool {
	return ct.UnixNano() != nilTime
}
