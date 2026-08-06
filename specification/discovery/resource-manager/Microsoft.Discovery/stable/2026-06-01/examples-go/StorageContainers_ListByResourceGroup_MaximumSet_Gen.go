package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/StorageContainers_ListByResourceGroup_MaximumSet_Gen.json
func ExampleStorageContainersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageContainersClient().NewListByResourceGroupPager("rgdiscovery", nil)
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
		// page = armdiscovery.StorageContainersClientListByResourceGroupResponse{
		// 	StorageContainerListResult: armdiscovery.StorageContainerListResult{
		// 		Value: []*armdiscovery.StorageContainer{
		// 			{
		// 				ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/StorageContainers/ltlcopua"),
		// 				Name: to.Ptr("ltlcopua"),
		// 				Tags: map[string]*string{
		// 					"key4240": to.Ptr("omppnvnqh"),
		// 				},
		// 				Location: to.Ptr("uksouth"),
		// 				Type: to.Ptr("Microsoft.Discovery/StorageContainers"),
		// 				SystemData: &armdiscovery.SystemData{
		// 					CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 					CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 					LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 				},
		// 				Properties: &armdiscovery.StorageContainerProperties{
		// 					ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
		// 					StorageStore: &armdiscovery.AzureStorageBlobStore{
		// 						Kind: to.Ptr(armdiscovery.StorageStoreTypeAzureStorageBlob),
		// 						StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/storageaccount"),
		// 						MountProtocol: to.Ptr(armdiscovery.BlobStorageMountProtocolNFS),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/azv"),
		// 	},
		// }
	}
}
