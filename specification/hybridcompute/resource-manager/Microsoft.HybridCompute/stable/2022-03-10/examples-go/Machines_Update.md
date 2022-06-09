```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2022-03-10/examples/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewMachinesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myResourceGroup",
		"myMachine",
		armhybridcompute.MachineUpdate{
			Identity: &armhybridcompute.Identity{
				Type: to.Ptr("SystemAssigned"),
			},
			Properties: &armhybridcompute.MachineUpdateProperties{
				LocationData: &armhybridcompute.LocationData{
					Name: to.Ptr("Redmond"),
				},
				OSProfile: &armhybridcompute.OSProfile{
					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
							PatchMode:      to.Ptr(armhybridcompute.PatchModeTypesManual),
						},
					},
					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.Ptr(armhybridcompute.AssessmentModeTypesImageDefault),
							PatchMode:      to.Ptr(armhybridcompute.PatchModeTypesManual),
						},
					},
				},
				ParentClusterResourceID:    to.Ptr("{AzureStackHCIResourceId}"),
				PrivateLinkScopeResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/privateLinkScopes/privateLinkScopeName"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv1.0.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.
