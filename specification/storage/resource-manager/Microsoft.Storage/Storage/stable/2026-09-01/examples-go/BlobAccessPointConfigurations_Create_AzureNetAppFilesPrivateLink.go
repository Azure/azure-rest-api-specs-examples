package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-09-01/BlobAccessPointConfigurations_Create_AzureNetAppFilesPrivateLink.json
func ExampleBlobAccessPointConfigurationsClient_BeginCreate_createBlobAccessPointConfigurationWithAzureNetAppFilesPrivateLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBlobAccessPointConfigurationsClient().BeginCreate(ctx, "testrg", "teststorageaccount", "anfprivatelinkaccesspoint", armstorage.BlobAccessPointConfiguration{
		Location: to.Ptr("eastus"),
		Properties: &armstorage.BlobAccessPointConfigurationProperties{
			State:       to.Ptr(armstorage.BlobAccessPointConfigurationStateActive),
			Description: to.Ptr("Azure NetApp Files S3-compatible source over Private Link"),
			Source: &armstorage.BlobAccessPointAzureNetAppFilesSourceProperties{
				SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeAzureNetAppFiles),
				Connection: &armstorage.BlobAccessPointPrivateLinkConnectionProperties{
					ConnectionType:      to.Ptr(armstorage.BlobAccessPointConnectionTypePrivateLink),
					PrivateLinkIDType:   to.Ptr(armstorage.BlobAccessPointPrivateLinkIDTypeResourceID),
					PrivateLinkID:       to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/privateLinkServices/anfpls"),
					PrivateLinkGroupID:  to.Ptr("s3"),
					PrivateLinkLocation: to.Ptr("eastus"),
					RequestMessage:      to.Ptr("Blob Access Point connection request for Azure NetApp Files"),
					Endpoint:            to.Ptr("https://anf.contoso.com/bucket1"),
					TLSVerification:     to.Ptr(armstorage.BlobAccessPointTLSVerificationPerform),
				},
				Auth: &armstorage.BlobAccessPointAccessKeyAuthProperties{
					AuthType:        to.Ptr(armstorage.BlobAccessPointRemoteAuthTypeAccessKey),
					AccessKeyID:     to.Ptr("example-access-key"),
					SecretAccessKey: to.Ptr("example-secret-key"),
					SigningRegion:   to.Ptr("us-east-1"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.BlobAccessPointConfigurationsClientCreateResponse{
	// 	BlobAccessPointConfiguration: armstorage.BlobAccessPointConfiguration{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/blobAccessPointConfigurations/anfprivatelinkaccesspoint"),
	// 		Name: to.Ptr("anfprivatelinkaccesspoint"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations"),
	// 		Location: to.Ptr("eastus"),
	// 		SystemData: &armstorage.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 		Properties: &armstorage.BlobAccessPointConfigurationProperties{
	// 			UniqueID: to.Ptr("33333333-4444-5555-6666-777777777777"),
	// 			State: to.Ptr(armstorage.BlobAccessPointConfigurationStateActive),
	// 			Description: to.Ptr("Azure NetApp Files S3-compatible source over Private Link"),
	// 			Source: &armstorage.BlobAccessPointAzureNetAppFilesSourceProperties{
	// 				SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeAzureNetAppFiles),
	// 				Connection: &armstorage.BlobAccessPointPrivateLinkConnectionProperties{
	// 					ConnectionType: to.Ptr(armstorage.BlobAccessPointConnectionTypePrivateLink),
	// 					PrivateLinkIDType: to.Ptr(armstorage.BlobAccessPointPrivateLinkIDTypeResourceID),
	// 					PrivateLinkID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Network/privateLinkServices/anfpls"),
	// 					PrivateLinkGroupID: to.Ptr("s3"),
	// 					PrivateLinkLocation: to.Ptr("eastus"),
	// 					RequestMessage: to.Ptr("Blob Access Point connection request for Azure NetApp Files"),
	// 					Endpoint: to.Ptr("https://anf.contoso.com/bucket1"),
	// 					TLSVerification: to.Ptr(armstorage.BlobAccessPointTLSVerificationPerform),
	// 					PrivateEndpointName: to.Ptr("anfprivateendpoint"),
	// 				},
	// 				Auth: &armstorage.BlobAccessPointAccessKeyAuthProperties{
	// 					AuthType: to.Ptr(armstorage.BlobAccessPointRemoteAuthTypeAccessKey),
	// 					AccessKeyID: to.Ptr("example-access-key"),
	// 					SigningRegion: to.Ptr("us-east-1"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.ResourceProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
