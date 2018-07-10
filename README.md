# go-e2e-test

# run

## clone
```
mkdir -p go/src
cd go/src
git clone https://git.dmm.com/zhang-shuguang/go-e2e-test.git
```

## create path
```
export GOPATH="$GOPATH:/path/to/go"
```

## install dependency
```
dep ensure
```

## run test
```
go test -v ./e2e
```
