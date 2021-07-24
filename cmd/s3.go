package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3CreateBucketAPI interface {
	CreateBucket(ctx context.Context,
		params *s3.CreateBucketInput,
		optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
}

func createS3Bucket(c context.Context, api S3CreateBucketAPI, bucketName string, region string) (*s3.CreateBucketOutput, error) {
	in := &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	}
	return api.CreateBucket(c, in)
}

type S3PutPublicAccessBlockAPI interface {
	PutPublicAccessBlock(ctx context.Context,
		params *s3.PutPublicAccessBlockInput,
		optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error)
}

func enableAllPublicAccessBlock(c context.Context, api S3PutPublicAccessBlockAPI, bucketName string) (*s3.PutPublicAccessBlockOutput, error) {
	in := &s3.PutPublicAccessBlockInput{
		Bucket: &bucketName,
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       true,
			BlockPublicPolicy:     true,
			IgnorePublicAcls:      true,
			RestrictPublicBuckets: true,
		},
	}
	return api.PutPublicAccessBlock(c, in)
}

type S3PutBucketEncryptionAPI interface {
	PutBucketEncryption(ctx context.Context,
		params *s3.PutBucketEncryptionInput,
		optFns ...func(*s3.Options)) (*s3.PutBucketEncryptionOutput, error)
}

func enableBucketEncryptionAES256(c context.Context, api S3PutBucketEncryptionAPI, bucketName string) (*s3.PutBucketEncryptionOutput, error) {
	in := &s3.PutBucketEncryptionInput{
		Bucket: &bucketName,
		ServerSideEncryptionConfiguration: &types.ServerSideEncryptionConfiguration{
			Rules: []types.ServerSideEncryptionRule{
				{
					ApplyServerSideEncryptionByDefault: &types.ServerSideEncryptionByDefault{
						SSEAlgorithm: types.ServerSideEncryptionAes256,
					},
				},
			},
		},
	}
	return api.PutBucketEncryption(c, in)
}

type S3PutBucketVersioningAPI interface {
	PutBucketVersioning(ctx context.Context,
		params *s3.PutBucketVersioningInput,
		optFns ...func(*s3.Options)) (*s3.PutBucketVersioningOutput, error)
}

func enableBucketVersioning(c context.Context, api S3PutBucketVersioningAPI, bucketName string) (*s3.PutBucketVersioningOutput, error) {
	in := &s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		VersioningConfiguration: &types.VersioningConfiguration{
			Status: types.BucketVersioningStatusEnabled,
		},
	}
	return api.PutBucketVersioning(c, in)
}
