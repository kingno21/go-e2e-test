# go-e2e-test

# run

## dependency
```
# install gvm
https://github.com/moovweb/gvm

brew install dep

gvm install 1.10.3
gvm use 1.10.3 --default

dep ensure
```
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
