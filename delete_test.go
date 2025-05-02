package s3_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	dps3 "github.com/ONSdigital/dp-s3/v3"
	"github.com/ONSdigital/dp-s3/v3/mock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDelete(t *testing.T) {
	Convey("Given an S3 client configured with a bucket and region", t, func() {
		ctx := context.Background()

		bucket := "myBucket"
		objKey := "my/object/key"
		region := "eu-north-1"

		sdkMock := &mock.S3SDKClientMock{
			DeleteObjectFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
				return &s3.DeleteObjectOutput{}, nil
			},
		}

		cli := dps3.InstantiateClient(sdkMock, nil, nil, nil, bucket, region, aws.Config{})

		Convey("When Delete is called with a valid key", func() {
			err := cli.Delete(ctx, objKey)

			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the DeleteObject function is called with the correct parameters", func() {
				So(len(sdkMock.DeleteObjectCalls()), ShouldEqual, 1)
				So(sdkMock.DeleteObjectCalls()[0].In, ShouldResemble, &s3.DeleteObjectInput{
					Bucket: &bucket,
					Key:    &objKey,
				})
			})
		})
	})
}
