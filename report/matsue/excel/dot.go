package excel

import (
	"fmt"
	"strings"

	"github.com/spiegel-im-spiegel/cov19jpn/values/date"
)

const (
	header = `digraph infections {
	graph [
		charset = "UTF-8",
		layout = fdp
	];
	node [
		fontname="Inconsolata",
		fontcolor = black,
		style = "solid,filled",
		color = black
		fillcolor = white
	];
	edge [
		color = coral3
	];
`
	footer = "}"
)

func ConvDot(infections []*Infection, start, end date.Date) (string, error) {
	if len(infections) == 0 {
		return "", nil
	}

	lastDay := infections[len(infections)-1].Date
	declr := &strings.Builder{}
	rel := &strings.Builder{}
	outside := map[string]string{}
	for _, infection := range infections {
		if !start.IsZero() && infection.Date.Before(start) {
			continue
		}
		if !end.IsZero() && infection.Date.After(end) {
			continue
		}
		if infection.Date.Before(lastDay.AddDay(-7)) {
			if !infection.InsideFlag || len(infection.FromOutside) > 0 {
				declr.WriteString(fmt.Sprintf("\t%s[color=crimson]\n", infection.NodeMatsue))
			} else {
				declr.WriteString(fmt.Sprintf("\t%s\n", infection.NodeMatsue))
			}
		} else if !infection.InsideFlag {
			declr.WriteString(fmt.Sprintf("\t%s[color=crimson]\n", infection.NodeMatsue))
		} else if len(infection.FromOutside) > 0 {
			declr.WriteString(fmt.Sprintf("\t%s[color=crimson,fillcolor=burlywood1]\n", infection.NodeMatsue))
		} else {
			declr.WriteString(fmt.Sprintf("\t%s[fillcolor=burlywood1]\n", infection.NodeMatsue))
		}
		for _, node := range infection.FromOutside {
			if !strings.EqualFold(node, "other") {
				outside[node] = ""
				rel.WriteString(fmt.Sprintf("\t%s->%s\n", node, infection.NodeMatsue))
			}
		}
		for _, node := range infection.FromInside {
			rel.WriteString(fmt.Sprintf("\t%s->%s\n", node, infection.NodeMatsue))
		}
	}
	for k := range outside {
		declr.WriteString(fmt.Sprintf("\t%s[color=crimson]\n", k))
	}

	bldr := &strings.Builder{}
	bldr.WriteString(header)
	bldr.WriteString(declr.String())
	bldr.WriteString(rel.String())
	bldr.WriteString(footer)
	return bldr.String(), nil
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
