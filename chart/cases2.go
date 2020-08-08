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

type newCasesData2 struct {
	date          string
	newCases      float64
	newCasesTokyo float64
	newDeaths     float64
}

func newNewCasesData2(rp report.Report) newCasesData2 {
	return newCasesData2{
		date:          rp.Date().String(),
		newCases:      (float64)(rp.CasesByDay()),
		newCasesTokyo: (float64)(rp.CasesByDayTokyo()),
		newDeaths:     (float64)(rp.DeathsByDay()),
	}
}

func importNewCasesData2(rps report.Reports, start, end values.Date) ([]newCasesData2, error) {
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
	data := []newCasesData2{newNewCasesData2(rp)}
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
		data = append(data, newNewCasesData2(rp))
	}
	return data, nil
}

func BarChartNewCases2(rps report.Reports, start, end values.Date, outPath string) error {
	data, err := importNewCasesData2(rps, start, end)
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	labelX := []string{}
	dataY1 := plotter.Values{}
	dataY1b := plotter.XYs{}
	dataY2 := plotter.Values{}
	maxCases := 0.0
	for i, d := range data {
		labelX = append(labelX, d.date)
		dataY1 = append(dataY1, d.newCases)
		maxCases = max(maxCases, d.newCases)
		dataY1b = append(dataY1b, plotter.XY{X: (float64)(i), Y: d.newCasesTokyo})
		maxCases = max(maxCases, d.newCasesTokyo)
		dataY2 = append(dataY2, d.newDeaths)
		maxCases = max(maxCases, d.newDeaths)
	}
	maxCases = (float64)((((int)(maxCases) / 100) + 1) * 100)

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

	//new line chart
	line, err := plotter.NewLine(dataY1b)
	if err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	line.Color = plotutil.Color(4)
	p.Add(line)

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
	p.Y.Max = maxCases

	//legend
	p.Legend.Add("New confirmed cases by day", bar1)
	p.Legend.Add("New positive PCR test results by day in Tokyo", line)
	p.Legend.Add("New deaths by day", bar2)
	p.Legend.Top = true  //top
	p.Legend.Left = true //left
	p.Legend.XOffs = 0
	p.Legend.YOffs = -10

	//title
	p.Title.Text = "Confirmed COVID-2019 Cases in Japan"

	//output image
	if err := p.Save(20.0*(vg.Length)(len(data)+2), 30*vg.Centimeter, outPath); err != nil {
		return errs.Wrap(err, "", errs.WithContext("start", start), errs.WithContext("outPath", outPath))
	}
	return nil
}
