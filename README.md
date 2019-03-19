# ipfs-s3c-storj-plugin
This experimental ipfs plugin combine code from https://github.com/RTradeLtd/storj-ipfs-ds-plugin and https://github.com/ipfs/go-ds-s3.
Based on these two repo.

#Usage 
1.  clone this repo to your $GOPATH/src/github.com

2.  
cd $GOPATH/src/github.com/ipfs-s3c-storj-plugin
gx install --local
make install
./build/ipfs init --profile s3-storjds

