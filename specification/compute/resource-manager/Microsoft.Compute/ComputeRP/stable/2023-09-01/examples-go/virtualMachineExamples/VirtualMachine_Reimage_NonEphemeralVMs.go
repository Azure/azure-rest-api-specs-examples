package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e4009d2f8d3bf0271757e522c7d1c1997e193d44/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineExamples/VirtualMachine_Reimage_NonEphemeralVMs.json
func ExampleVirtualMachinesClient_BeginReimage_reimageANonEphemeralVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginReimage(ctx, "myResourceGroup", "myVMName", &armcompute.VirtualMachinesClientBeginReimageOptions{Parameters: &armcompute.VirtualMachineReimageParameters{
		ExactVersion: to.Ptr("aaaaaa"),
		OSProfile: &armcompute.OSProfileProvisioningData{
			AdminPassword: to.Ptr("{your-password}"),
			CustomData:    to.Ptr("{your-custom-data}"),
		},
		TempDisk: to.Ptr(true),
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
