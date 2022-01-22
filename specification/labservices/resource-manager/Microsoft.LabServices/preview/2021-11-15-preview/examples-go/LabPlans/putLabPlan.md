Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.2.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armlabservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/putLabPlan.json
func ExampleLabPlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabPlansClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		armlabservices.LabPlan{
			Location: to.StringPtr("<location>"),
			Properties: &armlabservices.LabPlanProperties{
				DefaultAutoShutdownProfile: &armlabservices.AutoShutdownProfile{
					DisconnectDelay:          to.StringPtr("<disconnect-delay>"),
					IdleDelay:                to.StringPtr("<idle-delay>"),
					NoConnectDelay:           to.StringPtr("<no-connect-delay>"),
					ShutdownOnDisconnect:     armlabservices.EnableStateEnabled.ToPtr(),
					ShutdownOnIdle:           armlabservices.ShutdownOnIdleModeUserAbsence.ToPtr(),
					ShutdownWhenNotConnected: armlabservices.EnableStateEnabled.ToPtr(),
				},
				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					ClientSSHAccess: armlabservices.ConnectionTypePublic.ToPtr(),
					WebRdpAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
					WebSSHAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
				},
				DefaultNetworkProfile: &armlabservices.LabPlanNetworkProfile{
					SubnetID: to.StringPtr("<subnet-id>"),
				},
				SharedGalleryID: to.StringPtr("<shared-gallery-id>"),
				SupportInfo: &armlabservices.SupportInfo{
					Email:        to.StringPtr("<email>"),
					Instructions: to.StringPtr("<instructions>"),
					Phone:        to.StringPtr("<phone>"),
					URL:          to.StringPtr("<url>"),
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
	log.Printf("Response result: %#v\n", res.LabPlansClientCreateOrUpdateResult)
}
```
