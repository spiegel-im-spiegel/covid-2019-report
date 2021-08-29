package excel

import (
	"strings"

	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/xuri/excelize/v2"
)

type Infection struct {
	NodeMatsue  string
	NodeShimane string
	Date        date.Date
	InsideFlag  bool
	FromInside  []string
	FromOutside []string
}

func NewInfections(xlsx *excelize.File, sheetIndex int) ([]*Infection, error) {
	rows, err := xlsx.Rows(xlsx.GetSheetName(sheetIndex))
	if err != nil {
		var errSheet excelize.ErrSheetNotExist
		if errs.As(err, &errSheet) {
			return nil, errs.Wrap(ErrInvalidSheetName, errs.WithCause(err))
		}
		return nil, errs.Wrap(err)
	}

	infections := []*Infection{}
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			return nil, errs.Wrap(err)
		}
		if len(cols) < 3 {
			return nil, errs.Wrap(ErrInvalidExcelData, errs.WithContext("cols", cols))
		}
		i := &Infection{
			NodeMatsue:  "m" + cols[0],
			NodeShimane: "s" + cols[1],
			Date:        date.FromString(cols[2]),
			InsideFlag:  false,
			FromInside:  []string{},
			FromOutside: []string{},
		}
		if len(cols) >= 4 && strings.EqualFold(cols[3], "1") {
			i.InsideFlag = true
		}
		if len(cols) >= 5 && len(cols[4]) > 0 {
			i.FromInside = append(i.FromInside, "m"+cols[4])
		}
		if len(cols) >= 6 && len(cols[5]) > 0 {
			i.FromInside = append(i.FromInside, "m"+cols[5])
		}
		if len(cols) >= 7 && len(cols[6]) > 0 {
			i.FromInside = append(i.FromInside, "m"+cols[6])
		}
		if len(cols) >= 8 && len(cols[7]) > 0 {
			i.FromInside = append(i.FromInside, "m"+cols[7])
		}
		if len(cols) >= 9 && len(cols[8]) > 0 {
			i.FromOutside = append(i.FromOutside, cols[8])
		}
		if len(cols) >= 10 && len(cols[9]) > 0 {
			i.FromOutside = append(i.FromOutside, cols[9])
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
