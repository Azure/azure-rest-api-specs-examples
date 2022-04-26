Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.4.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryPlans_Update.json
func ExampleReplicationRecoveryPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<recovery-plan-name>",
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
								ID:               to.Ptr("<id>"),
								VirtualMachineID: to.Ptr("<virtual-machine-id>"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					},
					{
						EndGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
						GroupType:       to.Ptr(armrecoveryservicessiterecovery.RecoveryPlanGroupTypeBoot),
						ReplicationProtectedItems: []*armrecoveryservicessiterecovery.RecoveryPlanProtectedItem{
							{
								ID:               to.Ptr("<id>"),
								VirtualMachineID: to.Ptr("<virtual-machine-id>"),
							}},
						StartGroupActions: []*armrecoveryservicessiterecovery.RecoveryPlanAction{},
					}},
			},
		},
		&armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
