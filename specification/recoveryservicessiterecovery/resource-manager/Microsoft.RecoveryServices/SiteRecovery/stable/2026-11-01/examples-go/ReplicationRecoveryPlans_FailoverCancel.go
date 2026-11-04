package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v4"
)

// Generated from example definition: 2026-11-01/ReplicationRecoveryPlans_FailoverCancel.json
func ExampleReplicationRecoveryPlansClient_BeginFailoverCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationRecoveryPlansClient().BeginFailoverCancel(ctx, "resourceGroupPS1", "vault1", "RPtest1", nil)
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
	// res = armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCancelResponse{
	// 	RecoveryPlan: armrecoveryservicessiterecovery.RecoveryPlan{
	// 		Name: to.Ptr("RPtest1"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans"),
	// 		ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationRecoveryPlans/RPtest1"),
	// 		Properties: &armrecoveryservicessiterecovery.RecoveryPlanProperties{
	// 			AllowedOperations: []*string{
	// 				to.Ptr("PlannedFailover"),
	// 			},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/5276a7bc-12a3-43a1-bc53-9bf80e0be87b"),
	// 				ScenarioName: to.Ptr("CancelFailover"),
	// 				StartTime: to.Ptr(time.Date(2021, time.February, 24, 8, 52, 42, 158759200, time.UTC)),
	// 			},
	// 			CurrentScenarioStatus: to.Ptr("Succeeded"),
	// 			CurrentScenarioStatusDescription: to.Ptr("Completed"),
	// 			FailoverDeploymentModel: to.Ptr("ResourceManager"),
	// 			FriendlyName: to.Ptr("RPtest1"),
	// 			Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
	// 			},
	// 			LastPlannedFailoverTime: to.Ptr(time.Date(2021, time.February, 24, 7, 33, 49, 137244800, time.UTC)),
	// 			LastTestFailoverTime: to.Ptr(time.Date(1601, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 			PrimaryFabricFriendlyName: to.Ptr("cloud1"),
	// 			PrimaryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			RecoveryFabricFriendlyName: to.Ptr("cloud1"),
	// 			RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			ReplicationProviders: []*string{
	// 				to.Ptr("InMageRcmFailback"),
	// 			},
	// 		},
	// 	},
	// }
}
