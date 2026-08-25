package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/LaunchBulkInstancesOperation_ListVirtualMachines_MaximumSet_Gen.json
func ExampleLaunchBulkInstancesOperationClient_NewListVirtualMachinesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLaunchBulkInstancesOperationClient().NewListVirtualMachinesPager("rgBulkactions", "useast2euap", "b038ec94-0860-42a5-b149-f1ce5f144e15", &armbulkactions.LaunchBulkInstancesOperationClientListVirtualMachinesOptions{
		Filter:    to.Ptr("onywxjwswbhlbkbbusgmkfgabdku"),
		Skiptoken: to.Ptr("tcbhwfqtoiwnlbjdbsnukxpgpa")})
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
		// page = armbulkactions.LaunchBulkInstancesOperationClientListVirtualMachinesResponse{
		// 	VirtualMachineListResult: armbulkactions.VirtualMachineListResult{
		// 		Value: []*armbulkactions.VirtualMachine{
		// 			{
		// 				Name: to.Ptr("tfzaoaeonndvvodjtefvomwwcrbe"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
		// 				Type: to.Ptr("gn"),
		// 				OperationStatus: to.Ptr(armbulkactions.VMOperationStatusCreating),
		// 				Error: &armbulkactions.APIError{
		// 					Code: to.Ptr("aqflblhj"),
		// 					Target: to.Ptr("dwviiuxyqxhktohwuflyuaguaqh"),
		// 					Message: to.Ptr("magtob"),
		// 					Details: []*armbulkactions.APIErrorBase{
		// 						{
		// 							Code: to.Ptr("emlgsqdgjxzlc"),
		// 							Target: to.Ptr("kxwrjuqcxmtqvjfdhsh"),
		// 							Message: to.Ptr("kwfytcvltzldrfrbbjrxtoexgzr"),
		// 						},
		// 					},
		// 					Innererror: &armbulkactions.BulkInstancesInnerError{
		// 						ExceptionType: to.Ptr("skecnvde"),
		// 						ErrorDetail: to.Ptr("ssejyhopfvcartkes"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
