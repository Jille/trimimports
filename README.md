# trimimports

[![Go Reference](https://pkg.go.dev/badge/github.com/Jille/trimimports.svg)](https://pkg.go.dev/github.com/Jille/trimimports)

This library is intended for use with generated code. You generate all possibly needed imports, and then Trim it to remove those your generated code didn't end up using.

I've found that manually writing logic for when to import a package is error prone and often breaks when the code changes. This library simplifies it to just emitting all of the possible imports, and trimming them to what was used afterwards.
