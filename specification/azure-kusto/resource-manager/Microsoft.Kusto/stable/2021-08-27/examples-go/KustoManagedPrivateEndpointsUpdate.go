package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoManagedPrivateEndpointsUpdate.json
func ExampleManagedPrivateEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewManagedPrivateEndpointsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<managed-private-endpoint-name>",
		armkusto.ManagedPrivateEndpoint{
			Properties: &armkusto.ManagedPrivateEndpointProperties{
				GroupID:               to.StringPtr("<group-id>"),
				PrivateLinkResourceID: to.StringPtr("<private-link-resource-id>"),
				RequestMessage:        to.StringPtr("<request-message>"),
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
	log.Printf("Response result: %#v\n", res.ManagedPrivateEndpointsClientUpdateResult)
}
