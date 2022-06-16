package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-11-01/examples/ReplicationvCenters_Create.json
func ExampleReplicationvCentersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationvCentersClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<fabric-name>",
		"<vcenter-name>",
		armrecoveryservicessiterecovery.AddVCenterRequest{
			Properties: &armrecoveryservicessiterecovery.AddVCenterRequestProperties{
				FriendlyName:    to.StringPtr("<friendly-name>"),
				IPAddress:       to.StringPtr("<ipaddress>"),
				Port:            to.StringPtr("<port>"),
				ProcessServerID: to.StringPtr("<process-server-id>"),
				RunAsAccountID:  to.StringPtr("<run-as-account-id>"),
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
	log.Printf("Response result: %#v\n", res.ReplicationvCentersClientCreateResult)
}
