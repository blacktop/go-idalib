package idalib

// #cgo CXXFLAGS: -I${SRCDIR}/include -std=c++11
// #cgo darwin LDFLAGS: -L${SRCDIR}/idasdk90/lib/arm64_mac_clang_64 -Wl,-rpath,'/Applications/IDA Professional 9.0.app/Contents/MacOS' -lida -lidalib
import "C"
