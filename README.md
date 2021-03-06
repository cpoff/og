Og-Lang *v0.7.2* (Optimistic Golang)
===

### [Documentation Website](https://champii.github.io/og)
### [GoDoc](https://godoc.org/github.com/Champii/og)

We are looking for a better name !  
[Renaming Og](https://github.com/Champii/og/issues/7)

# Disclamer

<h3> This software is in its early stage.</h3>

# Quick overview

This language aim to save developer time as much as possible by implementing syntactic sugar like templates or auto return. It might borrow some concepts from functionnal languages like function composition or currying.

This is a quick and dirty syntactic coloration made by hand in optic to highlight main features of the language, please don't judge me  (Or better, please contribute to make a better one :] )

![Overview](https://github.com/Champii/og/raw/master/docs/_media/overview_color.png)

## Demo

<p float="left">
  <img height="220" src="https://github.com/Champii/og/raw/master/docs/_media/hello_preview.gif" />
  <img height="220" src="https://github.com/Champii/og/raw/master/docs/_media/og_preview.gif" />
</p>

# Build

Here is the procedure to regenerate the parser from the grammar if you want to make changes to it.  
If you just want to (re)build the binary, you can call `make build` or just `go build` (needs a previously generated parser from grammar. See below)  
You will need `Java`, the Antlr4 library is in `./parser/antlr4-4.7.1-SNAPSHOT-complete.jar`

```bash
# Get Og
go get -u github.com/champii/og 
cd $GOPATH/src/github.com/champii/og

# This will regenerate the grammar,
# Compile the existing sources from the previous Og (`og lib`)
# And run the tests.
# Needs the last official `og` binary version at global scope.
make

# It cleans the `lib` folder,
# Then compiles og from the previous global version (`og lib`)
# Then recomiles it from itself (`./og lib`)
# And run the tests
make re

# Simple exemple
og exemples/import.og
```

The current build time of the project is around 5s for all sources files with `./og` alone, and around 20s for full rebootstrap with `make re` (That bootstraps from old version then rebootstraps from itself, with `go build` and `go test` each time). 

# Quotes

<table>
  <tr><td><b>"Golang On Steroids"</b></td>         <td>- <em>Socrates</em></td></tr>
  <tr><td><b>"The Code like it should be. 5/7"</b></td><td>- <em>Mahatma Gandhi</em></td></tr>
  <tr><td><b>"(..Recursive Facepalm..)"</b></td> <td>- <em>Google</em></td></tr>
</table>
