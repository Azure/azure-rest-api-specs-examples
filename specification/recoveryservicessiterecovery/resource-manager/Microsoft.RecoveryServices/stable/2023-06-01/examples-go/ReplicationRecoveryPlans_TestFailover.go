package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryPlans_TestFailover.json
func ExampleReplicationRecoveryPlansClient_BeginTestFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationRecoveryPlansClient().BeginTestFailover(ctx, "vault1", "resourceGroupPS1", "RPtest1", armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInput{
		Properties: &armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInputProperties{
			FailoverDirection: to.Ptr(armrecoveryservicessiterecovery.PossibleOperationsDirectionsPrimaryToRecovery),
			NetworkID:         to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
			NetworkType:       to.Ptr("VmNetworkAsInput"),
			ProviderSpecificDetails: []armrecoveryservicessiterecovery.RecoveryPlanProviderSpecificFailoverInputClassification{
				&armrecoveryservicessiterecovery.RecoveryPlanHyperVReplicaAzureFailoverInput{
					InstanceType: to.Ptr("HyperVReplicaAzure"),
				}},
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
	// 			to.Ptr("TestFailoverCleanup")},
	// 			CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 				JobID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/d40bfb40-aaaa-4c0d-87d3-41b15439a84b"),
	// 				ScenarioName: to.Ptr("TestFailover"),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-04-27T07:00:58.8191916Z"); return t}()),
	// 			},
	// 			CurrentScenarioStatus: to.Ptr("Suspended"),
	// 			CurrentScenarioStatusDescription: to.Ptr("WaitingForStopTestFailover"),
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
	// 				},
	// 				{
	// 					EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{
	// 					},
	// 					GroupType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
	// 					ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
	// 						{
	// 							ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
	// 							VirtualMachineID: to.Ptr("c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
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
