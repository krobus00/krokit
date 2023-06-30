package kit

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	Credential aws.CredentialsProvider
	Region     string
	Endpoint   string

	PartSize    int64
	Concurrency int
}

type s3Client struct {
	client *s3.Client

	PartSize    int64
	Concurrency int
}

type S3Client interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput) (*manager.UploadOutput, error)
	CreatePresignedURL(ctx context.Context, params *s3.GetObjectInput) (*v4.PresignedHTTPRequest, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
}

func NewS3Client(config *S3Config, optFns ...func(*s3.Options)) (S3Client, error) {
	client := s3.NewFromConfig(aws.Config{
		Credentials: config.Credential,
		Region:      config.Region,
	}, optFns...)

	if client == nil {
		return nil, errors.New("creating an S3 SDK client failed")
	}

	partSize := config.PartSize
	if partSize < manager.MinUploadPartSize {
		partSize = manager.MinUploadPartSize
	}

	concurrency := config.Concurrency
	if concurrency < manager.DefaultUploadConcurrency {
		concurrency = manager.DefaultUploadConcurrency
	}

	return &s3Client{
		client:      client,
		PartSize:    config.PartSize,
		Concurrency: config.Concurrency,
	}, nil
}

func (i *s3Client) PutObject(ctx context.Context, params *s3.PutObjectInput) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(i.client, func(u *manager.Uploader) {
		u.PartSize = i.PartSize
		u.Concurrency = i.Concurrency
	})

	return uploader.Upload(ctx, params)
}

func (i *s3Client) CreatePresignedURL(ctx context.Context, params *s3.GetObjectInput) (*v4.PresignedHTTPRequest, error) {
	return s3.NewPresignClient(i.client).PresignGetObject(ctx, params)
}

func (i *s3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	return i.client.DeleteObjects(ctx, params)
}
