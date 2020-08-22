package report

import (
	"encoding/csv"
	"io"

	"github.com/spiegel-im-spiegel/covid-2019-report/cases"
	"github.com/spiegel-im-spiegel/covid-2019-report/ecode"
	"github.com/spiegel-im-spiegel/errs"
)

func importWHOCSV(r io.Reader) ([]cases.Cases, error) {
	cr := csv.NewReader(r)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading

	cs := make([]cases.Cases, 0, 128)
	header := false
	for {
		elms, err := cr.Read()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if len(elms) < 8 {
			return nil, errs.Wrap(ecode.ErrInvalidRecord, errs.WithContext("record", elms))
		}
		if !header {
			c, err := cases.New(elms[0], elms[5], elms[7])
			if err != nil {
				return nil, errs.Wrap(err)
			}
			cs = append(cs, c)
		} else {
			header = false
		}
	}
	return cs, nil
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
