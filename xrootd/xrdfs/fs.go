// Copyright ©2018 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xrdfs contains structures representing the XRootD-based filesystem.
package xrdfs

import (
	"context"
)

// FileSystem implements access to a collection of named files over XRootD.
type FileSystem interface {
	// Dirlist returns the contents of a directory together with the stat information.
	Dirlist(ctx context.Context, path string) ([]EntryStat, error)

	// Open returns the file handle for a file together with the compression and the stat info.
	Open(ctx context.Context, path string, mode OpenMode, options OpenOptions) (File, error)

	// RemoveFile removes the file at path.
	RemoveFile(ctx context.Context, path string) error

	// Truncate changes the size of the named file.
	Truncate(ctx context.Context, path string, size int64) error

	// Stat returns the entry stat info for the given path.
	Stat(ctx context.Context, path string) (EntryStat, error)

	// VirtualStat returns the virtual filesystem stat info for the given path.
	// Note that path needs not to be an existing filesystem object, it is used as a path prefix in order to
	// filter out servers and partitions that could not be used to hold objects whose path starts
	// with the specified path prefix.
	VirtualStat(ctx context.Context, path string) (VirtualFSStat, error)

	// Mkdir creates a new directory with the specified name and permission bits.
	Mkdir(ctx context.Context, path string, perm OpenMode) error

	// MkdirAll creates a directory named path, along with any necessary parents,
	// and returns nil, or else returns an error.
	// The permission bits perm are used for all directories that MkdirAll creates.
	MkdirAll(ctx context.Context, path string, perm OpenMode) error

	// RemoveDir removes a directory.
	// The directory to be removed must be empty.
	RemoveDir(ctx context.Context, path string) error

	// RemoveAll removes path and any children it contains.
	// It removes everything it can but returns the first error it encounters.
	// If the path does not exist, RemoveAll returns nil (no error.)
	RemoveAll(ctx context.Context, path string) error

	// Rename renames (moves) oldpath to newpath.
	Rename(ctx context.Context, oldpath, newpath string) error

	// Chmod changes the permissions of the named file to perm.
	Chmod(ctx context.Context, path string, mode OpenMode) error

	// Statx obtains type information for one or more paths.
	// Only a limited number of flags is meaningful such as StatIsExecutable, StatIsDir, StatIsOther, StatIsOffline.
	Statx(ctx context.Context, paths []string) ([]StatFlags, error)
}

// OpenMode is the mode in which path is to be opened.
// The mode is an "or`d" combination of ModeXyz flags.
type OpenMode uint16

const (
	OpenModeOwnerRead    OpenMode = 0x100 // OpenModeOwnerRead indicates that owner has read access.
	OpenModeOwnerWrite   OpenMode = 0x080 // OpenModeOwnerWrite indicates that owner has write access.
	OpenModeOwnerExecute OpenMode = 0x040 // OpenModeOwnerExecute indicates that owner has execute access.

	OpenModeGroupRead    OpenMode = 0x020 // OpenModeGroupRead indicates that group has read access.
	OpenModeGroupWrite   OpenMode = 0x010 // OpenModeGroupWrite indicates that group has write access.
	OpenModeGroupExecute OpenMode = 0x008 // OpenModeGroupExecute indicates that group has execute access.

	OpenModeOtherRead    OpenMode = 0x004 // OpenModeOtherRead indicates that owner has read access.
	OpenModeOtherWrite   OpenMode = 0x002 // OpenModeOtherWrite indicates that owner has write access.
	OpenModeOtherExecute OpenMode = 0x001 // OpenModeOtherExecute indicates that owner has execute access.
)

// OpenOptions are the options to apply when path is opened.
type OpenOptions uint16

const (
	// OpenOptionsCompress specifies that file is opened even when compressed.
	OpenOptionsCompress OpenOptions = 1 << iota
	// OpenOptionsDelete specifies that file is opened deleting any existing file.
	OpenOptionsDelete
	// OpenOptionsForce specifies that file is opened ignoring  file usage rules.
	OpenOptionsForce
	// OpenOptionsNew specifies that file is opened only if it does not already exist.
	OpenOptionsNew
	// OpenOptionsOpenRead specifies that file is opened only for reading.
	OpenOptionsOpenRead
	// OpenOptionsOpenUpdate specifies that file is opened only for reading and writing.
	OpenOptionsOpenUpdate
	// OpenOptionsAsync specifies that file is opened for asynchronous i/o.
	OpenOptionsAsync
	// OpenOptionsRefresh specifies that cached information on the file's location need to be updated.
	OpenOptionsRefresh
	// OpenOptionsMkPath specifies that directory path is created if it does not already exist.
	OpenOptionsMkPath
	// OpenOptionsOpenAppend specifies that file is opened only for appending.
	OpenOptionsOpenAppend
	// OpenOptionsReturnStatus specifies that file status information should be returned in the response.
	OpenOptionsReturnStatus
	// OpenOptionsReplica specifies that file is opened for replica creation.
	OpenOptionsReplica
	// OpenOptionsPOSC specifies that Persist On Successful Close (POSC) processing should be enabled.
	OpenOptionsPOSC
	// OpenOptionsNoWait specifies that file is opened only if it does not cause a wait.
	OpenOptionsNoWait
	// OpenOptionsSequentiallyIO specifies that file will be read or written sequentially.
	OpenOptionsSequentiallyIO
	// OpenOptionsNone specifies that file is opened without specific options.
	OpenOptionsNone OpenOptions = 0
)
