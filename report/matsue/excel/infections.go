package excel

import (
	"io"
	"strings"

	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/errs"

	"github.com/spiegel-im-spiegel/csvdata"
	"github.com/spiegel-im-spiegel/csvdata/exceldata"
)

type Infection struct {
	NodeMatsue  string
	NodeShimane string
	Date        date.Date
	InsideFlag  bool
	FromInside  []string
	FromOutside []string
}

func NewInfections(path, sheetName string) ([]*Infection, error) {
	ods, err := exceldata.OpenFile(path, "")
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("path", path))
	}
	r, err := exceldata.New(ods, sheetName)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("path", path), errs.WithContext("sheetName", sheetName))
	}
	rows := csvdata.NewRows(r, true)
	rows.Close() // dummy

	infections := []*Infection{}
	for {
		if err := rows.Next(); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return infections, errs.Wrap(err)
		}
		i := &Infection{
			NodeMatsue:  "m" + rows.Get(0),
			NodeShimane: "s" + rows.Get(1),
			Date:        date.FromString(rows.Get(2)),
			InsideFlag:  false,
			FromInside:  []string{},
			FromOutside: []string{},
		}
		if strings.EqualFold(rows.Get(3), "1") {
			i.InsideFlag = true
		}
		for n := 4; n <= 7; n++ {
			if s := rows.Get(n); len(s) > 0 {
				i.FromInside = append(i.FromInside, "m"+s)
			}
		}
		for n := 8; n <= 9; n++ {
			if s := rows.Get(n); len(s) > 0 {
				i.FromOutside = append(i.FromOutside, s)
			}
		}
		infections = append(infections, i)
	}
	return infections, nil
}

/* Copyright 2021 Spiegel
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
