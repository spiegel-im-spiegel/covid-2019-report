package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spiegel-im-spiegel/covid-2019-report/chart"
	"github.com/spiegel-im-spiegel/covid-2019-report/imgutil"
	"github.com/spiegel-im-spiegel/covid-2019-report/report"
	"github.com/spiegel-im-spiegel/covid-2019-report/values"
)

const (
	casesFile         = "./covid-2019-new-cases-in-japan.png"
	casesFile2        = "./covid-2019-new-cases-in-japan2.png"
	fatalityRateFile  = "./covid-2019-fatality-rate-in-japan.png"
	fatalityRateFile2 = "./covid-2019-fatality-rate-in-japan2.png"
	allFile           = "./covid-2019-cases-in-japan.png"
	allFile2          = "./covid-2019-cases-in-japan2.png"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	whoCSV := os.Args[1]
	tokyoCSV := os.Args[2]
	fmt.Println("build chart by:", whoCSV, tokyoCSV)

	csvFile1, err := os.Open(whoCSV)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	defer csvFile1.Close()

	csvFile2, err := os.Open(tokyoCSV)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	defer csvFile2.Close()

	rps, err := report.ImportCSV(csvFile1, csvFile2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	start, err := values.DateFrom("2020-03-11")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	end, err := values.DateFrom("2020-05-25")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}

	if err := chart.BarChartNewCases(rps, start, end, casesFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := chart.LineChartFatalityRate(rps, start, end, fatalityRateFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := imgutil.ConcatImageFiles(allFile, casesFile, fatalityRateFile); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}

	start, err = values.DateFrom("2020-05-25")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	end = values.NewDate(time.Time{})

	if err := chart.BarChartNewCases2(rps, start, end, casesFile2); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := chart.LineChartFatalityRate(rps, start, end, fatalityRateFile2); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if err := imgutil.ConcatImageFiles(allFile2, casesFile2, fatalityRateFile2); err != nil {
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
