Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ApproveRejectPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<private-endpoint-connection-name>",
		armdatafactory.PrivateLinkConnectionApprovalRequestResource{
			Properties: &armdatafactory.PrivateLinkConnectionApprovalRequest{
				PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
					Description:     to.StringPtr("<description>"),
					ActionsRequired: to.StringPtr("<actions-required>"),
					Status:          to.StringPtr("<status>"),
				},
			},
		},
		&armdatafactory.PrivateEndpointConnectionClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientCreateOrUpdateResult)
}
```
