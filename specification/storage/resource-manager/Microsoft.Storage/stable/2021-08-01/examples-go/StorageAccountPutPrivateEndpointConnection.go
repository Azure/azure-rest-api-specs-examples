package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/StorageAccountPutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Put(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-name>",
		armstorage.PrivateEndpointConnection{
			Properties: &armstorage.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armstorage.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armstorage.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientPutResult)
}
