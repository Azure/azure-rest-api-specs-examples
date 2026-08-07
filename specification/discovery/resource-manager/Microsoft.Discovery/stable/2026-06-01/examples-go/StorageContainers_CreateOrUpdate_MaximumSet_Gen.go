package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/StorageContainers_CreateOrUpdate_MaximumSet_Gen.json
func ExampleStorageContainersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageContainersClient().BeginCreateOrUpdate(ctx, "rgdiscovery", "49af599cddb38a473b", armdiscovery.StorageContainer{
		Properties: &armdiscovery.StorageContainerProperties{
			StorageStore: &armdiscovery.AzureStorageBlobStore{
				Kind:             to.Ptr(armdiscovery.StorageStoreTypeAzureStorageBlob),
				StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/storageaccount"),
				MountProtocol:    to.Ptr(armdiscovery.BlobStorageMountProtocolNFS),
			},
		},
		Tags: map[string]*string{
			"key4240": to.Ptr("omppnvnqh"),
		},
		Location: to.Ptr("uksouth"),
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
	// res = armdiscovery.StorageContainersClientCreateOrUpdateResponse{
	// 	StorageContainer: armdiscovery.StorageContainer{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/storageContainers/49af599cddb38a473b"),
	// 		Name: to.Ptr("49af599cddb38a473b"),
	// 		Tags: map[string]*string{
	// 			"key4240": to.Ptr("omppnvnqh"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/storageContainers"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.StorageContainerProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			StorageStore: &armdiscovery.AzureStorageBlobStore{
	// 				Kind: to.Ptr(armdiscovery.StorageStoreTypeAzureStorageBlob),
	// 				StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Storage/storageAccounts/storageaccount"),
	// 				MountProtocol: to.Ptr(armdiscovery.BlobStorageMountProtocolNFS),
	// 			},
	// 		},
	// 	},
	// }
}
