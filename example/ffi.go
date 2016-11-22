package main

import "../lua"
import "fmt"

var lcode string = 
`local ffi = require("ffi")
ffi.cdef[[
int printf(const char *fmt, ...);
]]
ffi.C.printf("Hello %s from FFI!\n", "world")`

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()
	
	// this one is needed to enable the FFI module
	L.OpenFFI()

	err := L.DoString(lcode)

	// Should print nil if there is no error
	fmt.Printf("Ciao %v\n", err)
}