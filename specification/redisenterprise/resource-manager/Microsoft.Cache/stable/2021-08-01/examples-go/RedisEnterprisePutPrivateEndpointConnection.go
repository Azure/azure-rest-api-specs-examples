package armredisenterprise_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2021-08-01/examples/RedisEnterprisePutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredisenterprise.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPut(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<private-endpoint-connection-name>",
		armredisenterprise.PrivateEndpointConnection{
			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armredisenterprise.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
