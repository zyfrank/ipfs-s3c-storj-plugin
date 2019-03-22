package main

import (
	config "gx/ipfs/QmUAuYuiafnJRZxDDX7MuruMNsicYNuyub5vUeAcupUBNs/go-ipfs-config"
	"os"
)

var profile config.Profile = config.Profile{
	Description: "Replaces default datastore config with experimental storj",
	Transform: func(c *config.Config) error {
		c.Datastore.Spec = map[string]interface{}{
			"type": "mount",
			"mounts": []interface{}{
				map[string]interface{}{
					"mountpoint": "/blocks",
					"name":       DatastoreType,
					"type":       "log",
					"child": map[string]interface{}{
						"accessKey":     os.Getenv("STORJ_ACCESS_KEY"),
						"secretKey":     os.Getenv("STORJ_SECRET_KEY"),
						"bucket":        "ipfs",
						"region":        "us-east-1",
						"endpoint":      "http://127.0.0.1:9000",
						"rootDirectory": "",
						"type":          DatastoreType,
						"logPath":       "",
					},
				},
				map[string]interface{}{
					"child": map[string]interface{}{
						"compression": "none",
						"path":        "datastore",
						"type":        "levelds",
					},
					"mountpoint": "/",
					"prefix":     "leveldb.datastore",
					"type":       "measure",
				},
			},
		}
		return nil
	},
}
