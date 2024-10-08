package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
)

// Generated from example definition: 2024-03-01/StandbyVirtualMachinePoolRuntimeViews_Get.json
func ExampleStandbyVirtualMachinePoolRuntimeViewsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstandbypool.NewClientFactory("00000000-0000-0000-0000-000000000009", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStandbyVirtualMachinePoolRuntimeViewsClient().Get(ctx, "rgstandbypool", "pool", "latest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstandbypool.StandbyVirtualMachinePoolRuntimeViewsClientGetResponse{
	// 	StandbyVirtualMachinePoolRuntimeViewResource: &armstandbypool.StandbyVirtualMachinePoolRuntimeViewResource{
	// 		Properties: &armstandbypool.StandbyVirtualMachinePoolRuntimeViewResourceProperties{
	// 			InstanceCountSummary: []*armstandbypool.VirtualMachineInstanceCountSummary{
	// 				{
	// 					Zone: to.Ptr[int64](1),
	// 					InstanceCountsByState: []*armstandbypool.PoolResourceStateCount{
	// 						{
	// 							State: to.Ptr("creating"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("running"),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocating"),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocated"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("starting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr("deleting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Zone: to.Ptr[int64](2),
	// 					InstanceCountsByState: []*armstandbypool.PoolResourceStateCount{
	// 						{
	// 							State: to.Ptr("creating"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("running"),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocating"),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocated"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("starting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr("deleting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Zone: to.Ptr[int64](3),
	// 					InstanceCountsByState: []*armstandbypool.PoolResourceStateCount{
	// 						{
	// 							State: to.Ptr("creating"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("running"),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocating"),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr("deallocated"),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr("starting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr("deleting"),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000009/resourceGroups/rgstandbypool/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/pool/runtimeViews/latest"),
	// 		Name: to.Ptr("pool"),
	// 		Type: to.Ptr("Microsoft.StandbyPool/standbyVirtualMachinePools/runtimeViews"),
	// 		SystemData: &armstandbypool.SystemData{
	// 			CreatedBy: to.Ptr("pooluser@microsoft.com"),
	// 			CreatedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-14T23:31:59.679Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pooluser@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armstandbypool.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-14T23:31:59.679Z"); return t}()),
	// 		},
	// 	},
	// }
}
