package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-07-01/Endpoints_CreateOrUpdate_AzureStorageNfsFileShare.json
func ExampleEndpointsClient_CreateOrUpdate_endpointsCreateOrUpdateAzureStorageNfsFileShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", "examples-endpointName", armstoragemover.Endpoint{
		Properties: &armstoragemover.AzureStorageNfsFileShareEndpointProperties{
			EndpointType:             to.Ptr(armstoragemover.EndpointTypeAzureStorageNfsFileShare),
			StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
			FileShareName:            to.Ptr("examples-fileshare"),
			Description:              to.Ptr("Example Storage File Share Endpoint Description"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragemover.EndpointsClientCreateOrUpdateResponse{
	// 	Endpoint: &armstoragemover.Endpoint{
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName"),
	// 		Name: to.Ptr("examples-endpointName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
	// 		Properties: &armstoragemover.AzureStorageNfsFileShareEndpointProperties{
	// 			EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageNfsFileShare),
	// 			StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
	// 			FileShareName: to.Ptr("examples-fileshare"),
	// 			Description: to.Ptr("Example Storage File Share Endpoint Description"),
	// 			ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
