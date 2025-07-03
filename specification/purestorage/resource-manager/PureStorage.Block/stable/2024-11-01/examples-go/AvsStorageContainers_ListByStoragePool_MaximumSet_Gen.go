package armpurestorageblock_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purestorageblock/armpurestorageblock"
)

// Generated from example definition: 2024-11-01/AvsStorageContainers_ListByStoragePool_MaximumSet_Gen.json
func ExampleAvsStorageContainersClient_NewListByStoragePoolPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurestorageblock.NewClientFactory("BC47D6CC-AA80-4374-86F8-19D94EC70666", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvsStorageContainersClient().NewListByStoragePoolPager("rgpurestorage", "spName", nil)
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
		// page = armpurestorageblock.AvsStorageContainersClientListByStoragePoolResponse{
		// 	AvsStorageContainerListResult: armpurestorageblock.AvsStorageContainerListResult{
		// 		Value: []*armpurestorageblock.AvsStorageContainer{
		// 			{
		// 				Properties: &armpurestorageblock.AvsStorageContainerProperties{
		// 					Space: &armpurestorageblock.Space{
		// 						TotalUsed: to.Ptr[int64](28),
		// 						Unique: to.Ptr[int64](4),
		// 						Snapshots: to.Ptr[int64](5),
		// 						Shared: to.Ptr[int64](9),
		// 					},
		// 					ResourceName: to.Ptr("name"),
		// 					ProvisionedLimit: to.Ptr[int64](24),
		// 					Datastore: to.Ptr("name"),
		// 					Mounted: to.Ptr(true),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("acezain"),
		// 				Type: to.Ptr("rhdmvrbnpxcydalkvtfsb"),
		// 				SystemData: &armpurestorageblock.SystemData{
		// 					CreatedBy: to.Ptr("ruoitchmuomrbscg"),
		// 					CreatedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.341Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("thfyhokbrldzmghuylqbwpbublj"),
		// 					LastModifiedByType: to.Ptr(armpurestorageblock.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-04T05:29:25.345Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
