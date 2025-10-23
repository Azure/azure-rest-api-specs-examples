package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-07-01-preview/Buckets_GenerateCredentials.json
func ExampleBucketsClient_GenerateCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBucketsClient().GenerateCredentials(ctx, "myRG", "account1", "pool1", "volume1", "bucket1", armnetapp.BucketCredentialsExpiry{
		KeyPairExpiryDays: to.Ptr[int32](3),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.BucketsClientGenerateCredentialsResponse{
	// 	BucketGenerateCredentials: &armnetapp.BucketGenerateCredentials{
	// 		AccessKey: to.Ptr("<REDACTED>"),
	// 		KeyPairExpiry: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-08-15T13:23:33Z"); return t}()),
	// 		SecretKey: to.Ptr("<REDACTED>"),
	// 	},
	// }
}
