Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.4.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/putLab.json
func ExampleLabsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		armlabservices.Lab{
			Location: to.Ptr("<location>"),
			Properties: &armlabservices.LabProperties{
				Description: to.Ptr("<description>"),
				AutoShutdownProfile: &armlabservices.AutoShutdownProfile{
					DisconnectDelay:          to.Ptr("<disconnect-delay>"),
					IdleDelay:                to.Ptr("<idle-delay>"),
					NoConnectDelay:           to.Ptr("<no-connect-delay>"),
					ShutdownOnDisconnect:     to.Ptr(armlabservices.EnableStateEnabled),
					ShutdownOnIdle:           to.Ptr(armlabservices.ShutdownOnIdleModeUserAbsence),
					ShutdownWhenNotConnected: to.Ptr(armlabservices.EnableStateEnabled),
				},
				ConnectionProfile: &armlabservices.ConnectionProfile{
					ClientRdpAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					ClientSSHAccess: to.Ptr(armlabservices.ConnectionTypePublic),
					WebRdpAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
					WebSSHAccess:    to.Ptr(armlabservices.ConnectionTypeNone),
				},
				LabPlanID: to.Ptr("<lab-plan-id>"),
				SecurityProfile: &armlabservices.SecurityProfile{
					OpenAccess: to.Ptr(armlabservices.EnableStateDisabled),
				},
				Title: to.Ptr("<title>"),
				VirtualMachineProfile: &armlabservices.VirtualMachineProfile{
					AdditionalCapabilities: &armlabservices.VirtualMachineAdditionalCapabilities{
						InstallGpuDrivers: to.Ptr(armlabservices.EnableStateDisabled),
					},
					AdminUser: &armlabservices.Credentials{
						Username: to.Ptr("<username>"),
					},
					CreateOption: to.Ptr(armlabservices.CreateOptionTemplateVM),
					ImageReference: &armlabservices.ImageReference{
						Offer:     to.Ptr("<offer>"),
						Publisher: to.Ptr("<publisher>"),
						SKU:       to.Ptr("<sku>"),
						Version:   to.Ptr("<version>"),
					},
					SKU: &armlabservices.SKU{
						Name: to.Ptr("<name>"),
					},
					UsageQuota:        to.Ptr("<usage-quota>"),
					UseSharedPassword: to.Ptr(armlabservices.EnableStateDisabled),
				},
				NetworkProfile: &armlabservices.LabNetworkProfile{
					SubnetID: to.Ptr("<subnet-id>"),
				},
				State: to.Ptr(armlabservices.LabStateDraft),
			},
		},
		&armlabservices.LabsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
