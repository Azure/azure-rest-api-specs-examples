package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationRecoveryPlans_PlannedFailover.json
func ExampleReplicationRecoveryPlansClient_BeginPlannedFailover() {
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
	poller, err := client.BeginPlannedFailover(ctx,
		"RPtest1",
		armrecoveryservicessiterecovery.RecoveryPlanPlannedFailoverInput{
			Properties: &armrecoveryservicessiterecovery.RecoveryPlanPlannedFailoverInputProperties{
				FailoverDirection: to.Ptr(armrecoveryservicessiterecovery.PossibleOperationsDirectionsPrimaryToRecovery),
				ProviderSpecificDetails: []armrecoveryservicessiterecovery.RecoveryPlanProviderSpecificFailoverInputClassification{
					&armrecoveryservicessiterecovery.RecoveryPlanHyperVReplicaAzureFailoverInput{
						InstanceType: to.Ptr("HyperVReplicaAzure"),
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
