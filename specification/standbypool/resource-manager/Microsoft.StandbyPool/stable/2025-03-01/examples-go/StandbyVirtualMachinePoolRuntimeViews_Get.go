package armstandbypool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool/v2"
)

// Generated from example definition: 2025-03-01/StandbyVirtualMachinePoolRuntimeViews_Get.json
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
	// 					InstanceCountsByState: []*armstandbypool.PoolVirtualMachineStateCount{
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateCreating),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateRunning),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocating),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocated),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateStarting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeleting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernating),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernated),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Zone: to.Ptr[int64](2),
	// 					InstanceCountsByState: []*armstandbypool.PoolVirtualMachineStateCount{
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateCreating),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateRunning),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocating),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocated),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateStarting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeleting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernating),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernated),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 				{
	// 					Zone: to.Ptr[int64](3),
	// 					InstanceCountsByState: []*armstandbypool.PoolVirtualMachineStateCount{
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateCreating),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateRunning),
	// 							Count: to.Ptr[int64](20),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocating),
	// 							Count: to.Ptr[int64](10),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeallocated),
	// 							Count: to.Ptr[int64](100),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateStarting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateDeleting),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernating),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 						{
	// 							State: to.Ptr(armstandbypool.PoolVirtualMachineStateHibernated),
	// 							Count: to.Ptr[int64](0),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstandbypool.ProvisioningStateSucceeded),
	// 			Status: &armstandbypool.PoolStatus{
	// 				Code: to.Ptr(armstandbypool.HealthStateCodeHealthy),
	// 				Message: to.Ptr("The pool is healthy."),
	// 			},
	// 			Prediction: &armstandbypool.StandbyVirtualMachinePoolPrediction{
	// 				ForecastValues: &armstandbypool.StandbyVirtualMachinePoolForecastValues{
	// 					InstancesRequestedCount: []*int64{
	// 						to.Ptr[int64](24),
	// 						to.Ptr[int64](10),
	// 						to.Ptr[int64](200),
	// 						to.Ptr[int64](12),
	// 						to.Ptr[int64](5),
	// 						to.Ptr[int64](10),
	// 						to.Ptr[int64](15),
	// 						to.Ptr[int64](23),
	// 						to.Ptr[int64](56),
	// 						to.Ptr[int64](38),
	// 						to.Ptr[int64](12),
	// 						to.Ptr[int64](19),
	// 					},
	// 				},
	// 				ForecastStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-14T01:34:59.228Z"); return t}()),
	// 				ForecastInfo: to.Ptr("{\"forecastAccuracy\": 85, \"seriesUnitIntervalInMins\": 60, \"instancesRequestedCount_recentHistory\": \"[9, 4, 2, 8, 8, 2, 3, 6, 5, 3, 2, 6]\"}"),
	// 			},
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
