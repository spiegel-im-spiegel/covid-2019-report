package chart

import (
	"math"

	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type HistgramData struct {
	period     values.Period
	cases      float64
	deaths     float64
	casesTokyo float64
}

func newHistgramData(record *histogram.HistData) HistgramData {
	return HistgramData{
		period: record.Period,
		cases:  (float64)(record.Cases),
		deaths: (float64)(record.Deaths),
	}
}

func ImportHistgramData(global, tokyo []*histogram.HistData) []HistgramData {
	data := []HistgramData{}
	for _, h := range global {
		data = append(data, newHistgramData(h))
	}
	for _, h := range tokyo {
		for i := 0; i < len(data); i++ {
			if data[i].period.Contains(h.Period.End) {
				data[i].casesTokyo += (float64)(h.Cases)
			}
		}
	}
	return data
}

func BarChartHistCases(data []HistgramData, outPath string) error {
	labelX := []string{}
	dataY := plotter.Values{}
	dataY2 := plotter.XYs{}
	maxCases := 0.0
	for i, d := range data {
		labelX = append(labelX, d.period.StringEnd())
		dataY = append(dataY, d.cases)
		maxCases = max(maxCases, d.cases)
		dataY2 = append(dataY2, plotter.XY{X: (float64)(i), Y: d.casesTokyo})
		maxCases = max(maxCases, d.casesTokyo)
	}
	maxCases = (float64)((((int)(maxCases) / 400) + 1) * 400)

	//default font
	plot.DefaultFont = "Helvetica"
	plotter.DefaultFont = "Helvetica"

	//new plot
	p, err := plot.New()
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2)
	bar.Offset = 0
	bar.Horizontal = false
	p.Add(bar)

	//new line chart
	line, err := plotter.NewLine(dataY2)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
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
	p.Y.Label.Text = "Confirmed cases"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = maxCases

	//legend
	p.Legend.Add("New confirmed cases by 7 days", bar)
	p.Legend.Add("New positive PCR test results by 7 days in Tokyo", line)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "Confirmed COVID-2019 Cases in Japan"

	//output image
	if err := p.Save(20.0*(vg.Length)(len(data)+2), 10*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	return nil
}

func BarChartHistDeaths(data []HistgramData, outPath string) error {
	labelX := []string{}
	dataY := plotter.Values{}
	maxDeaths := 0.0
	for _, d := range data {
		labelX = append(labelX, d.period.StringEnd())
		dataY = append(dataY, d.deaths)
		maxDeaths = max(maxDeaths, d.deaths)
	}
	maxDeaths = (float64)((((int)(maxDeaths) / 200) + 1) * 200)

	//default font
	plot.DefaultFont = "Helvetica"
	plotter.DefaultFont = "Helvetica"

	//new plot
	p, err := plot.New()
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(10))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2)
	bar.Offset = 0
	bar.Horizontal = false
	p.Add(bar)

	//labels of X
	p.NominalX(labelX...)
	p.X.Label.Text = "Date of report"
	//p.X.Padding = 0
	p.X.Tick.Label.Rotation = math.Pi / 2.5
	p.X.Tick.Label.XAlign = draw.XRight
	p.X.Tick.Label.YAlign = draw.YCenter

	//labels of Y
	p.Y.Label.Text = "Deaths"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = maxDeaths

	//legend
	p.Legend.Add("New deaths by 7 days", bar)
	p.Legend.Top = true   //top
	p.Legend.Left = false //left
	p.Legend.XOffs = -5
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "COVID-2019 deaths in Japan"

	//output image
	if err := p.Save(20.0*(vg.Length)(len(data)+2), 8*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
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