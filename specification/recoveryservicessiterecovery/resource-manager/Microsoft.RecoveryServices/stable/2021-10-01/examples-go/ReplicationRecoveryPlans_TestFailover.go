package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationRecoveryPlans_TestFailover.json
func ExampleReplicationRecoveryPlansClient_BeginTestFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginTestFailover(ctx,
		"<recovery-plan-name>",
		armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInput{
			Properties: &armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInputProperties{
				FailoverDirection: armrecoveryservicessiterecovery.PossibleOperationsDirectionsPrimaryToRecovery.ToPtr(),
				NetworkID:         to.StringPtr("<network-id>"),
				NetworkType:       to.StringPtr("<network-type>"),
				ProviderSpecificDetails: []armrecoveryservicessiterecovery.RecoveryPlanProviderSpecificFailoverInputClassification{
					&armrecoveryservicessiterecovery.RecoveryPlanHyperVReplicaAzureFailoverInput{
						RecoveryPlanProviderSpecificFailoverInput: armrecoveryservicessiterecovery.RecoveryPlanProviderSpecificFailoverInput{
							InstanceType: to.StringPtr("<instance-type>"),
						},
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
