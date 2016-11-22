Go Bindings for the lua C API
=========================

Simplest way to install:

	# go get -u github.com/dnestorov/golua/lua

Will work as long as you point CGO to the required libraries:

On Windows

	# set CGO_LDFLAGS=-L<path_to_lua_lib> -L<path_to_libluaffifb>
	
On Linux/UNIX 

	# set CGO_LDFLAGS=-L<path_to_lua_lib> -L<path_to_libluaffifb>

Or add the libraries path to lua.go

	#cgo CFLAGS: -Ilua -Iffi

You can then try to run the examples:

	$ cd /usr/local/go/src/pkg/github.com/aarzilli/golua/example/
	$ go run basic.go
	$ go run alloc.go
	$ go run panic.go
	$ go run userdata.go

QUICK START
---------------------

Create a new Virtual Machine with:

```go
L := lua.NewState()
L.OpenLibs()
defer L.Close()
```

Lua's Virtual Machine is stack based, you can call lua functions like this:

```go
// push "print" function on the stack
L.GetGlobal("print")
// push the string "Hello World!" on the stack
L.PushString("Hello World!")
// call print with one argument, expecting no results
L.Call(1, 0)
```

Of course this isn't very useful, more useful is executing lua code from a file or from a string:

```go
// executes a string of lua code
err := L.DoString("...")
// executes a file
err = L.DoFile(filename)
```

You will also probably want to publish go functions to the virtual machine, you can do it by:

```go
func adder(L *lua.State) int {
	a := L.ToInteger(1)
	b := L.ToInteger(2)
	L.PushInteger(a + b)
	return 1 // number of return values
}

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	L.Register("adder", adder)
	L.DoString("print(adder(2, 2))")
}
```

To enable the FFI functionality
```go
func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()
	
	// this one is needed to enable the FFI module
	L.OpenFFI()

	// more code comes here...
}
```

ON ERROR HANDLING
---------------------

Lua's exceptions are incompatible with Go, golua works around this incompatibility by setting up protected execution environments in `lua.State.DoString`, `lua.State.DoFile`  and lua.State.Call and turning every exception into a Go panic.

This means that:
1. In general you can't do any exception handling from Lua, `pcall` and `xpcall` are renamed to `unsafe_pcall` and `unsafe_xpcall`. They are only safe to be called from Lua code that never calls back to Go. Use at your own risk.
2. The call to lua.State.Error, present in previous versions of this library, has been removed as it is nonsensical
3. Method calls on a newly created `lua.State` happen in an unprotected environment, if Lua throws an exception as a result your program will be terminated. If this is undesirable perform your initialization like this:

```go
func LuaStateInit(L *lua.State) int {
	… initialization goes here…
	return 0
}

…
L.PushGoFunction(LuaStateInit)
err := L.Call(0, 0)
…
```

ON THREADS AND COROUTINES
---------------------

'lua.State' is not thread safe, but the library itself is. Lua's coroutines exist but (to my knowledge) have never been tested and are likely to encounter the same problems that errors have, use at your own peril.

ODDS AND ENDS
---------------------

* Compiling from source yields only a static link library (liblua.a), you can either produce the dynamic link library on your own or use the `luaa` build tag.

ORIGINAL CONTRIBUTORS
---------------------

* Adam Fitzgerald (original author)
* Alessandro Arzilli
* Steve Donovan
* Harley Laue
* James Nurmi
* Ruitao
* Xushiwei
* Isaint
* hsinhoyeh
* Viktor Palmkvist
* HongZhen Peng
* Admin36
* Pierre Neidhardt (@Ambrevar)

SEE ALSO
---------------------

- [Luar](https://github.com/stevedonovan/luar/) is a reflection layer on top of golua API providing a simplified way to publish go functions to a Lua VM.
- [Golua unicode](https://bitbucket.org/ambrevar/golua/) is an extension library that adds unicode support to golua and replaces lua regular expressions with re2.

Licensing
-------------
GoLua is released under the MIT license.
Please see the LICENSE file for more information.

Lua is Copyright (c) Lua.org, PUC-Rio.  All rights reserved.
