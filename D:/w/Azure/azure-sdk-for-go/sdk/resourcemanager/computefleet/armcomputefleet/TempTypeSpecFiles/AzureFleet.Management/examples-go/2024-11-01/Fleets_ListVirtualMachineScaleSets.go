package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/TempTypeSpecFiles/AzureFleet.Management/examples/2024-11-01/Fleets_ListVirtualMachineScaleSets.json
func ExampleFleetsClient_NewListVirtualMachineScaleSetsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("1DC2F28C-A625-4B0E-9748-9885A3C9E9EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListVirtualMachineScaleSetsPager("rgazurefleet", "myFleet", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcomputefleet.FleetsClientListVirtualMachineScaleSetsResponse{
		// 	VirtualMachineScaleSetListResult: armcomputefleet.VirtualMachineScaleSetListResult{
		// 		Value: []*armcomputefleet.VirtualMachineScaleSet{
		// 			{
		// 				Name: to.Ptr("myVmss"),
		// 				ID: to.Ptr("/subscriptions/7B0CD4DB-3381-4013-9B31-FB6E6FD0FF1C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet/virtualMachineScaleSets/myVmss"),
		// 				Type: to.Ptr("Microsoft.AzureFleet/fleets/virtualMachineScaleSets"),
		// 				OperationStatus: to.Ptr(armcomputefleet.ProvisioningStateCreating),
		// 				Error: &armcomputefleet.APIError{
		// 					Details: []*armcomputefleet.APIErrorBase{
		// 						{
		// 							Code: to.Ptr("gzhtOverconstrainedAllocationRequestyosk"),
		// 							Target: to.Ptr("qfthabhrmndhfizfnrwpgxvnokpy"),
		// 							Message: to.Ptr("Allocation Failed"),
		// 						},
		// 					},
		// 					Innererror: &armcomputefleet.InnerError{
		// 						ExceptionType: to.Ptr("sfaomfpoaptnbxchrfskm"),
		// 						ErrorDetail: to.Ptr("ihjwbwykq"),
		// 					},
		// 					Code: to.Ptr("OverconstrainedAllocationRequest"),
		// 					Target: to.Ptr("nhaj"),
		// 					Message: to.Ptr("Allocation Failed"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
