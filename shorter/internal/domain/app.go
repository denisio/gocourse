package domain

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	AppDir = filepath.Join(filepath.Dir(b), "../..")
)
