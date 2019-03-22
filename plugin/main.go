package main

import (
	"fmt"
	s3ds "github.com/ipfs-s3c-storj-plugin"
	"gx/ipfs/QmPDEJTb3WBHmvubsLXCaqRPC8dRgvFz7A4p96dxZbJuWL/go-ipfs/plugin"
	"gx/ipfs/QmPDEJTb3WBHmvubsLXCaqRPC8dRgvFz7A4p96dxZbJuWL/go-ipfs/repo"
	"gx/ipfs/QmPDEJTb3WBHmvubsLXCaqRPC8dRgvFz7A4p96dxZbJuWL/go-ipfs/repo/fsrepo"
	config "gx/ipfs/QmUAuYuiafnJRZxDDX7MuruMNsicYNuyub5vUeAcupUBNs/go-ipfs-config"
)

var Plugins = []plugin.Plugin{
	&S3Plugin{},
}

type S3Plugin struct{}

func (s3p S3Plugin) Name() string {
	return "s3-datastore-plugin"
}

func (s3p S3Plugin) Version() string {
	return "0.0.2"
}

func (s3p S3Plugin) Init() error {
	config.Profiles["s3c-storjds"] = profile
	return nil
}

var DatastoreType = "s3ds"

func (s3p S3Plugin) DatastoreTypeName() string {
	return DatastoreType
}

func DsConfigHandler(m map[string]interface{}) (fsrepo.DatastoreConfig, error) {
	region, ok := m["region"].(string)
	if !ok {
		return nil, fmt.Errorf("s3ds: no region specified")
	}

	bucket, ok := m["bucket"].(string)
	if !ok {
		return nil, fmt.Errorf("s3ds: no bucket specified")
	}

	accessKey, ok := m["accessKey"].(string)
	if !ok {
		return nil, fmt.Errorf("s3ds: no accessKey specified")
	}

	secretKey, ok := m["secretKey"].(string)
	if !ok {
		return nil, fmt.Errorf("s3ds: no secretKey specified")
	}

	endpoint, ok := m["endpoint"].(string)
	if !ok {
		return nil, fmt.Errorf("ds-storj: unable to convert endpoint to string type")
	}
	if endpoint == "" {
		return nil, fmt.Errorf("ds-storj: endpoint configuration is empty")
	}

	var rootDirectory string
	if v, ok := m["rootDirectory"]; ok {
		rootDirectory, ok = v.(string)
		if !ok {
			return nil, fmt.Errorf("s3ds: rootDirectory not a string")
		}
	}

	var workers int
	if v, ok := m["workers"]; ok {
		workersf, ok := v.(float64)
		workers = int(workersf)
		switch {
		case !ok:
			return nil, fmt.Errorf("s3ds: workers not a number")
		case workers <= 0:
			return nil, fmt.Errorf("s3ds: workers <= 0: %f", workersf)
		case float64(workers) != workersf:
			return nil, fmt.Errorf("s3ds: workers is not an integer: %f", workersf)
		}
	}

	return &S3Config{
		cfg: s3ds.Config{
			Region:    region,
			Bucket:    bucket,
			AccessKey: accessKey,
			SecretKey: secretKey,
			Endpoint:  endpoint,
			//	SessionToken:   sessionToken,
			RootDirectory: rootDirectory,
			Workers:       workers,
			//	RegionEndpoint: endpoint,
		},
	}, nil
}
func (s3p S3Plugin) DatastoreConfigParser() fsrepo.ConfigFromMap {
	return DsConfigHandler
}

type S3Config struct {
	cfg s3ds.Config
}

func (s3c *S3Config) DiskSpec() fsrepo.DiskSpec {
	return fsrepo.DiskSpec{
		"bucket":        s3c.cfg.Bucket,
		"region":        s3c.cfg.Region,
		"endpoint":      s3c.cfg.Endpoint,
		"rootDirectory": s3c.cfg.RootDirectory,
	}
}

func (s3c *S3Config) Create(path string) (repo.Datastore, error) {
	d, err := s3ds.NewS3Datastore(s3c.cfg)
	if err != nil {
		return nil, err
	}
	if err := d.BucketExists(s3c.cfg.Bucket); err != nil {
		if err := d.CreateBucket(s3c.cfg.Bucket); err != nil {
			return nil, err
		}
	}
	return d, nil
}
