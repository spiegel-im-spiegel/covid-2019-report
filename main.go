package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/covid-2019-report/chart"
	"github.com/spiegel-im-spiegel/covid-2019-report/imgutil"
	"github.com/spiegel-im-spiegel/covid-2019-report/report"
	"github.com/spiegel-im-spiegel/covid-2019-report/values"
)

const (
	casesFile        = "./covid-2019-new-cases-in-japan.png"
	fatalityRateFile = "./covid-2019-fatality-rate-in-japan.png"
	allFile          = "./covid-2019-cases-in-japan.png"
)

func main() {
	rps, err := report.ImportCSV(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	start, err := values.DateFrom("2020-03-11")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}

	if err := chart.BarChartNewCases(rps, start, casesFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := chart.LineChartFatalityRate(rps, start, fatalityRateFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := imgutil.ConcatImageFiles(allFile, []string{casesFile, fatalityRateFile}); err != nil {
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
