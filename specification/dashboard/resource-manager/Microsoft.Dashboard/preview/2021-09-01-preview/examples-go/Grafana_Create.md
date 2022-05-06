Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdashboard%2Farmdashboard%2Fv0.2.0/sdk/resourcemanager/dashboard/armdashboard/README.md) on how to add the SDK to your project and authenticate.

```go
package armdashboard_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_Create.json
func ExampleGrafanaClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdashboard.NewGrafanaClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armdashboard.ManagedGrafana{
			Identity: &armdashboard.ManagedIdentity{
				Type: to.Ptr(armdashboard.IdentityTypeSystemAssigned),
			},
			Location: to.Ptr("<location>"),
			Properties: &armdashboard.ManagedGrafanaProperties{
				ProvisioningState: to.Ptr(armdashboard.ProvisioningStateAccepted),
				ZoneRedundancy:    to.Ptr(armdashboard.ZoneRedundancyEnabled),
			},
			SKU: &armdashboard.ResourceSKU{
				Name: to.Ptr("<name>"),
			},
			Tags: map[string]*string{
				"Environment": to.Ptr("Dev"),
			},
		},
		&armdashboard.GrafanaClientBeginCreateOptions{ResumeToken: ""})
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
