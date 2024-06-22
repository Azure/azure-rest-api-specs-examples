package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ee6d9fd7687d4b67117c5a167c191a7e7e70b53/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/StorageMovers_CreateOrUpdate.json
func ExampleStorageMoversClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageMoversClient().CreateOrUpdate(ctx, "examples-rg", "examples-storageMoverName", armstoragemover.StorageMover{
		Location: to.Ptr("eastus2"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
		Properties: &armstoragemover.Properties{
			Description: to.Ptr("Example Storage Mover Description"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageMover = armstoragemover.StorageMover{
	// 	Name: to.Ptr("examples-storageMoverName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers"),
	// 	ID: to.Ptr("/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName"),
	// 	SystemData: &armstoragemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-01T02:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armstoragemover.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armstoragemover.Properties{
	// 		Description: to.Ptr("Example Storage Mover Description"),
	// 	},
	// }
}
