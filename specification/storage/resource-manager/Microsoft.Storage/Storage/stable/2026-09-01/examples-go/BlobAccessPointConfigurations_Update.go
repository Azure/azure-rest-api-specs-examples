package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-09-01/BlobAccessPointConfigurations_Update.json
func ExampleBlobAccessPointConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBlobAccessPointConfigurationsClient().BeginUpdate(ctx, "testrg", "teststorageaccount", "testaccesspointconfig", armstorage.BlobAccessPointConfigurationUpdate{
		Properties: &armstorage.BlobAccessPointConfigurationPropertiesUpdate{
			State:       to.Ptr(armstorage.BlobAccessPointConfigurationStateInactive),
			Description: to.Ptr("Updated NetApp ONTAP S3-compatible source"),
			Source: &armstorage.BlobAccessPointNetAppONTAPSourcePropertiesUpdate{
				SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeNetAppONTAP),
				Auth: &armstorage.BlobAccessPointAccessKeyAuthPropertiesUpdate{
					AuthType:        to.Ptr(armstorage.BlobAccessPointRemoteAuthTypeAccessKey),
					SecretAccessKey: to.Ptr("rotated-secret-key"),
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
	// res = armstorage.BlobAccessPointConfigurationsClientUpdateResponse{
	// 	BlobAccessPointConfiguration: armstorage.BlobAccessPointConfiguration{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/blobAccessPointConfigurations/testaccesspointconfig"),
	// 		Name: to.Ptr("testaccesspointconfig"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations"),
	// 		Location: to.Ptr("eastus"),
	// 		SystemData: &armstorage.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)),
	// 		},
	// 		Properties: &armstorage.BlobAccessPointConfigurationProperties{
	// 			UniqueID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 			State: to.Ptr(armstorage.BlobAccessPointConfigurationStateInactive),
	// 			Description: to.Ptr("Updated NetApp ONTAP S3-compatible source"),
	// 			Source: &armstorage.BlobAccessPointNetAppONTAPSourceProperties{
	// 				SourceType: to.Ptr(armstorage.BlobAccessPointSourceTypeNetAppONTAP),
	// 				Connection: &armstorage.BlobAccessPointEndpointConnectionProperties{
	// 					ConnectionType: to.Ptr(armstorage.BlobAccessPointConnectionTypeEndpoint),
	// 					Endpoint: to.Ptr("https://netapp.contoso.com/bucket1"),
	// 					TLSVerification: to.Ptr(armstorage.BlobAccessPointTLSVerificationPerform),
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
