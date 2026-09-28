package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-10-06-preview/VirtualMachineBulkOperations_BulkCancel_PartialSuccess.json
func ExampleVirtualMachineBulkOperationsClient_BulkCancelOperations_twoResponseWithPartiallySuccessfulResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineBulkOperationsClient().BulkCancelOperations(ctx, "example-rg", "eastus", armbulkactions.CancelOperationsContent{
		OperationIDs: []*string{
			to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			to.Ptr("dddddddd-dddd-dddd-dddd-dddddddddddd"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.VirtualMachineBulkOperationsClientBulkCancelOperationsResponse{
	// 	CancelOperationsResponse: armbulkactions.CancelOperationsResponse{
	// 		Results: []*armbulkactions.ResourceOperation{
	// 			{
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
	// 					ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Compute/virtualMachines/bulk-vm-01"),
	// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
	// 					State: to.Ptr(armbulkactions.OperationStateCancelled),
	// 					CompletedAt: to.Ptr(time.Date(2026, time.August, 31, 18, 5, 0, 0, time.UTC)),
	// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
	// 						ErrorCode: to.Ptr("OperationCancelledByUser"),
	// 						ErrorDetails: to.Ptr("Operation: aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa was cancelled by the user."),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				ErrorCode: to.Ptr("OperationNotFound"),
	// 				Operation: &armbulkactions.ResourceOperationDetails{
	// 					OperationID: to.Ptr("dddddddd-dddd-dddd-dddd-dddddddddddd"),
	// 					State: to.Ptr(armbulkactions.OperationState("Unknown")),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
