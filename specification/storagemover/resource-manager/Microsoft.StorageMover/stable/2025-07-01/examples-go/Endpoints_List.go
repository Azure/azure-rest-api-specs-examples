package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: 2025-07-01/Endpoints_List.json
func ExampleEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("60bcfc77-6589-4da2-b7fd-f9ec9322cf95", cred, nil)
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
		// page = armstoragemover.EndpointsClientListResponse{
		// 	EndpointList: armstoragemover.EndpointList{
		// 		NextLink: to.Ptr("https://management.azure.com/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints?$skiptoken=fake-continue-token"),
		// 		Value: []*armstoragemover.Endpoint{
		// 			{
		// 				Name: to.Ptr("examples-endpointName1"),
		// 				Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 				ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName1"),
		// 				Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
		// 					Description: to.Ptr("Example Storage Container Endpoint 1 Description"),
		// 					BlobContainerName: to.Ptr("examples-blobcontainer1"),
		// 					EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
		// 					ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 					StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examples-endpointName2"),
		// 				Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 				ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName2"),
		// 				Properties: &armstoragemover.NfsMountEndpointProperties{
		// 					Description: to.Ptr("Example Storage Container Endpoint 2 Description"),
		// 					EndpointType: to.Ptr(armstoragemover.EndpointTypeNfsMount),
		// 					Export: to.Ptr("/"),
		// 					Host: to.Ptr("0.0.0.0"),
		// 					NfsVersion: to.Ptr(armstoragemover.NfsVersionNFSv4),
		// 					ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("examples-endpointName3"),
		// 				Type: to.Ptr("Microsoft.StorageMover/storageMovers/endpoints"),
		// 				ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/examples-endpointName3"),
		// 				Properties: &armstoragemover.AzureStorageBlobContainerEndpointProperties{
		// 					Description: to.Ptr("Example Storage Container Endpoint 3 Description"),
		// 					BlobContainerName: to.Ptr("examples-blobcontainer3"),
		// 					EndpointType: to.Ptr(armstoragemover.EndpointTypeAzureStorageBlobContainer),
		// 					ProvisioningState: to.Ptr(armstoragemover.ProvisioningStateSucceeded),
		// 					StorageAccountResourceID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
