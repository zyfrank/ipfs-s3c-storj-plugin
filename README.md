# ipfs-s3c-storj-plugin
This experimental ipfs plugin combine code from https://github.com/RTradeLtd/storj-ipfs-ds-plugin and https://github.com/ipfs/go-ds-s3.

# Usage 

（make sure backup ~/.ipfs/, then delete this dir）

clone this repo to your $GOPATH/src/github.com 

cd $GOPATH/src/github.com/ipfs-s3c-storj-plugin 

make install

./build/ipfs init --profile s3c-storjds

change $IPFS_PATH/config (commonly it is ~/.ipfs/config),  input your "accessKey" and "secretKey" which are used to access storj s3 gateway

start your test storj env. by using storj-sim network run

when you use ./build/ipfs add *, the file will be fed to storj


