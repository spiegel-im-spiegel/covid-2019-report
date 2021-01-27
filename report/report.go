package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/google"
	"github.com/spiegel-im-spiegel/cov19data/google/entity"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/covid-2019-report/chart"
	"github.com/spiegel-im-spiegel/fetch"
)

func getAllData() ([]*entity.JapanData, error) {
	impt, err := google.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, err
	}
	defer impt.Close()
	return impt.Data()
}

func outputFile(path string, dat []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, bytes.NewReader(dat)); err != nil {
		return err
	}
	return nil

}

func exportCsv(data []*entity.JapanData, pc values.PrefJpCode, path string) error {
	b, err := entity.ExportCSV(
		data,
		filter.WithPeriod(
			values.NewPeriod(
				values.Yesterday().AddDay(-27),
				values.Yesterday().AddDay(6),
			),
		),
		filter.WithPrefJpCode(pc),
	)
	if err != nil {
		return err
	}
	return outputFile(path, b)
}

func exportHistData(data []*entity.JapanData, pc values.PrefJpCode) ([]*histogram.HistData, error) {
	return entity.ExportHistgram(
		data,
		values.NewPeriod(
			values.Yesterday().AddDay(-27),
			values.Yesterday().AddDay(6),
		),
		7,
		filter.WithPrefJpCode(pc),
	)
}

func main() {
	data, err := getAllData()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//shimane data
	if err := exportCsv(data, values.PrefJpCode(32), "shimane/shimane-covid19-cases-recently.csv"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if hist, err := exportHistData(data, values.PrefJpCode(32)); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else if err := chart.BarChartHistCasesByPref(chart.ImportHistgramDataByPref(hist), "Shimane", "shimane/shimane-covid19-cases-histgram.png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// if err := exportHistCsv(data, values.PrefJpCode(32), "shimane/shimane-covid19-cases-histgram.csv"); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }
	//hiroshima data
	if err := exportCsv(data, values.PrefJpCode(34), "hiroshima/hiroshima-covid19-cases-recently.csv"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if hist, err := exportHistData(data, values.PrefJpCode(34)); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else if err := chart.BarChartHistCasesByPref(chart.ImportHistgramDataByPref(hist), "Hiroshima", "hiroshima/hiroshima-covid19-cases-histgram.png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
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
