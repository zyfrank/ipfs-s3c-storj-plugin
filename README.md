# ipfs-s3c-storj-plugin
This experimental ipfs plugin combine code from https://github.com/RTradeLtd/storj-ipfs-ds-plugin and https://github.com/ipfs/go-ds-s3.

# Usage 

clone this repo to your $GOPATH/src/github.com 

cd $GOPATH/src/github.com/ipfs-s3c-storj-plugin 

gx install --local 
(here I personnally gx publish 3 ipfs packages:go-ipfs-config,go-ipfs and iptb.  maybe in your site it is hard to get these three packages
just check https://github.com/zyfrank/go-ipfs, https://github.com/zyfrank/go-ipfs-config and https://github.com/zyfrank/iptb, clone to your local env. ,switch to branch storj-s3c-plugin, then use gx publish -f to your local IPFS node, so gx install --local can find these three packages)

make install

./build/ipfs init --profile s3-storjds

change $IPFS_PATH/config (commonly it is ~/.ipfs/config),  input your "accessKey" and "secretKey" which are used to access storj s3 gateway

start your test storj env. by using storj-sim network run

now start ipfs daemon  ./build/ipfs daemon

when you ./build/ipfs add *, the file will be feed to storj


