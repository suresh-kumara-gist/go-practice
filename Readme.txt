https://golang.org/doc/code.html#Introduction


$  go env GOPATH
/Users/kumsu25/go


$ export GOPATH=$(go env GOPATH)

$ export GOPATH=$HOME/Sites/devdesktop/go-practice

$ export PATH=$PATH:$(go env GOPATH)/bin
export GOBIN=$GOPATH/bin/


 515  echo $PATH
 516  mkdir -p $GOPATH/src/github.com/user
 517  mkdir $GOPATH/src/github.com/user/hello
 518  go install github.com/user/hello
 519  $GOPATH/bin/hello


create libraries
$ go build github.com/user/stringutil
$ go build


$ go test github.com/user/stringutil


$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!

$ go env
$ export GOBIN=$GOPATH/bin/

import "github.com/golang/example/stringutil"



https://tour.golang.org/basics/17

https://tour.golang.org/list

https://tour.golang.org/flowcontrol/1

https://tour.golang.org/moretypes/1

*** https://golang.org/doc/effective_go.html

https://dev.to/codehakase/how-i-learned-go-programming
https://tour.golang.org/welcome/1

https://golang.org/doc/#articles

https://golang.org/doc/code.html#Introduction



export GOBIN=$GOPATH/bin/
C02VG1JPG8WN:go-practice kumsu25$ go install src/github.com/advanced/advanced.go
C02VG1JPG8WN:go-practice kumsu25$ $GOBIN/advanced

https://tour.golang.org/moretypes/4
