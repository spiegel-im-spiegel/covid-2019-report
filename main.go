package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/cov19data"
	"github.com/spiegel-im-spiegel/cov19data/client"
	"github.com/spiegel-im-spiegel/cov19data/entity"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/covid-2019-report/chart"
	"github.com/spiegel-im-spiegel/covid-2019-report/imgutil"
)

const (
	histcasesFile  = "./covid-2019-new-cases-histgram-in-japan.png"
	histdeathsFile = "./covid-2019-new-deaths-histgram-in-japan.png"
	histallFile    = "./covid-2019-histgram-in-japan.png"
)

func main() {
	data, err := cov19data.ImportWHOCSV(
		client.Default(),
		entity.WithFilterPeriod(values.NewPeriod(values.NewDate(2020, time.Month(3), 11), values.NewDateTime(time.Time{}))),
		entity.WithCountryCode(values.CC_JP),
		entity.WithRegionCode(values.WPRO),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	h, err := cov19data.MakeHistgramWHO(data, 7)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	histChart := chart.ImportHistgramData(h)

	if err := chart.BarChartHistCases(histChart, histcasesFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := chart.BarChartHistDeaths(histChart, histdeathsFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := imgutil.ConcatImageFiles(histallFile, histcasesFile, histdeathsFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
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
