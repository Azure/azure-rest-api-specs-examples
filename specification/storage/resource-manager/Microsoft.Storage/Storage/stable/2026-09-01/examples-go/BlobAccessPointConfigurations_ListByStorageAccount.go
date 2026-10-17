package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-09-01/BlobAccessPointConfigurations_ListByStorageAccount.json
func ExampleBlobAccessPointConfigurationsClient_NewListByStorageAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBlobAccessPointConfigurationsClient().NewListByStorageAccountPager("testrg", "teststorageaccount", nil)
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
		// page = armstorage.BlobAccessPointConfigurationsClientListByStorageAccountResponse{
		// 	BlobAccessPointConfigurationListResult: armstorage.BlobAccessPointConfigurationListResult{
		// 		Value: []*armstorage.BlobAccessPointConfiguration{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/blobAccessPointConfigurations/testaccesspointconfig"),
		// 				Name: to.Ptr("testaccesspointconfig"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations"),
		// 				Location: to.Ptr("eastus"),
		// 				SystemData: &armstorage.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 				Properties: &armstorage.BlobAccessPointConfigurationProperties{
		// 					UniqueID: to.Ptr("11111111-2222-3333-4444-555555555555"),
		// 					State: to.Ptr(armstorage.BlobAccessPointConfigurationStateActive),
		// 					Description: to.Ptr("NetApp ONTAP S3-compatible source"),
		// 					Source: &armstorage.BlobAccessPointNetAppONTAPSourceProperties{
		// 						SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeNetAppONTAP),
		// 						Connection: &armstorage.BlobAccessPointEndpointConnectionProperties{
		// 							ConnectionType: to.Ptr(armstorage.BlobAccessPointConnectionTypeEndpoint),
		// 							Endpoint: to.Ptr("https://netapp.contoso.com/bucket1"),
		// 							TLSVerification: to.Ptr(armstorage.BlobAccessPointTLSVerificationPerform),
		// 						},
		// 						Auth: &armstorage.BlobAccessPointAccessKeyAuthProperties{
		// 							AuthType: to.Ptr(armstorage.BlobAccessPointRemoteAuthTypeAccessKey),
		// 							AccessKeyID: to.Ptr("example-access-key"),
		// 							SigningRegion: to.Ptr("us-east-1"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/blobAccessPointConfigurations/privatelinkaccesspointconfig"),
		// 				Name: to.Ptr("privatelinkaccesspointconfig"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations"),
		// 				Location: to.Ptr("eastus"),
		// 				SystemData: &armstorage.SystemData{
		// 					CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)),
		// 				},
		// 				Properties: &armstorage.BlobAccessPointConfigurationProperties{
		// 					UniqueID: to.Ptr("22222222-3333-4444-5555-666666666666"),
		// 					State: to.Ptr(armstorage.BlobAccessPointConfigurationStateActive),
		// 					Description: to.Ptr("Private-link-backed S3-compatible source"),
		// 					Source: &armstorage.BlobAccessPointDellOneFSSourceProperties{
		// 						SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeDellOneFS),
		// 						Connection: &armstorage.BlobAccessPointPrivateLinkConnectionProperties{
		// 							ConnectionType: to.Ptr(armstorage.BlobAccessPointConnectionTypePrivateLink),
		// 							PrivateLinkIDType: to.Ptr(armstorage.BlobAccessPointPrivateLinkIDTypeResourceID),
		// 							PrivateLinkID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/privateLinkServices/testpls"),
		// 							PrivateLinkGroupID: to.Ptr("s3"),
		// 							PrivateLinkLocation: to.Ptr("eastus"),
		// 							RequestMessage: to.Ptr("Blob Access Point connection request"),
		// 							Endpoint: to.Ptr("https://onefs.contoso.com/bucket2"),
		// 							TLSVerification: to.Ptr(armstorage.BlobAccessPointTLSVerificationPerform),
		// 							PrivateEndpointName: to.Ptr("testpe"),
		// 						},
		// 						Auth: &armstorage.BlobAccessPointAccessKeyAuthProperties{
		// 							AuthType: to.Ptr(armstorage.BlobAccessPointRemoteAuthTypeAccessKey),
		// 							AccessKeyID: to.Ptr("example-access-key-2"),
		// 							SigningRegion: to.Ptr("us-east-1"),
		// 							HostOverride: to.Ptr("onefs.contoso.com"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.ResourceProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
