IPFS_PATH ?= ${HOME}/.ipfs
IPFSVERSION = QmPDEJTb3WBHmvubsLXCaqRPC8dRgvFz7A4p96dxZbJuWL
IPFSCMDBUILDPATH=vendor/gx/ipfs/$(IPFSVERSION)/go-ipfs/cmd/ipfs
REPOROOT=$(shell pwd)
deps:
	gx --verbose install --local

build: deps

	rm -rf $(REPOROOT)/build
	mkdir $(REPOROOT)/build
	(go build  -o=build/s3c-storj-plugin.so  -buildmode=plugin ./plugin ;  chmod a+x build/s3c-storj-plugin.so)
	(cd $(IPFSCMDBUILDPATH) ; go build ; cp ipfs $(REPOROOT)/build)

install: build
	mkdir -p ${IPFS_PATH}/plugins
	rm -f ${IPFS_PATH}/plugins/s3c-storj-plugin.so
	cp build/s3c-storj-plugin.so ${IPFS_PATH}/plugins/
