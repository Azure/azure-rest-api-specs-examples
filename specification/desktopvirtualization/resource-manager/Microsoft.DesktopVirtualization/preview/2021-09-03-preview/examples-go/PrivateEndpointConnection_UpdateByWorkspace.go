package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
)

// x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateEndpointConnection_UpdateByWorkspace.json
func ExamplePrivateEndpointConnectionsClient_UpdateByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdesktopvirtualization.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.UpdateByWorkspace(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<private-endpoint-connection-name>",
		armdesktopvirtualization.PrivateEndpointConnection{
			Properties: &armdesktopvirtualization.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armdesktopvirtualization.PrivateLinkServiceConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          armdesktopvirtualization.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientUpdateByWorkspaceResult)
}
