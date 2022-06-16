package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryPlans_Update.json
func ExampleReplicationRecoveryPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient("vault1",
		"resourceGroupPS1",
		"c183865e-6077-46f2-a3b1-deb0f4f4650a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"RPtest1",
		armrecoveryservicessiterecovery.UpdateRecoveryPlanInput{
			Properties: &armrecoveryservicessiterecovery.UpdateRecoveryPlanInputProperties{
				Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
					{
						EndGroupActions:           []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:                 to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeShutdown),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{},
						StartGroupActions:         []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions:           []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:                 to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeFailover),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{},
						StartGroupActions:         []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:       to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
							{
								ID:               to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
								VirtualMachineID: to.Ptr("f8491e4f-817a-40dd-a90c-af773978c75b"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:       to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
							{
								ID:               to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
								VirtualMachineID: to.Ptr("c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
