# Goshim

A Go dependency management system inspired by Bundler and rbenv.

```go
# Download and install goshim
go get github.com/kevin-cantwell/goshim

# Updates the source file you specify and modifies your path by inserting a 'go' shim.
# Also installs a ~/.goshim directory that contains project-specific
# Maybe 'create' isn't the best command
goshim create ~/.profile
```

### goshim commands

`install [source_file]`:
1. Installs the shim to `$HOME/.goshim/go` if it does not already exist.
1. If needed, adds the shim to your path by modifying `source_file` with: `export PATH=$HOME/.goshim/go:$PATH`

`init`:
1. Creates a `.goget` file in the current directory if one does not already exist.
1. Executes

`get`:
1. Identical in usage to `go get` except that it accepts 

`exec [cmd]`:
1. Executes the command in a goshim environment. Useful if some other tool requires $GOPATH to be set, but does not call any go executables

`export`:
1. Exports all environment variables known to go env
