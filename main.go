package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

func getSystemLibrary() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	switch runtime.GOOS {
	case "darwin":
		return "/Applications/IDA Professional 9.0.app/Contents/MacOS/libidalib.dylib"
	case "linux":
		return filepath.Join(homedir, "ida-pro-9.0", "libidalib.so")
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func main() {
	idalib, err := purego.Dlopen(getSystemLibrary(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	var init_library func() int
	purego.RegisterLibFunc(&init_library, idalib, "init_library")

	var open_database func(string, bool) int
	purego.RegisterLibFunc(&open_database, idalib, "open_database")

	var close_database func(bool)
	purego.RegisterLibFunc(&close_database, idalib, "close_database")

	if ret := init_library(); ret != 0 {
		panic(fmt.Errorf("init_library()=%d", ret))
	}

	if ret := open_database("/Users/blacktop/RE/15.1_OLD_IDA/kernel.release.t6020.i64", false); ret != 0 {
		panic(fmt.Errorf("open_database()=%d", ret))
	}
	defer close_database(false)
}
