Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.2.1/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationRecoveryPlans_UnplannedFailover.json
func ExampleReplicationRecoveryPlansClient_BeginUnplannedFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationRecoveryPlansClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUnplannedFailover(ctx,
		"<recovery-plan-name>",
		armrecoveryservicessiterecovery.RecoveryPlanUnplannedFailoverInput{
			Properties: &armrecoveryservicessiterecovery.RecoveryPlanUnplannedFailoverInputProperties{
				FailoverDirection: armrecoveryservicessiterecovery.PossibleOperationsDirections("PrimaryToRecovery").ToPtr(),
				ProviderSpecificDetails: []armrecoveryservicessiterecovery.RecoveryPlanProviderSpecificFailoverInputClassification{
					&armrecoveryservicessiterecovery.RecoveryPlanHyperVReplicaAzureFailoverInput{
						InstanceType: to.StringPtr("<instance-type>"),
					}},
				SourceSiteOperations: armrecoveryservicessiterecovery.SourceSiteOperations("Required").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ReplicationRecoveryPlansClientUnplannedFailoverResult)
}
```
