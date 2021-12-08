Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.1.0/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/putLab.json
func ExampleLabsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		armlabservices.Lab{
			TrackedResource: armlabservices.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armlabservices.LabProperties{
				LabUpdateProperties: armlabservices.LabUpdateProperties{
					Description: to.StringPtr("<description>"),
					AutoShutdownProfile: &armlabservices.AutoShutdownProfile{
						DisconnectDelay:          to.StringPtr("<disconnect-delay>"),
						IdleDelay:                to.StringPtr("<idle-delay>"),
						NoConnectDelay:           to.StringPtr("<no-connect-delay>"),
						ShutdownOnDisconnect:     armlabservices.EnableStateEnabled.ToPtr(),
						ShutdownOnIdle:           armlabservices.ShutdownOnIdleModeUserAbsence.ToPtr(),
						ShutdownWhenNotConnected: armlabservices.EnableStateEnabled.ToPtr(),
					},
					ConnectionProfile: &armlabservices.ConnectionProfile{
						ClientRdpAccess: armlabservices.ConnectionTypePublic.ToPtr(),
						ClientSSHAccess: armlabservices.ConnectionTypePublic.ToPtr(),
						WebRdpAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
						WebSSHAccess:    armlabservices.ConnectionTypeNone.ToPtr(),
					},
					LabPlanID: to.StringPtr("<lab-plan-id>"),
					SecurityProfile: &armlabservices.SecurityProfile{
						OpenAccess: armlabservices.EnableStateDisabled.ToPtr(),
					},
					Title: to.StringPtr("<title>"),
					VirtualMachineProfile: &armlabservices.VirtualMachineProfile{
						AdditionalCapabilities: &armlabservices.VirtualMachineAdditionalCapabilities{
							InstallGpuDrivers: armlabservices.EnableStateDisabled.ToPtr(),
						},
						AdminUser: &armlabservices.Credentials{
							Username: to.StringPtr("<username>"),
						},
						CreateOption: armlabservices.CreateOptionTemplateVM.ToPtr(),
						ImageReference: &armlabservices.ImageReference{
							Offer:     to.StringPtr("<offer>"),
							Publisher: to.StringPtr("<publisher>"),
							SKU:       to.StringPtr("<sku>"),
							Version:   to.StringPtr("<version>"),
						},
						SKU: &armlabservices.SKU{
							Name: to.StringPtr("<name>"),
						},
						UsageQuota:        to.StringPtr("<usage-quota>"),
						UseSharedPassword: armlabservices.EnableStateDisabled.ToPtr(),
					},
				},
				NetworkProfile: &armlabservices.LabNetworkProfile{
					SubnetID: to.StringPtr("<subnet-id>"),
				},
				State: armlabservices.LabStateDraft.ToPtr(),
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
	log.Printf("Lab.ID: %s\n", *res.ID)
}
```
