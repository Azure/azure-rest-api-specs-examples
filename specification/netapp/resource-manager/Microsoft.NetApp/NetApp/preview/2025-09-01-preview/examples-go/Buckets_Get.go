package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Buckets_Get.json
func ExampleBucketsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBucketsClient().Get(ctx, "myRG", "account1", "pool1", "volume1", "bucket1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetapp.BucketsClientGetResponse{
	// 	Bucket: &armnetapp.Bucket{
	// 		Name: to.Ptr("account1/pool1/volume1/bucket1"),
	// 		Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/buckets"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/buckets/bucket1"),
	// 		Properties: &armnetapp.BucketProperties{
	// 			Path: to.Ptr("/path"),
	// 			FileSystemUser: &armnetapp.FileSystemUser{
	// 				NfsUser: &armnetapp.NfsUser{
	// 					GroupID: to.Ptr[int64](1000),
	// 					UserID: to.Ptr[int64](1001),
	// 				},
	// 			},
	// 			Permissions: to.Ptr(armnetapp.BucketPermissionsReadOnly),
	// 			ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
	// 			Server: &armnetapp.BucketServerProperties{
	// 				CertificateCommonName: to.Ptr("www.example.com"),
	// 				CertificateExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-08-15T13:23:32Z"); return t}()),
	// 				Fqdn: to.Ptr("fullyqualified.domainname.com"),
	// 				IPAddress: to.Ptr("1.2.3.4"),
	// 			},
	// 			Status: to.Ptr(armnetapp.CredentialsStatusCredentialsExpired),
	// 		},
	// 	},
	// }
}
