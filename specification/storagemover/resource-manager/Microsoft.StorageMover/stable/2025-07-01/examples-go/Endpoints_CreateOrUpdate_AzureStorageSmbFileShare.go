package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-07-01/Endpoints_CreateOrUpdate_AzureStorageSmbFileShare.json
func ExampleEndpointsClient_CreateOrUpdate_endpointsCreateOrUpdateAzureStorageSmbFileShare() {
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
		Properties: &armstoragemover.AzureStorageSmbFileShareEndpointProperties{
			Description:              to.Ptr("Example Storage File Share Endpoint Description"),
			EndpointType:             to.Ptr(armstoragemover.EndpointTypeAzureStorageSmbFileShare),
			FileShareName:            to.Ptr("examples-fileshare"),
			StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
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
	// 		Name: to.Ptr("examples-endpointName"),
	// 		Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
	// 		ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName"),
	// 		Properties: &armstoragemover.AzureStorageSmbFileShareEndpointProperties{
	// 			Description: to.Ptr("Example Storage File Share Endpoint Description"),
	// 			EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageSmbFileShare),
	// 			FileShareName: to.Ptr("examples-fileshare"),
	// 			ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
	// 			StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
	// 		},
	// 	},
	// }
}
