package main

import (
	"context"

	"github.com/goark/cov19data"
	"github.com/goark/cov19data/filter"
	"github.com/goark/cov19data/histogram"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

func getGlobalHist(p values.Period) ([]*histogram.HistData, error) {
	impt, err := cov19data.NewWeb(context.Background(), fetch.New())
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer impt.Close()
	return impt.Histogram(
		p,
		7,
		filter.WithCountryCode(values.CC_JP),
		filter.WithRegionCode(values.WPRO),
	)
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
