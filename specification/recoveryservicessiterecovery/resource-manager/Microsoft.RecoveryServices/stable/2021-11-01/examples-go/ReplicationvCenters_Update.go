package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationvCenters_Update.json
func ExampleReplicationvCentersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationvCentersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<fabric-name>",
		"<vcenter-name>",
		armrecoveryservicessiterecovery.UpdateVCenterRequest{
			Properties: &armrecoveryservicessiterecovery.UpdateVCenterRequestProperties{
				IPAddress: to.StringPtr("<ipaddress>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationvCentersClientUpdateResult)
}
