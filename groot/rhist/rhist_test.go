// Copyright 2018 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rhist

import (
	"go-hep.org/x/hep/groot/internal/rtests"
	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rcont"
	"go-hep.org/x/hep/groot/root"
)

var HistoTestCases = []struct {
	Name string
	Skip bool
	Want rtests.ROOTer
}{
	{
		Name: "TH1F",
		Want: &H1F{
			th1: th1{
				Named:     *rbase.NewNamed("h1f", "my-title"),
				attline:   rbase.AttLine{Color: 602, Style: 1, Width: 1},
				attfill:   rbase.AttFill{Color: 0, Style: 1001},
				attmarker: rbase.AttMarker{Color: 1, Style: 1, Width: 1},
				ncells:    102,
				xaxis: taxis{
					Named: *rbase.NewNamed("xaxis", ""),
					attaxis: rbase.AttAxis{
						Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.005, LabelSize: 0.035, Ticks: 0.03, TitleOffset: 1, TitleSize: 0.035, TitleColor: 1, TitleFont: 42,
					},
					nbins: 100, xmin: 0, xmax: 100,
					xbins: rcont.ArrayD{Data: nil},
					first: 0, last: 0, bits2: 0x0, time: false, tfmt: "",
					labels:  nil,
					modlabs: nil,
				},
				yaxis: taxis{
					Named: *rbase.NewNamed("yaxis", ""),
					attaxis: rbase.AttAxis{
						Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.005, LabelSize: 0.035, Ticks: 0.03, TitleOffset: 1, TitleSize: 0.035, TitleColor: 1, TitleFont: 42,
					},
					nbins: 1, xmin: 0, xmax: 1,
					xbins: rcont.ArrayD{Data: nil},
					first: 0, last: 0, bits2: 0x0, time: false, tfmt: "",
					labels:  nil,
					modlabs: nil,
				},
				zaxis: taxis{
					Named: *rbase.NewNamed("zaxis", ""),
					attaxis: rbase.AttAxis{
						Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.005, LabelSize: 0.035, Ticks: 0.03, TitleOffset: 1, TitleSize: 0.035, TitleColor: 1, TitleFont: 42,
					},
					nbins: 1, xmin: 0, xmax: 1,
					xbins: rcont.ArrayD{Data: nil},
					first: 0, last: 0, bits2: 0x0, time: false, tfmt: "",
					labels:  nil,
					modlabs: nil,
				},
				boffset: 0, bwidth: 1000,
				entries: 10,
				tsumw:   10, tsumw2: 16, tsumwx: 278, tsumwx2: 9286,
				max: -1111, min: -1111,
				norm:    0,
				contour: rcont.ArrayD{Data: nil},
				sumw2: rcont.ArrayD{
					Data: []float64{
						1,
						0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
						1,
					},
				},
				opt:    "",
				funcs:  *rcont.NewList("", []root.Object{}),
				buffer: nil,
				erropt: 0,
			},
			arr: rcont.ArrayF{
				Data: []float32{
					1,
					0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					1,
				},
			},
		},
	},
	{
		Name: "TH2F",
		Want: &H2F{
			th2: th2{
				th1: th1{
					Named:     *rbase.NewNamed("h2f", "my title"),
					attline:   rbase.AttLine{Color: 602, Style: 1, Width: 1},
					attfill:   rbase.AttFill{Color: 0, Style: 1001},
					attmarker: rbase.AttMarker{Color: 1, Style: 1, Width: 1},
					ncells:    144,
					xaxis: taxis{
						Named: *rbase.NewNamed("xaxis", ""),
						attaxis: rbase.AttAxis{
							Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.004999999888241291, LabelSize: 0.03500000014901161,
							Ticks: 0.029999999329447746, TitleOffset: 1, TitleSize: 0.03500000014901161, TitleColor: 1, TitleFont: 42,
						},
						nbins:   10,
						xmin:    0,
						xmax:    10,
						xbins:   rcont.ArrayD{},
						first:   0,
						last:    0,
						bits2:   0x0,
						time:    false,
						tfmt:    "",
						labels:  nil,
						modlabs: nil,
					},
					yaxis: taxis{
						Named: *rbase.NewNamed("yaxis", ""),
						attaxis: rbase.AttAxis{
							Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.004999999888241291, LabelSize: 0.03500000014901161,
							Ticks: 0.029999999329447746, TitleOffset: 1, TitleSize: 0.03500000014901161, TitleColor: 1, TitleFont: 42,
						},
						nbins:   10,
						xmin:    0,
						xmax:    10,
						xbins:   rcont.ArrayD{},
						first:   0,
						last:    0,
						bits2:   0x0,
						time:    false,
						tfmt:    "",
						labels:  nil,
						modlabs: nil,
					},
					zaxis: taxis{
						Named: *rbase.NewNamed("zaxis", ""),
						attaxis: rbase.AttAxis{
							Ndivs: 510, AxisColor: 1, LabelColor: 1, LabelFont: 42, LabelOffset: 0.004999999888241291, LabelSize: 0.03500000014901161,
							Ticks: 0.029999999329447746, TitleOffset: 1, TitleSize: 0.03500000014901161, TitleColor: 1, TitleFont: 42,
						},
						nbins:   1,
						xmin:    0,
						xmax:    1,
						xbins:   rcont.ArrayD{},
						first:   0,
						last:    0,
						bits2:   0x0,
						time:    false,
						tfmt:    "",
						labels:  nil,
						modlabs: nil,
					},
					boffset: 0,
					bwidth:  1000,
					entries: 13,
					tsumw:   9,
					tsumw2:  29,
					tsumwx:  21,
					tsumwx2: 55,
					max:     -1111,
					min:     -1111,
					norm:    0,
					contour: rcont.ArrayD{},
					sumw2: rcont.ArrayD{
						Data: []float64{
							1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
							0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 25, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
							0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
						},
					},
					opt:    "",
					funcs:  *rcont.NewList("", []root.Object{}),
					buffer: nil,
					erropt: 0,
				},
				scale:   1,
				tsumwy:  21,
				tsumwy2: 55,
				tsumwxy: 55,
			},
			arr: rcont.ArrayF{
				Data: []float32{
					1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 1, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,
					0, 1, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0,
					0, 0, 0, 0, 0, 1,
				},
			},
		},
	},
}
