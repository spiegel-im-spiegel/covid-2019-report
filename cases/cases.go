package cases

import (
	"strconv"

	"github.com/spiegel-im-spiegel/covid-2019-report/values"
	"github.com/spiegel-im-spiegel/errs"
)

type Cases struct {
	Date   values.Date
	Total  int64 `json:"TotalConfirmedCases"`
	Deaths int64 `json:"TotalDeaths"`
}

func New(date, total, deaths string) (Cases, error) {
	c := Cases{}
	dt, err := values.DateFrom(date)
	if err != nil {
		return c, errs.Wrap(err, "Invalid data", errs.WithContext("date", date))
	}
	c.Date = dt

	t, err := strconv.ParseInt(total, 10, 64)
	if err != nil {
		return c, errs.Wrap(err, "Invalid data", errs.WithContext("total", total))
	}
	c.Total = t

	d, err := strconv.ParseInt(deaths, 10, 64)
	if err != nil {
		return c, errs.Wrap(err, "Invalid data", errs.WithContext("deaths", deaths))
	}
	c.Deaths = d

	return c, err
}

func (c Cases) FatalityRate() float64 {
	if c.Total == 0 {
		return 0.0
	}
	return (float64)(c.Deaths) / (float64)(c.Total)
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
