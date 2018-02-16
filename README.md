Developed with:

```
go version go1.9 windows/amd64
```

# Installation
1. [Install Go](https://golang.org/doc/install) if not already installed.

1. Navigate to the desired `cmd` subdirectory (`./cmd/npic` or `./cmd/recordc`).

1. Create the executable.
	* `go install` will create the executable in `$GOBIN` (`$GOPATH/bin`), making it accessible via `$PATH` (`npic`|`recordc`).
	* `go build` will create the executable in the directory in which it is called, requiring local execution (`./npic`|`./recordc`)

# Part 1
`npic` takes a single parameter, `-file`, that should point to a newline-delimited list of NPIs to be tested. For convenience, a file was created using the examples from [data_eng/README.md]() at [data_eng/npis.txt]().

Example:
```sh
npic -file ../../data_eng/npis.txt
```

# Part 2
`recordc` has one required parameter and one optional parameter.

`-file` is required and should point to a delimited (pipe, comma, or space) file of records. Each record must be on its own line (newline-delimited).

`-sort` takes one of three parameters: `provider`, `dob`, or `lastname`.
* `provider`:	Sorts records by provider type ascending with a subsort by last name ascending.
* `dob`: Sorts records by date of birth ascending.
* `lastname`: Sorts records by last name descending.

Examples:
```sh
recordc -file ../../data_eng/pipe_delimited.txt
```

```sh
recordc -file ../../data_eng/comma_delimited.txt -sort provider
```
