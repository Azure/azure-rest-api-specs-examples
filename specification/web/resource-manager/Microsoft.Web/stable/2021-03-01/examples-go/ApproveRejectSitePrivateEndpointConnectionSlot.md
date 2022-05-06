Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappservice%2Farmappservice%2Fv0.4.0/sdk/resourcemanager/appservice/armappservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armappservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ApproveRejectSitePrivateEndpointConnectionSlot.json
func ExampleWebAppsClient_BeginApproveOrRejectPrivateEndpointConnectionSlot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armappservice.NewWebAppsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginApproveOrRejectPrivateEndpointConnectionSlot(ctx,
		"<resource-group-name>",
		"<name>",
		"<private-endpoint-connection-name>",
		"<slot>",
		armappservice.PrivateLinkConnectionApprovalRequestResource{
			Properties: &armappservice.PrivateLinkConnectionApprovalRequest{
				PrivateLinkServiceConnectionState: &armappservice.PrivateLinkConnectionState{
					Description:     to.Ptr("<description>"),
					ActionsRequired: to.Ptr("<actions-required>"),
					Status:          to.Ptr("<status>"),
				},
			},
		},
		&armappservice.WebAppsClientBeginApproveOrRejectPrivateEndpointConnectionSlotOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
