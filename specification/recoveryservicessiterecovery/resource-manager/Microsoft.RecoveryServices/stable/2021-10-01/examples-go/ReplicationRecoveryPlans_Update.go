package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationRecoveryPlans_Update.json
func ExampleReplicationRecoveryPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<recovery-plan-name>",
		armrecoveryservicessiterecovery.UpdateRecoveryPlanInput{
			Properties: &armrecoveryservicessiterecovery.UpdateRecoveryPlanInputProperties{
				Groups: []*armrecoveryservicessiterecovery.RecoveryPlanGroup{
					{
						EndGroupActions:           []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:                 armrecoveryservicessiterecovery.RecoveryPlanGroupTypeShutdown.ToPtr(),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{},
						StartGroupActions:         []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions:           []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:                 armrecoveryservicessiterecovery.RecoveryPlanGroupTypeFailover.ToPtr(),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{},
						StartGroupActions:         []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:       armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot.ToPtr(),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
							{
								ID:               to.StringPtr("<id>"),
								VirtualMachineID: to.StringPtr("<virtual-machine-id>"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:       armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot.ToPtr(),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
							{
								ID:               to.StringPtr("<id>"),
								VirtualMachineID: to.StringPtr("<virtual-machine-id>"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RecoveryPlan.ID: %s\n", *res.ID)
}
