package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	ds "gx/ipfs/QmUadX5EcvrBmxAV9sE7wUWtWSqxns5K84qKJBixmcT1w9/go-datastore"
)

// BucketExists is used to lookup if the designated bucket exists
func (s *S3Bucket) BucketExists(name string) error {
	listParam := &s3.ListBucketsInput{}
	out, err := s.S3.ListBuckets(listParam)
	if err != nil {
		return parseError(err)
	}
	for _, v := range out.Buckets {
		if *v.Name == name {
			return nil
		}
	}
	return ds.ErrNotFound
}

// CreateBucket is used to create a bucket
func (s *S3Bucket) CreateBucket(name string) error {
	createParam := &s3.CreateBucketInput{
		Bucket: aws.String(name),
	}
	// create bucket ensure we have initialize client correct
	_, err := s.S3.CreateBucket(createParam)
	return parseError(err)
}
