# Test

## Install local go dependencies

```
cd ..
git clone git@github.com:iden3/go-iden3-core.git
git clone git@github.com:iden3/go-iden3-crypto.git
```

## Gomobile build

```
go mod vendor
gomobile init
GO111MODULE=off go get github.com/ethereum/go-ethereum
ln -s $PWD ~/go/src/
cp -r \
  "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
  "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
# GO111MODULE=off gomobile bind -target=android test 
GOFLAGS=-mod=vendor gomobile bind -target=android test
```
