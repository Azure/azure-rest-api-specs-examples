package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationRecoveryPlans_Create.json
func ExampleReplicationRecoveryPlansClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationRecoveryPlansClient().BeginCreate(ctx, "vault1", "resourceGroupPS1", "RPtest1", armrecoveryservicessiterecovery.CreateRecoveryPlanInput{
		Properties: &armrecoveryservicessiterecovery.CreateRecoveryPlanInputProperties{
			FailoverDeploymentModel: to.Ptr(armrecoveryservicessiterecovery.FailoverDeploymentModelResourceManager),
			Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
				{
					EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					GroupType:       to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
					ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
						{
							ID:               to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
							VirtualMachineID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
						}},
					StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
				}},
			PrimaryFabricID:  to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
			RecoveryFabricID: to.Ptr("Microsoft Azure"),
		},
	}, nil)
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
	// 			to.Ptr("PlannedFailover"),
	// 			to.Ptr("UnplannedFailover"),
	// 			to.Ptr("TestFailover")},
	// 			FailoverDeploymentModel: to.Ptr("ResourceManager"),
	// 			FriendlyName: to.Ptr("RPtest1"),
	// 			Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
	// 				{
	// 					EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 					GroupType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeShutdown),
	// 					ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
	// 					},
	// 					StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 				},
	// 				{
	// 					EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 					GroupType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeFailover),
	// 					ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
	// 					},
	// 					StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 				},
	// 				{
	// 					EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 					GroupType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
	// 					ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
	// 						{
	// 							ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 							VirtualMachineID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
	// 					}},
	// 					StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 			}},
	// 			PrimaryFabricFriendlyName: to.Ptr("cloud1"),
	// 			PrimaryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 			RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
	// 			RecoveryFabricID: to.Ptr("Microsoft Azure"),
	// 			ReplicationProviders: []*string{
	// 				to.Ptr("HyperVReplicaAzure")},
	// 			},
	// 		}
}
