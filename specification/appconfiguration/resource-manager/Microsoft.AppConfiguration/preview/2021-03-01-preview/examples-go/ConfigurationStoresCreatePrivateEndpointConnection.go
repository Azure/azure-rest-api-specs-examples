package armappconfiguration_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresCreatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		"<private-endpoint-connection-name>",
		armappconfiguration.PrivateEndpointConnection{
			Properties: &armappconfiguration.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armappconfiguration.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armappconfiguration.ConnectionStatus("Approved").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientCreateOrUpdateResult)
}
