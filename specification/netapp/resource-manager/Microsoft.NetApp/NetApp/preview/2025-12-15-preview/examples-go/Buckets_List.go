package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v10"
)

// Generated from example definition: 2025-12-15-preview/Buckets_List.json
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
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/buckets/bucket1"),
		// 				Name: to.Ptr("account1/pool1/volume1/bucket1"),
		// 				Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/buckets"),
		// 				Properties: &armnetapp.BucketProperties{
		// 					Path: to.Ptr("/path"),
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 					FileSystemUser: &armnetapp.FileSystemUser{
		// 						NfsUser: &armnetapp.NfsUser{
		// 							UserID: to.Ptr[int64](1001),
		// 							GroupID: to.Ptr[int64](1000),
		// 						},
		// 					},
		// 					Status: to.Ptr(armnetapp.CredentialsStatusActive),
		// 					Server: &armnetapp.BucketServerProperties{
		// 						Fqdn: to.Ptr("fullyqualified.domainname.com"),
		// 						CertificateCommonName: to.Ptr("www.example.com"),
		// 						CertificateExpiryDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-08-15T13:23:32Z"); return t}()),
		// 						IPAddress: to.Ptr("1.2.3.4"),
		// 					},
		// 					AkvDetails: &armnetapp.AzureKeyVaultDetails{
		// 						CertificateAkvDetails: &armnetapp.CertificateAkvDetails{
		// 							CertificateKeyVaultURI: to.Ptr("https://REDACTED.vault.azure.net/"),
		// 							CertificateName: to.Ptr("my-certificate"),
		// 						},
		// 						CredentialsAkvDetails: &armnetapp.CredentialsAkvDetails{
		// 							CredentialsKeyVaultURI: to.Ptr("https://REDACTED.vault.azure.net/"),
		// 							SecretName: to.Ptr("my-secret"),
		// 						},
		// 					},
		// 					Permissions: to.Ptr(armnetapp.BucketPermissionsReadOnly),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
