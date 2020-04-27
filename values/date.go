package values

import (
	"strconv"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

//Time is wrapper class of time.Time
type Date struct {
	time.Time
}

//NewDate returns Time instance
func NewDate(tm time.Time) Date {
	return Date{tm}
}

var timeTemplate = []string{
	"2006-01-02",
	time.RFC3339,
}

func DateFrom(s string) (Date, error) {
	if len(s) == 0 || strings.EqualFold(s, "null") {
		return NewDate(time.Time{}), nil
	}
	var lastErr error
	for _, tmplt := range timeTemplate {
		if tm, err := time.Parse(tmplt, s); err != nil {
			lastErr = errs.Wrap(err, "", errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
		} else {
			return Date{tm}, nil
		}
	}
	return NewDate(time.Time{}), lastErr
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
func (t *Date) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	tm, err := DateFrom(s)
	if err != nil {
		return err
	}
	*t = tm
	return nil
}

//MarshalJSON returns time string with RFC3339 format
func (t *Date) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("\"\""), nil
	}
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

func (t Date) String() string {
	return t.Format("2006-01-02")
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
