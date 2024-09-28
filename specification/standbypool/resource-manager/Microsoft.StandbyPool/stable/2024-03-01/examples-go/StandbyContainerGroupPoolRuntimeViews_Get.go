package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: 2024-03-01/StandbyContainerGroupPoolRuntimeViews_Get.json
func ExampleStandbyContainerGroupPoolRuntimeViewsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyContainerGroupPoolRuntimeViewsClient().Get(ctx, "rgstandbypool", "pool", "latest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstandbypool.StandbyContainerGroupPoolRuntimeViewsClientGetResponse{
	// 	StandbyContainerGroupPoolRuntimeViewResource: &armstandbypool.StandbyContainerGroupPoolRuntimeViewResource{
	// 		Properties: &armstandbypool.StandbyContainerGroupPoolRuntimeViewResourceProperties{
	// 			InstanceCountSummary: []*armstandbypool.ContainerGroupInstanceCountSummary{
	// 				{
	// 					InstanceCountsByState: []*armstandbypool.PoolResourceStateCount{
	// 						{
	// 							State: to.Ptr("creating"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("running"),
	// 							Count: to.Ptr[int64](500),
	// 						},
	// 						{
	// 							State: to.Ptr("deleting"),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyContainerGroupPools/pool/runtimeViews/latest"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyContainerGroupPools/runtimeViews"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T16:33:22.210Z"); return t}()),
	// 		},
	// 	},
	// }
}
