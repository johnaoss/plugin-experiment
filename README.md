# plugin-experiment

This repo is a bunch of experiments I've made w.r.t. having runtime plugins in Go.

# Installation

```bash
go get -u github.com/johnaoss/plugin-experiment
```

# Development

Before developing, you must run `go generate ./...` on the project root to generate 
the required interface definitions. Aside from that, you'll need to be in the `GOPATH` as 
the interpreter `yaegi` doesn't support modules at this time.

# Roadmap

Currently I plan on either figuring out how to modify `genexports` and `yaegi` in order to satisfy the requirements I need.

In the future, I'll try using the offical Go `plugin` package, but that would require me to compile all files given.

# Licensing

This package contains a bunch of code from [containuous/yaegi](https://github.com/containous/yaegi). Primarily under `exports` and `generate`. 

I've modified their `goexports` command from `yaegi` in the `generate/genexports.go` file.