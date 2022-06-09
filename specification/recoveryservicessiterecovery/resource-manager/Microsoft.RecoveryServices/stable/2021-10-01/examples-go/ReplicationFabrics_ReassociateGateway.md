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

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationFabrics_ReassociateGateway.json
func ExampleReplicationFabricsClient_BeginReassociateGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationFabricsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginReassociateGateway(ctx,
		"<fabric-name>",
		armrecoveryservicessiterecovery.FailoverProcessServerRequest{
			Properties: &armrecoveryservicessiterecovery.FailoverProcessServerRequestProperties{
				ContainerName:         to.StringPtr("<container-name>"),
				SourceProcessServerID: to.StringPtr("<source-process-server-id>"),
				TargetProcessServerID: to.StringPtr("<target-process-server-id>"),
				UpdateType:            to.StringPtr("<update-type>"),
				VMsToMigrate: []*string{
					to.StringPtr("Vm1"),
					to.StringPtr("Vm2")},
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
	log.Printf("Fabric.ID: %s\n", *res.ID)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicessiterecovery%2Fv0.1.0/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/README.md) on how to add the SDK to your project and authenticate.
