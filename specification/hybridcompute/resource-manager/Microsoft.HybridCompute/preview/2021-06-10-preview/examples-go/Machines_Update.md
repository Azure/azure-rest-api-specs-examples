Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.2.1/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.MachineUpdate{
			Identity: &armhybridcompute.Identity{
				Type: to.StringPtr("<type>"),
			},
			Properties: &armhybridcompute.MachineUpdateProperties{
				LocationData: &armhybridcompute.LocationData{
					Name: to.StringPtr("<name>"),
				},
				OSProfile: &armhybridcompute.OSProfile{
					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.StringPtr("<assessment-mode>"),
						},
					},
					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.StringPtr("<assessment-mode>"),
						},
					},
				},
				ParentClusterResourceID:    to.StringPtr("<parent-cluster-resource-id>"),
				PrivateLinkScopeResourceID: to.StringPtr("<private-link-scope-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MachinesClientUpdateResult)
}
```
