# gocover

A example for go coverage

## Test

Get package:

```
go get github.com/winlinvip/gocover
```

Run the utest `go test ./... -v`:

```
go test ./... -v
#    === RUN   TestMin
#    --- PASS: TestMin (0.00s)
#    PASS
#    ok  	_/Users/winlin/git/gocover/core	0.007s
```

## Cover

First, generate cover files:

```
mkdir -p cover &&
go test github.com/winlinvip/gocover/core -cover -coverprofile cover/core.cover &&
go test github.com/winlinvip/gocover/protocol -cover -coverprofile cover/protocol.cover
```

Then, combile the data:

```
echo 'mode: set' > system.cover &&
tail -q -n +2 cover/*.cover >> system.cover
```

Finally, generate the report:

```
go tool cover -html=system.cover -o system.html
```

Please read the [report](https://winlinvip.github.io/gocover/system.html#file0) here.

## Simple Usage

To get the coverage of a specified package, please run:

```
go test github.com/winlinvip/gocover/protocol -cover -coverprofile out.cover &&
go tool cover -html=out.cover

