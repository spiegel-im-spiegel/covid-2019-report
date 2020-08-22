package report

import (
	"io"
)

func ImportCSV(readerJp, readerTokyo io.Reader) (Reports, error) {
	cs, err := importWHOCSV(readerJp)
	if err != nil {
		return nil, err
	}
	csTokyo, err := importTokyoCSV(readerTokyo)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(cs); i++ {
		if ct, ok := csTokyo[cs[i].Date.String()]; ok {
			//fmt.Printf("%v: %v\n", cs[i].Date, ct)
			cs[i].NewsTokyo = ct
		}
	}
	return New(cs), nil
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
