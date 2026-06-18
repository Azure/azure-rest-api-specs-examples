package armresiliencemanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resiliencemanagement/armresiliencemanagement"
)

// Generated from example definition: 2026-04-01-preview/Drills_AddOrUpdateResources_MaximumSet_Gen.json
func ExampleDrillsClient_BeginAddOrUpdateResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresiliencemanagement.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDrillsClient().BeginAddOrUpdateResources(ctx, "sampleServiceGroupName", "qmn", "drill1", armresiliencemanagement.AddOrUpdateResourcesRequest{
		FaultDurationInMin: to.Ptr[int32](0),
		ResourceLists: &armresiliencemanagement.ResourceLists{
			IncludeResources: []*armresiliencemanagement.IncludeOrUpdateResource{
				{
					FaultProperties: &armresiliencemanagement.FaultProperties{
						CustomFault: &armresiliencemanagement.CustomFaultDetails{
							FaultName:        to.Ptr("umofuzwgczqwyzcoakmrdrkjknykdonhypxibwrweggltsmjayvnlzroxdfalwkfsqvuqtfwhhzcnemndbgxdiciqs"),
							ScriptResourceID: to.Ptr("/subscriptions/191973cd-9c54-41e0-ac19-25dd9a92d5a8/resourceGroups/abhinkRG/providers/Microsoft.Automation/automationAccounts/abhinkAcc/runbooks/viveksi"),
						},
						DefaultFault: &armresiliencemanagement.FaultDetails{
							FaultUrn:         to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
							FaultName:        to.Ptr("shutdown"),
							TargetResourceID: to.Ptr("/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/vm1"),
						},
						OverriddenDefaultFault: &armresiliencemanagement.FaultDetails{
							FaultUrn:         to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
							FaultName:        to.Ptr("shutdown"),
							TargetResourceID: to.Ptr("/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/vm1"),
						},
					},
					ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/drills/drill1/drillResources/2c9b3a1f-f96e-42c2-98fe-15005da8a133"),
				},
			},
			ExcludeResources: []*string{
				to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/drills/drill1/drillResources/c2191964-be24-4849-8faf-d9569576c708"),
			},
			UpdateResources: []*armresiliencemanagement.IncludeOrUpdateResource{
				{
					FaultProperties: &armresiliencemanagement.FaultProperties{
						CustomFault: &armresiliencemanagement.CustomFaultDetails{
							FaultName:        to.Ptr("umofuzwgczqwyzcoakmrdrkjknykdonhypxibwrweggltsmjayvnlzroxdfalwkfsqvuqtfwhhzcnemndbgxdiciqs"),
							ScriptResourceID: to.Ptr("/subscriptions/191973cd-9c54-41e0-ac19-25dd9a92d5a8/resourceGroups/abhinkRG/providers/Microsoft.Automation/automationAccounts/abhinkAcc/runbooks/viveksi"),
						},
						DefaultFault: &armresiliencemanagement.FaultDetails{
							FaultUrn:         to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
							FaultName:        to.Ptr("shutdown"),
							TargetResourceID: to.Ptr("/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/vm1"),
						},
						OverriddenDefaultFault: &armresiliencemanagement.FaultDetails{
							FaultUrn:         to.Ptr("urn:csci:microsoft:virtualMachine:shutdown/1.0"),
							FaultName:        to.Ptr("shutdown"),
							TargetResourceID: to.Ptr("/subscriptions/f2edfd5d-5496-4683-b94f-b3588c579009/resourceGroups/testRG/providers/Microsoft.Compute/virtualMachines/vm1"),
						},
					},
					ID: to.Ptr("/providers/Microsoft.Management/serviceGroups/sampleServiceGroupName/providers/Microsoft.AzureResilienceManagement/drills/drill1/drillResources/c26bea42-c34c-4e6f-8cf4-15043e18c8bc"),
				},
			},
		},
		ForceInclusionAndUpdate: to.Ptr(armresiliencemanagement.ForceInclusionAndUpdateEnable),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
