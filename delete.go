package s3

import (
	"context"
	"fmt"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Delete removes the object with the specified key from the S3 bucket.
func (cli *Client) Delete(ctx context.Context, key string) error {
	_, err := cli.sdkClient.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &cli.bucketName,
		Key:    &key,
	})
	if err != nil {
		return NewError(fmt.Errorf("error deleting object from s3: %w", err), log.Data{
			"bucket_name": cli.bucketName,
			"s3_key":      key,
		})
	}
	return nil
}
