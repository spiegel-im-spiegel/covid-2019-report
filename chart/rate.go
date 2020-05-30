package chart

import (
	"errors"
	"math"

	"github.com/spiegel-im-spiegel/covid-2019-report/ecode"
	"github.com/spiegel-im-spiegel/covid-2019-report/report"
	"github.com/spiegel-im-spiegel/covid-2019-report/values"
	"github.com/spiegel-im-spiegel/errs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type fatalityRateData struct {
	date         string
	fatalityRate float64
}

func newFatalityRateData(rp report.Report) fatalityRateData {
	return fatalityRateData{
		date:         rp.Date().String(),
		fatalityRate: rp.FatalityRate(),
	}
}

func importFatalityRateData(rps report.Reports, start, end values.Date) ([]fatalityRateData, error) {
	var rp report.Report
	var err error
	if start.IsZero() {
		rp, err = rps.Top()
	} else {
		rp, err = rps.SearchByDate(start)
	}
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("end", end))
	}
	data := []fatalityRateData{newFatalityRateData(rp)}
	for {
		rp, err := rps.Next()
		if err != nil {
			if errors.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("end", end))
		}
		if !end.IsZero() && rp.Date().After(end) {
			break
		}
		data = append(data, newFatalityRateData(rp))
	}
	return data, nil
}

func LineChartFatalityRate(rps report.Reports, start, end values.Date, outPath string) error {
	data, err := importFatalityRateData(rps, start, end)
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	labelX := []string{}
	dataY := plotter.XYs{}
	for i, d := range data {
		labelX = append(labelX, d.date)
		dataY = append(dataY, plotter.XY{X: (float64)(i), Y: d.fatalityRate * 100.0})
	}

	//default font
	plot.DefaultFont = "Helvetica"
	plotter.DefaultFont = "Helvetica"

	//new plot
	p, err := plot.New()
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}

	//new line chart
	line, err := plotter.NewLine(dataY)
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	line.Color = plotutil.Color(4)
	p.Add(line)

	//labels of X
	p.NominalX(labelX...)
	p.X.Label.Text = "Date of report"
	//p.X.Padding = 0
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter

	//labels of Y
	p.Y.Label.Text = "Fatality rate (%)"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = 6.5

	//legend
	p.Legend.Add("Fatality rate", line)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "COVID-2019 Fatality Rate in Japan"

	//output image
	if err := p.Save(20.0*(vg.Length)(len(data)+2), 15*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	return nil
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
