package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/250861bb6a886b75255edfa0aa5ee2dd0d6e7a11/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Redeploy_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginRedeploy_virtualMachineScaleSetRedeployMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginRedeploy(ctx, "rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineScaleSetsClientBeginRedeployOptions{VMInstanceIDs: &armcompute.VirtualMachineScaleSetVMInstanceIDs{
		InstanceIDs: []*string{
			to.Ptr("aaaaaaaaaaaaaaaaa")},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
