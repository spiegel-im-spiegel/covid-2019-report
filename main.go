package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/covid-2019-report/chart"
	"github.com/spiegel-im-spiegel/covid-2019-report/imgutil"
	"github.com/spiegel-im-spiegel/errs"
)

const (
	HistCasesFile       = "./covid-2019-new-cases-histgram-in-japan.png"
	HistDeathsFile      = "./covid-2019-new-deaths-histgram-in-japan.png"
	HistCasesFileShort  = "./covid-2019-new-cases-histgram-in-japan-short.png"
	HistDeathsFileShort = "./covid-2019-new-deaths-histgram-in-japan-short.png"
	HistAllFile         = "./covid-2019-histgram-in-japan.png"
)

func makeGraph(from, to values.Date, histCasesFile, histDeathsFile string) error {
	p := values.NewPeriod(from, to)
	global, err := getGlobalHist(p)
	if err != nil {
		return errs.Wrap(err)
	}
	tokyo, err := getTokyoHist(p)
	if err != nil {
		return errs.Wrap(err)
	}
	histChart := chart.ImportHistgramData(global, tokyo)

	if len(histCasesFile) > 0 {
		if err := chart.BarChartHistCases(histChart, histCasesFile); err != nil {
			return errs.Wrap(err)
		}
	}
	if len(histDeathsFile) > 0 {
		if err := chart.BarChartHistDeaths(histChart, histDeathsFile); err != nil {
			return errs.Wrap(err)
		}
	}
	return nil
}

func main() {
	lastDay := values.Yesterday()
	if err := makeGraph(values.NewDate(2020, time.Month(3), 11), lastDay, HistCasesFile, HistDeathsFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := imgutil.ConcatImageFiles(HistAllFile, HistCasesFile, HistDeathsFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	// experimental
	if err := makeGraph(values.NewDate(2021, time.Month(10), 1), lastDay, HistCasesFileShort, HistDeathsFileShort); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
}

/* Copyright 2020-2021 Spiegel
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
