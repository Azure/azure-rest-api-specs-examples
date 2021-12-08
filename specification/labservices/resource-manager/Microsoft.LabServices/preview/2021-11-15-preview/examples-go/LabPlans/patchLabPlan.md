Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.1.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armlabservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/patchLabPlan.json
func ExampleLabPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		armlabservices.LabPlanUpdate{
			Properties: &armlabservices.LabPlanUpdateProperties{
				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					ClientSSHAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					WebRdpAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
					WebSSHAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
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
	log.Printf("LabPlan.ID: %s\n", *res.ID)
}
```
