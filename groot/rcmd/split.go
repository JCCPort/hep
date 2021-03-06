// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rcmd

import (
	"fmt"
	"log"
	"path"

	"go-hep.org/x/hep/groot"
	"go-hep.org/x/hep/groot/riofs"
	"go-hep.org/x/hep/groot/rtree"
)

// Split splits the tree from the input file into multiple trees,
// each with nevts entries.
// Split returns the name of the split output files, and an error, if any.
func Split(oname, fname, tname string, nevts int64, verbose bool) ([]string, error) {
	f, err := groot.Open(fname)
	if err != nil {
		return nil, fmt.Errorf(
			"could not open input file %q: %w",
			fname, err,
		)
	}
	defer f.Close()

	o, err := riofs.Dir(f).Get(tname)
	if err != nil {
		return nil, fmt.Errorf(
			"could not fet tree %q: %w", tname, err,
		)
	}

	tree, ok := o.(rtree.Tree)
	if !ok {
		return nil, fmt.Errorf("object %q is not a Tree", tname)
	}

	var (
		cur    int64
		n      = tree.Entries()
		nfiles = 0
	)
	for i := 0; cur < n; i++ {
		m, err := split(oname, tname, tree, i, cur, nevts, verbose)
		if err != nil {
			return nil, fmt.Errorf("could not split tree into file#%d: %w", i, err)
		}
		cur += m
		nfiles++
	}

	onames := make([]string, nfiles)
	for i := range onames {
		onames[i] = fmt.Sprintf(
			"%s-%d.root",
			oname[:len(oname)-len(".root")], i,
		)
	}

	return onames, nil
}

func split(oname, tname string, tree rtree.Tree, i int, beg, nevts int64, verbose bool) (int64, error) {
	oname = fmt.Sprintf(
		"%s-%d.root",
		oname[:len(oname)-len(".root")], i,
	)
	o, err := groot.Create(oname)
	if err != nil {
		return 0, fmt.Errorf("could not create output file %d: %w", i, err)
	}
	defer o.Close()

	var (
		dirName = path.Dir(tname)
		objName = path.Base(tname)
		dir     = riofs.Directory(o)
	)
	if dirName != "/" && dirName != "" && dirName != "." {
		_, err = riofs.Dir(o).Mkdir(dirName)
		if err != nil {
			return 0, fmt.Errorf("could not create output directory %q: %w", dirName, err)
		}
		odir, err := riofs.Dir(o).Get(dirName)
		if err != nil {
			return 0, fmt.Errorf("could not fetch output directory %q: %w", dirName, err)
		}
		dir = odir.(riofs.Directory)
	}
	wvars := rtree.WriteVarsFromTree(tree)
	w, err := rtree.NewWriter(
		dir, objName,
		wvars,
		rtree.WithTitle(tree.Title()),
	)
	if err != nil {
		return 0, fmt.Errorf("could not create tree writer: %w", err)
	}
	defer w.Close()

	var (
		tot   int64
		rvars = make([]rtree.ReadVar, len(wvars))
		src   = tree
		max   = beg + nevts
	)
	for i, wvar := range wvars {
		rvars[i] = rtree.ReadVar{
			Name:  wvar.Name,
			Value: wvar.Value,
		}
	}

	scan, err := rtree.NewScannerVars(src, rvars...)
	if err != nil {
		return 0, fmt.Errorf("could not create scanner: %w", err)
	}
	err = scan.SeekEntry(beg)
	if err != nil {
		return 0, fmt.Errorf("could not seek to entry %d: %w", beg, err)
	}

	if max > tree.Entries() {
		max = tree.Entries()
	}

	if verbose {
		log.Printf("splitting [%d, %d) into %q...", beg, max, oname)
	}

	for scan.Next() && scan.Entry() < max {
		err = scan.Scan()
		if err != nil {
			return tot, fmt.Errorf("rtree: could not read entry %d from tree: %w", scan.Entry(), err)
		}

		_, err = w.Write()
		if err != nil {
			return tot, fmt.Errorf("rtree: could not write entry %d to tree: %w", scan.Entry(), err)
		}
		tot++
	}

	err = scan.Err()
	if err != nil {
		return tot, fmt.Errorf("rtree: could not scan through tree: %w", err)
	}

	if verbose {
		log.Printf("splitting [%d, %d) into %q... [ok]", beg, max, oname)
	}

	return tot, nil
}
