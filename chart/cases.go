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

type newCasesData struct {
	date      string
	newCases  float64
	newDeaths float64
}

func newNewCasesData(rp report.Report) newCasesData {
	return newCasesData{
		date:      rp.Date().String(),
		newCases:  (float64)(rp.CasesByDay()),
		newDeaths: (float64)(rp.DeathsByDay()),
	}
}

func importNewCasesData(rps report.Reports, start values.Date) ([]newCasesData, error) {
	var rp report.Report
	var err error
	if start.IsZero() {
		rp, err = rps.Top()
	} else {
		rp, err = rps.SearchByDate(start)
	}
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("start", start))
	}
	data := []newCasesData{newNewCasesData(rp)}
	for {
		rp, err := rps.Next()
		if err != nil {
			if errors.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err, "", errs.WithContext("start", start))
		}
		data = append(data, newNewCasesData(rp))
	}
	return data, nil
}

func BarChartNewCases(rps report.Reports, start values.Date, outPath string) error {
	data, err := importNewCasesData(rps, start)
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	labelX := []string{}
	dataY1 := plotter.Values{}
	dataY2 := plotter.Values{}
	for _, d := range data {
		labelX = append(labelX, d.date)
		dataY1 = append(dataY1, d.newCases)
		dataY2 = append(dataY2, d.newDeaths)
	}

	//default font
	plot.DefaultFont = "Helvetica"
	plotter.DefaultFont = "Helvetica"

	//new plot
	p, err := plot.New()
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}

	//new bar chart
	bar1, err := plotter.NewBarChart(dataY1, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	bar1.LineStyle.Width = vg.Length(0)
	bar1.Color = plotutil.Color(2)
	bar1.Offset = -2
	bar1.Horizontal = false
	p.Add(bar1)
	bar2, err := plotter.NewBarChart(dataY2, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	bar2.LineStyle.Width = vg.Length(0)
	bar2.Color = plotutil.Color(6)
	bar2.Offset = 2
	bar2.Horizontal = false
	p.Add(bar2)

	//labels of X
	p.NominalX(labelX...)
	p.X.Label.Text = "Date of report"
	//p.X.Padding = 0
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter

	//labels of Y
	p.Y.Label.Text = "Confirmed cases"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = 800

	//legend
	p.Legend.Add("New confirmed cases by day", bar1)
	p.Legend.Add("New deaths by day", bar2)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "Confirmed COVID-2019 Cases in Japan"

	//output image
	if err := p.Save(20.0*(vg.Length)(len(data)), 15*vg.Centimeter, outPath); err != nil {
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
