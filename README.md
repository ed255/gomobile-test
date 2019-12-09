# Test

## Gomobile build

```
go get -u github.com/ed255/gomobile-test
cd $GOPATH/src/github.com/ed255/gomobile-test
export GO111MODULE=on
go mod download
go mod vendor
export GO111MODULE=off

gomobile bind -target=android
cp -r \
  "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
  "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
gomobile bind -target=android
```
