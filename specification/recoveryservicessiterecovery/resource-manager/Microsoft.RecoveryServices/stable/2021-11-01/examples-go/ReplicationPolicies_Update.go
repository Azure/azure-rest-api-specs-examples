package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationPolicies_Update.json
func ExampleReplicationPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationPoliciesClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<policy-name>",
		armrecoveryservicessiterecovery.UpdatePolicyInput{
			Properties: &armrecoveryservicessiterecovery.UpdatePolicyInputProperties{
				ReplicationProviderSettings: &armrecoveryservicessiterecovery.HyperVReplicaAzurePolicyInput{
					InstanceType: to.StringPtr("<instance-type>"),
				},
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
	log.Printf("Response result: %#v\n", res.ReplicationPoliciesClientUpdateResult)
}
