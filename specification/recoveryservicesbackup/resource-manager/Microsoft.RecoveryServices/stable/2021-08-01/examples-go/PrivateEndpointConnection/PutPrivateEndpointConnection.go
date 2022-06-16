package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/PrivateEndpointConnection/PutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPut(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<private-endpoint-connection-name>",
		armrecoveryservicesbackup.PrivateEndpointConnectionResource{
			Properties: &armrecoveryservicesbackup.PrivateEndpointConnection{
				PrivateEndpoint: &armrecoveryservicesbackup.PrivateEndpoint{
					ID: to.StringPtr("<id>"),
				},
				PrivateLinkServiceConnectionState: &armrecoveryservicesbackup.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armrecoveryservicesbackup.PrivateEndpointConnectionStatusApproved.ToPtr(),
				},
				ProvisioningState: armrecoveryservicesbackup.ProvisioningStateSucceeded.ToPtr(),
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
	log.Printf("PrivateEndpointConnectionResource.ID: %s\n", *res.ID)
}
