package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Pools_Get.json
func ExamplePoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPoolsClient().Get(ctx, "example-resource-group", "example-cluster", "example-pool-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.PoolsClientGetResponse{
	// 	Pool: armhorizondb.Pool{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01"),
	// 		Name: to.Ptr("example-pool-01"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/pools"),
	// 		Properties: &armhorizondb.PoolProperties{
	// 			State: to.Ptr(armhorizondb.PoolStateReady),
	// 			ComputeModel: &armhorizondb.PoolComputeModel{
	// 				Type: to.Ptr(armhorizondb.PoolComputeModelTypeProvisioned),
	// 				ComputeProcessor: to.Ptr(armhorizondb.ComputeProcessorTypeAMD),
	// 				ComputeGeneration: to.Ptr[int32](3),
	// 				VCores: to.Ptr[int32](8),
	// 			},
	// 			NodeCount: to.Ptr[int32](3),
	// 			HighAvailability: &armhorizondb.PoolHighAvailability{
	// 				Mode: to.Ptr(armhorizondb.PoolHighAvailabilityModeAnyZone),
	// 			},
	// 			ParameterGroup: &armhorizondb.PoolParameterGroupReference{
	// 				ParameterGroupResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/parameterGroups/production-parameters"),
	// 				ApplyMethod: to.Ptr(armhorizondb.ParameterApplyMethodImmediate),
	// 				SyncStatus: to.Ptr(armhorizondb.ParameterGroupSyncStatusInSync),
	// 			},
	// 			ScheduledMaintenance: &armhorizondb.SystemScheduledMaintenanceProperties{
	// 				MaintenanceManagementType: to.Ptr(armhorizondb.ScheduledMaintenanceManagementTypeSystem),
	// 			},
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
