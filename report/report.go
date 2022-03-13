package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/goark/cov19jpn/chart"
	"github.com/goark/cov19jpn/entity"
	"github.com/goark/cov19jpn/fetch"
	"github.com/goark/cov19jpn/filter"
	"github.com/goark/cov19jpn/values/prefcodejpn"
)

func getPrefCodes() []prefcodejpn.Code {
	prefcodes := []prefcodejpn.Code{}
	for i := uint(1); ; i++ {
		c := prefcodejpn.Code(i)
		prefcodes = append(prefcodes, c)
		if c == prefcodejpn.PREFCODE_MAX {
			break
		}
	}
	return prefcodes
}

func run() error {
	//fetch data
	r, err := fetch.Web(context.Background(), &http.Client{})
	if err != nil {
		return err
	}
	defer r.Close()
	es, err := fetch.Import(r, filter.New())
	if err != nil {
		return err
	}
	list := entity.NewList(es)
	list.Sort()

	//csv
	if err := func() error {
		file, err := os.Create("./google-forecast/forecast_JAPAN_PREFECTURE_28.csv")
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(file, bytes.NewReader(list.EncodeCSV())); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return err
	}

	//plots
	for _, pref := range getPrefCodes() {
		sublist := list.Filer(filter.New(pref))
		hlist := chart.New(sublist.StartDayMeasure(), sublist.EndDayMeasure().AddDay(7), 7, sublist)
		filename := fmt.Sprintf("./google-forecast/%s-%s-cov19-forecast.png", pref.String(), pref.Name())
		if err := chart.MakeHistChart(hlist, pref.Title(), filename); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
