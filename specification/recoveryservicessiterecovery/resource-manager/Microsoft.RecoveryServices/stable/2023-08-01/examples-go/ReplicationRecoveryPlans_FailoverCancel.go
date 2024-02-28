package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationRecoveryPlans_FailoverCancel.json
func ExampleReplicationRecoveryPlansClient_BeginFailoverCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationRecoveryPlansClient().BeginFailoverCancel(ctx, "vault1", "resourceGroupPS1", "RPtest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecoveryPlan = armrecoveryservicessiterecovery.RecoveryPlan{
	// 	Name: to.Ptr("RPtest1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationRecoveryPlans"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationRecoveryPlans/RPtest1"),
	// 	Properties: &armrecoveryservicessiterecovery.RecoveryPlanProperties{
	// 		AllowedOperations: []*string{
	// 			to.Ptr("PlannedFailover")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/5276a7bc-12a3-43a1-bc53-9bf80e0be87b"),
	// 				ScenarioName: to.Ptr("CancelFailover"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-24T08:52:42.158Z"); return t}()),
	// 			},
	// 			CurrentScenarioStatus: to.Ptr("Succeeded"),
	// 			CurrentScenarioStatusDescription: to.Ptr("Completed"),
	// 			FailoverDeploymentModel: to.Ptr("ResourceManager"),
	// 			FriendlyName: to.Ptr("RPtest1"),
	// 			Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
	// 			},
	// 			LastPlannedFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-24T07:33:49.137Z"); return t}()),
	// 			LastTestFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1601-01-01T00:00:00.000Z"); return t}()),
	// 			PrimaryFabricFriendlyName: to.Ptr("cloud1"),
	// 			PrimaryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			RecoveryFabricFriendlyName: to.Ptr("cloud1"),
	// 			RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			ReplicationProviders: []*string{
	// 				to.Ptr("InMageRcmFailback")},
	// 			},
	// 		}
}
