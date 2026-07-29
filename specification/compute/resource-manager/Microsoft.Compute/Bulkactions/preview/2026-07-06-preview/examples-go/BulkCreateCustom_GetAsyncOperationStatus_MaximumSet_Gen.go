package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/BulkCreateCustom_GetAsyncOperationStatus_MaximumSet_Gen.json
func ExampleBulkCreateCustomClient_GetAsyncOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBulkCreateCustomClient().GetAsyncOperationStatus(ctx, "eastus", "f1ac145b-9d8b-417d-8101-9962d03c0904", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.BulkCreateCustomClientGetAsyncOperationStatusResponse{
	// 	OperationStatusResult: armbulkactions.OperationStatusResult{
	// 		ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom/f1ac145b-9d8b-417d-8101-9962d03c0904"),
	// 		Name: to.Ptr("f1ac145b-9d8b-417d-8101-9962d03c0904"),
	// 		Status: to.Ptr("Failed"),
	// 		PercentComplete: to.Ptr[float64](100),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:35:00Z"); return t}()),
	// 		Operations: []*armbulkactions.OperationStatusResult{
	// 			{
	// 				ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/providers/Microsoft.Compute/locations/eastus/bulkCreateCustom/9d8e1f2a-3b4c-4d5e-6f7a-8b9c0d1e2f3a"),
	// 				Name: to.Ptr("9d8e1f2a-3b4c-4d5e-6f7a-8b9c0d1e2f3a"),
	// 				ResourceID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/virtualMachines/bulkvm-payments-0"),
	// 				Status: to.Ptr("Failed"),
	// 				PercentComplete: to.Ptr[float64](100),
	// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:20:00Z"); return t}()),
	// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-08T18:32:00Z"); return t}()),
	// 				Operations: []*armbulkactions.OperationStatusResult{
	// 				},
	// 				Error: &armbulkactions.ErrorDetail{
	// 					Code: to.Ptr("OverconstrainedAllocationRequest"),
	// 					Message: to.Ptr("Allocation failed. The requested VM size Standard_D2s_v5 is not available in zone 1 of region eastus."),
	// 					Target: to.Ptr("bulkvm-payments-0"),
	// 					Details: []*armbulkactions.ErrorDetail{
	// 					},
	// 					AdditionalInfo: []*armbulkactions.ErrorAdditionalInfo{
	// 						{
	// 							Type: to.Ptr("AllocationFailureDetails"),
	// 							Info: map[string]any{
	// 								"availableZones": []any{
	// 									"2",
	// 									"3",
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Error: &armbulkactions.ErrorDetail{
	// 			Code: to.Ptr("PartialAllocationFailure"),
	// 			Message: to.Ptr("One or more virtual machines in the bulk operation failed to allocate."),
	// 			Target: to.Ptr("bulkCreateCustom"),
	// 			Details: []*armbulkactions.ErrorDetail{
	// 			},
	// 			AdditionalInfo: []*armbulkactions.ErrorAdditionalInfo{
	// 			},
	// 		},
	// 	},
	// }
}
