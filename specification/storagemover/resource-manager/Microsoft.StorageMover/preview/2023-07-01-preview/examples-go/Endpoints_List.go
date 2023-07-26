package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_List.json
func ExampleEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEndpointsClient().NewListPager("examples-rg", "examples-storageMoverName", nil)
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
		// page.EndpointList = armstoragemover.EndpointList{
		// 	Value: []*armstoragemover.Endpoint{
		// 		{
		// 			Name: to.Ptr("examples-endpointName1"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName1"),
		// 			Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
		// 				Description: to.Ptr("Example Storage Container Endpoint 1 Description"),
		// 				EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
		// 				ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 				BlobContainerName: to.Ptr("examples-blobcontainer1"),
		// 				StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-endpointName2"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName2"),
		// 			Properties: &armstoragemover.NfsMountEndpointProperties{
		// 				Description: to.Ptr("Example Storage Container Endpoint 2 Description"),
		// 				EndpointType: to.Ptr(armstoragemover.EndpointTypeNfsMount),
		// 				ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 				Export: to.Ptr("/"),
		// 				Host: to.Ptr("0.0.0.0"),
		// 				NfsVersion: to.Ptr(armstoragemover.NfsVersionNFSv4),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examples-endpointName3"),
		// 			Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 			ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName3"),
		// 			Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
		// 				Description: to.Ptr("Example Storage Container Endpoint 3 Description"),
		// 				EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
		// 				ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 				BlobContainerName: to.Ptr("examples-blobcontainer3"),
		// 				StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
		// 			},
		// 	}},
		// }
	}
}
