package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/StorageContainers_ListByResourceGroup.json
func ExampleStorageContainersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageContainersClient().NewListByResourceGroupPager("test-rg", nil)
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
		// page = armazurestackhcivm.StorageContainersClientListByResourceGroupResponse{
		// 	StorageContainerListResult: armazurestackhcivm.StorageContainerListResult{
		// 		Value: []*armazurestackhcivm.StorageContainer{
		// 			{
		// 				Name: to.Ptr("Default_Container"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/storageContainers"),
		// 				ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
		// 					Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/Default_Container"),
		// 				Location: to.Ptr("West US2"),
		// 				Properties: &armazurestackhcivm.StorageContainerProperties{
		// 					Path: to.Ptr("C:\\container_storage"),
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
