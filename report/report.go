package report

import (
	"sort"

	"github.com/spiegel-im-spiegel/covid-2019-report/cases"
	"github.com/spiegel-im-spiegel/covid-2019-report/ecode"
	"github.com/spiegel-im-spiegel/covid-2019-report/values"
	"github.com/spiegel-im-spiegel/errs"
)

type Report interface {
	Date() values.Date
	TotalCases() int64
	TotalDeaths() int64
	CasesByDate() int64
	DeathsByDate() int64
	FatalityRate() float64
}

type report struct {
	today     cases.Cases
	yesterday cases.Cases
}

func newReport(today, yesterday cases.Cases) *report {
	return &report{today: today, yesterday: yesterday}
}

func (r *report) Date() values.Date {
	return r.today.Date
}

func (r *report) TotalCases() int64 {
	return r.today.Total
}

func (r *report) TotalDeaths() int64 {
	return r.today.Deaths
}

func (r *report) CasesByDate() int64 {
	return r.today.Total - r.yesterday.Total
}

func (r *report) DeathsByDate() int64 {
	return r.today.Deaths - r.yesterday.Deaths
}

func (r *report) FatalityRate() float64 {
	return r.today.FatalityRate()
}

type Reports interface {
	Len() int
	Index(i int) (Report, error)
	Next() (Report, error)
	Top() (Report, error)
	Last() (Report, error)
	SearchByDate(dt values.Date) (Report, error)
}

type reports struct {
	index int
	data  []cases.Cases
}

func New(cs []cases.Cases) Reports {
	if len(cs) == 0 {
		return &reports{index: -1, data: []cases.Cases{}}
	}
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Date.Before(cs[j].Date.Time)
	})
	return &reports{index: -1, data: cs}
}

func (rs *reports) Len() int {
	if rs == nil {
		return 0
	}
	return len(rs.data)
}

func (rs *reports) Index(i int) (Report, error) {
	if rs == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "")
	}
	if i < 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("index", i))
	}
	rs.index = i
	return rs.get()
}

func (rs *reports) Next() (Report, error) {
	if rs == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "")
	}
	rs.index++
	return rs.get()
}

func (rs *reports) Top() (Report, error) {
	if rs == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "")
	}
	rs.index = 0
	return rs.get()
}

func (rs *reports) Last() (Report, error) {
	if rs == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "")
	}
	if rs.Len() == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("reports.Len", rs.Len()))
	}
	rs.index = rs.Len() - 1
	return rs.get()
}

func (rs *reports) SearchByDate(dt values.Date) (Report, error) {
	if rs == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer, "")
	}
	if rs.Len() == 0 {
		return nil, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("reports.Len", rs.Len()))
	}
	for i, c := range rs.data {
		if c.Date.Year() == dt.Year() && c.Date.Month() == dt.Month() && c.Date.Day() == dt.Day() {
			rs.index = i
			return rs.get()
		}
	}
	return nil, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("Date", dt.String()))
}

func (rs *reports) get() (*report, error) {
	if rs.index < 0 || rs.index >= rs.Len() {
		return nil, errs.Wrap(ecode.ErrNoData, "", errs.WithContext("reports.index", rs.index))
	}
	if rs.index == 0 {
		return newReport(rs.data[0], cases.Cases{}), nil
	}
	return newReport(rs.data[rs.index], rs.data[rs.index-1]), nil
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
