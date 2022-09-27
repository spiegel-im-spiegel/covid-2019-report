package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goark/cov19data/values"
	"github.com/spiegel-im-spiegel/covid-2019-report/report/matsue/excel"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	infections, err := excel.NewInfections(args[0], "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	startTm, err := values.NewDateString("2022-01-01")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}

	s, err := excel.ConvDot(infections, startTm, values.Today())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	fmt.Println(s)
}

/* Copyright 2021-2022 Spiegel
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
