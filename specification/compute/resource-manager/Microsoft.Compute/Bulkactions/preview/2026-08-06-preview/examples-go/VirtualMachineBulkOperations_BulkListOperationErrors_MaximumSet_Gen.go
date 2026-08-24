package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/VirtualMachineBulkOperations_BulkListOperationErrors_MaximumSet_Gen.json
func ExampleVirtualMachineBulkOperationsClient_NewBulkListOperationErrorsPager_virtualMachineBulkOperationsBulkListOperationErrorsExample() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineBulkOperationsClient().NewBulkListOperationErrorsPager("rgBulkactions", "useast2euap", nil)
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
		// page = armbulkactions.VirtualMachineBulkOperationsClientBulkListOperationErrorsResponse{
		// 	ListBulkOperationErrorsResponse: armbulkactions.ListBulkOperationErrorsResponse{
		// 		Value: []*armbulkactions.ResourceOperation{
		// 			{
		// 				ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
		// 				ErrorCode: to.Ptr("TestErrorCode"),
		// 				ErrorDetails: to.Ptr("Test error details"),
		// 				Operation: &armbulkactions.ResourceOperationDetails{
		// 					OperationID: to.Ptr("af449548-8e1a-4079-874e-2caa4ff783cc"),
		// 					ResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
		// 					OpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
		// 					SubscriptionID: to.Ptr("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2"),
		// 					Deadline: to.Ptr(time.Date(2026, time.June, 11, 19, 35, 45, 98000000, time.UTC)),
		// 					DeadlineType: to.Ptr(armbulkactions.DeadlineTypeInitiateAt),
		// 					State: to.Ptr(armbulkactions.OperationStatePendingScheduling),
		// 					Timezone: to.Ptr("UTC"),
		// 					ResourceOperationError: &armbulkactions.ResourceOperationError{
		// 						ErrorCode: to.Ptr("OperationCancelledByUser"),
		// 						ErrorDetails: to.Ptr("Operation was cancelled by the user."),
		// 					},
		// 					FallbackOperationInfo: &armbulkactions.FallbackOperationInfo{
		// 						LastOpType: to.Ptr(armbulkactions.ResourceOperationTypeStart),
		// 						Status: to.Ptr("succeeded"),
		// 						Error: &armbulkactions.ResourceOperationError{
		// 							ErrorCode: to.Ptr("TestErrorCode"),
		// 							ErrorDetails: to.Ptr("Test error details"),
		// 						},
		// 					},
		// 					CompletedAt: to.Ptr(time.Date(2026, time.June, 11, 19, 35, 45, 98000000, time.UTC)),
		// 					RetryPolicy: &armbulkactions.RetryPolicy{
		// 						RetryCount: to.Ptr[int32](2),
		// 						RetryWindowInMinutes: to.Ptr[int32](19),
		// 						OnFailureAction: to.Ptr(armbulkactions.ResourceOperationTypeUnknown),
		// 					},
		// 					ResourceNotificationDetails: &armbulkactions.ResourceNotificationDetails{
		// 						ResourceContext: to.Ptr(""),
		// 					},
		// 				},
		// 				VirtualMachineInfo: &armbulkactions.VirtualMachineInfo{
		// 					VMSize: to.Ptr("Standard_D2ads_v5"),
		// 					Zone: to.Ptr("1"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/locations/useast2euap/listBulkOperationErrors?api-version=2026-08-06-preview&continuationToken=abc123"),
		// 	},
		// }
	}
}
