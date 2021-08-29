package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
	"github.com/spiegel-im-spiegel/covid-2019-report/report/matsue/excel"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
	}
	xlsx, no, err := excel.OpenXlsxFileSheet(args[0], "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	infections, err := excel.NewInfections(xlsx, no)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	s, err := excel.ConvDot(infections, date.FromString("2020-10-25"), date.Today())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	fmt.Println(s)
}
