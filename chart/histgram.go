package chart

import (
	"math"

	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
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
		cases:  record.Cases,
		deaths: record.Deaths,
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
				data[i].casesTokyo += h.Cases
			}
		}
	}
	return data
}

func ImportHistgramDataByPref(pref []*histogram.HistData) []HistgramData {
	data := []HistgramData{}
	for _, h := range pref {
		data = append(data, newHistgramData(h))
	}
	return data
}

func BarChartHistCases(data []HistgramData, outPath string) error {
	labelX := []string{}
	dataY := plotter.Values{}
	dataY2 := plotter.XYs{}
	dataY3 := plotter.Values{}
	maxCases := 0.0
	for i, d := range data {
		labelX = append(labelX, d.period.StringEnd())
		dataY = append(dataY, d.cases)
		maxCases = max(maxCases, d.cases)
		dataY3 = append(dataY3, d.deaths)
		maxCases = max(maxCases, d.deaths)
		dataY2 = append(dataY2, plotter.XY{X: (float64)(i), Y: d.casesTokyo})
		maxCases = max(maxCases, d.casesTokyo)
	}
	maxCases = (float64)((((int)(maxCases) / 400) + 1) * 400)

	//default font
	plot.DefaultFont = font.Font{
		Typeface: "Liberation",
		Variant:  "Sans",
	}

	//new plot
	p := plot.New()

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(5))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2)
	bar.Offset = 0
	bar.Horizontal = false
	p.Add(bar)
	bar3, err := plotter.NewBarChart(dataY3, vg.Points(5))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar3.LineStyle.Width = vg.Length(0)
	bar3.Color = plotutil.Color(7)
	bar3.Offset = 0
	bar3.Horizontal = false
	p.Add(bar3)

	//new line chart
	line, err := plotter.NewLine(dataY2)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	line.Color = plotutil.Color(3)
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
	p.Legend.Add("New deaths by 7 days", bar3)
	p.Legend.Add("New positive PCR test results by 7 days in Tokyo", line)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "Confirmed COVID-2019 Cases in Japan"

	//output image
	if err := p.Save(10.0*(vg.Length)(len(data)+2), 10*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	return nil
}

func BarChartHistCasesByPref(data []HistgramData, prefName, outPath string) error {
	labelX := []string{}
	dataY := plotter.Values{}
	maxCases := 0.0
	for _, d := range data {
		labelX = append(labelX, d.period.StringEnd())
		dataY = append(dataY, d.cases)
		maxCases = max(maxCases, d.cases)
	}
	maxCases = (float64)((((int)(maxCases) / 100) + 1) * 100)

	//default font
	plot.DefaultFont = font.Font{
		Typeface: "Liberation",
		Variant:  "Sans",
	}

	//new plot
	p := plot.New()

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(5))
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
	p.Y.Label.Text = "Confirmed cases"
	p.Y.Padding = 5
	p.Y.Min = 0
	p.Y.Max = maxCases

	//legend
	p.Legend.Add("New Cases", bar)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = 0

	//title
	p.Title.Text = "Cases in " + prefName

	//output image
	if err := p.Save(10.0*(vg.Length)(len(data)+2), 8*vg.Centimeter, outPath); err != nil {
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
	plot.DefaultFont = font.Font{
		Typeface: "Liberation",
		Variant:  "Sans",
	}

	//new plot
	p := plot.New()

	//new bar chart
	bar, err := plotter.NewBarChart(dataY, vg.Points(5))
	if err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(7)
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
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = 0

	//title
	p.Title.Text = "COVID-2019 deaths in Japan"

	//output image
	if err := p.Save(10.0*(vg.Length)(len(data)+2), 8*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, errs.WithContext("outPath", outPath))
	}
	return nil
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
