package request

import (
	"encoding/json"
	"time"
)

const DateLayout = "2006-01-02 15:04:05 -0700"

type Timestamp struct {
	t time.Time
}

func (st *Timestamp) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return err
	}
	st.t = t
	return nil
}

func (st *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.t.Format(DateLayout))
}

func (st *Timestamp) ToTime() time.Time {
	return st.t
}

func (st *Timestamp) String() string {
	return st.t.Format(DateLayout)
}
