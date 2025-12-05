package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Buckets_List.json
func ExampleBucketsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBucketsClient().NewListPager("myRG", "account1", "pool1", "volume1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armnetapp.BucketsClientListResponse{
		// 	BucketList: armnetapp.BucketList{
		// 		Value: []*armnetapp.Bucket{
		// 			{
		// 				Name: to.Ptr("account1/pool1/volume1/bucket1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/buckets"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/buckets/bucket1"),
		// 				Properties: &armnetapp.BucketProperties{
		// 					Path: to.Ptr("/path"),
		// 					FileSystemUser: &armnetapp.FileSystemUser{
		// 						NfsUser: &armnetapp.NfsUser{
		// 							GroupID: to.Ptr[int64](1000),
		// 							UserID: to.Ptr[int64](1001),
		// 						},
		// 					},
		// 					Permissions: to.Ptr(armnetapp.BucketPermissionsReadOnly),
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					Server: &armnetapp.BucketServerProperties{
		// 						CertificateCommonName: to.Ptr("www.example.com"),
		// 						CertificateExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-08-15T13:23:32Z"); return t}()),
		// 						Fqdn: to.Ptr("fullyqualified.domainname.com"),
		// 						IPAddress: to.Ptr("1.2.3.4"),
		// 					},
		// 					Status: to.Ptr(armnetapp.CredentialsStatusActive),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
