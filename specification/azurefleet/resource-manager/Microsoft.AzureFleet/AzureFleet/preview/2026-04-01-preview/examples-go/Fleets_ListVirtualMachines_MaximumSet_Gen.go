package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/v2"
)

// Generated from example definition: 2026-04-01-preview/Fleets_ListVirtualMachines_MaximumSet_Gen.json
func ExampleFleetsClient_NewListVirtualMachinesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("8F7446E8-AD7B-4D50-989A-2504374B8462", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFleetsClient().NewListVirtualMachinesPager("rgazurefleet", "testFleet", &armcomputefleet.FleetsClientListVirtualMachinesOptions{
		Filter:    to.Ptr("qppsnaauhedxu"),
		Skiptoken: to.Ptr("jxgpugummyphgx")})
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
		// page = armcomputefleet.FleetsClientListVirtualMachinesResponse{
		// 	VirtualMachineListResult: armcomputefleet.VirtualMachineListResult{
		// 		Value: []*armcomputefleet.VirtualMachine{
		// 			{
		// 				Name: to.Ptr("test-vm_aef123_0"),
		// 				ID: to.Ptr("/subscriptions/3453D930-6DDF-4466-B3B3-E1AEE9BD448C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet/virtualMachines/test-vm_aef123_0"),
		// 				Type: to.Ptr("lmzwlqfgyp"),
		// 				OperationStatus: to.Ptr(armcomputefleet.VMOperationStatusLaunching),
		// 			},
		// 			{
		// 				Name: to.Ptr("test-vm_aef123_1"),
		// 				ID: to.Ptr("/subscriptions/3453D930-6DDF-4466-B3B3-E1AEE9BD448C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet/virtualMachines/test-vm_aef123_1"),
		// 				Type: to.Ptr("lmzwlqfgyp"),
		// 				OperationStatus: to.Ptr(armcomputefleet.VMOperationStatusCreating),
		// 			},
		// 			{
		// 				Name: to.Ptr("test-vm_aef123_2"),
		// 				ID: to.Ptr("/subscriptions/3453D930-6DDF-4466-B3B3-E1AEE9BD448C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet/virtualMachines/test-vm_aef123_2"),
		// 				Type: to.Ptr("lmzwlqfgyp"),
		// 				OperationStatus: to.Ptr(armcomputefleet.VMOperationStatusSucceeded),
		// 			},
		// 			{
		// 				Name: to.Ptr("test-vm_aef123_3"),
		// 				ID: to.Ptr("/subscriptions/3453D930-6DDF-4466-B3B3-E1AEE9BD448C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/myFleet/virtualMachines/test-vm_aef123_3"),
		// 				OperationStatus: to.Ptr(armcomputefleet.VMOperationStatusFailed),
		// 				Error: &armcomputefleet.APIError{
		// 					Code: to.Ptr("obrkfckgffjphilhxacxdmrvbcrg"),
		// 					Target: to.Ptr("rnjvbpwlqnicuuyk"),
		// 					Message: to.Ptr("oiqdlxfufmymom"),
		// 					Details: []*armcomputefleet.APIErrorBase{
		// 						{
		// 							Code: to.Ptr("vtuyxkhxumtapurmhkv"),
		// 							Target: to.Ptr("wfkewwpzippudvjizcnwgzspbeq"),
		// 							Message: to.Ptr("hjoitvryetqmjv"),
		// 						},
		// 					},
		// 					Innererror: &armcomputefleet.InnerError{
		// 						ExceptionType: to.Ptr("bybvkxpntnkyrzdluphztenxzjy"),
		// 						ErrorDetail: to.Ptr("lsrcyfnkecxzioyagpgsbebgvfkwuc"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
