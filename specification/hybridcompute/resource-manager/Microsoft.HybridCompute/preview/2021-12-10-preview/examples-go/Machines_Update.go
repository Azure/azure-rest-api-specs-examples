package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-12-10-preview/examples/Machines_Update.json
func ExampleMachinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.MachineUpdate{
			Identity: &armhybridcompute.Identity{
				Type: to.Ptr("<type>"),
			},
			Properties: &armhybridcompute.MachineUpdateProperties{
				LocationData: &armhybridcompute.LocationData{
					Name: to.Ptr("<name>"),
				},
				OSProfile: &armhybridcompute.OSProfile{
					LinuxConfiguration: &armhybridcompute.OSProfileLinuxConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.Ptr("<assessment-mode>"),
							PatchMode:      to.Ptr("<patch-mode>"),
						},
					},
					WindowsConfiguration: &armhybridcompute.OSProfileWindowsConfiguration{
						PatchSettings: &armhybridcompute.PatchSettings{
							AssessmentMode: to.Ptr("<assessment-mode>"),
							PatchMode:      to.Ptr("<patch-mode>"),
						},
					},
				},
				ParentClusterResourceID:    to.Ptr("<parent-cluster-resource-id>"),
				PrivateLinkScopeResourceID: to.Ptr("<private-link-scope-resource-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
