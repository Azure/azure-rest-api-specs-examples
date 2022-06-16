package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ApproveRejectSitePrivateEndpointConnection.json
func ExampleWebAppsClient_BeginApproveOrRejectPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappservice.NewWebAppsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginApproveOrRejectPrivateEndpointConnection(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		armappservice.PrivateLinkConnectionApprovalRequestResource{
			Properties: &armappservice.PrivateLinkConnectionApprovalRequest{
				PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          to.StringPtr("<status>"),
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
	log.Printf("RemotePrivateEndpointConnectionARMResource.ID: %s\n", *res.ID)
}
