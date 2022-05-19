Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.5.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/putLabPlan.json
func ExampleLabPlansClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabPlansClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"testrg123",
		"testlabplan",
		armlabservices.LabPlan{
			Location: to.Ptr("westus"),
			Properties: &armlabservices.LabPlanProperties{
				DefaultAutoShutdownProfile: &armlabservices.AutoShutdownProfile{
					DisconnectDelay:          to.Ptr("00:05"),
					IdleDelay:                to.Ptr("01:00"),
					NoConnectDelay:           to.Ptr("01:00"),
					ShutdownOnDisconnect:     to.Ptr(armlabservices.EnableStateEnabled),
					ShutdownOnIdle:           to.Ptr(armlabservices.ShutdownOnIdleModeUserAbsence),
					ShutdownWhenNotConnected: to.Ptr(armlabservices.EnableStateEnabled),
				},
				DefaultConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					ClientSSHAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					WebRdpAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
					WebSSHAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
				},
				DefaultNetworkProfile: &armlabservices.LabPlanNetworkProfile{
					SubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
				},
				SharedGalleryID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Compute/galleries/testsig"),
				SupportInfo: &armlabservices.SupportInfo{
					Email:        to.Ptr("help@contoso.com"),
					Instructions: to.Ptr("Contact support for help."),
					Phone:        to.Ptr("+1-202-555-0123"),
					URL:          to.Ptr("help.contoso.com"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
