package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-08-06-preview/LaunchBulkInstancesOperation_GetOperationStatus_MaximumSet_Gen.json
func ExampleLaunchBulkInstancesOperationClient_GetOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLaunchBulkInstancesOperationClient().GetOperationStatus(ctx, "useast2euap", "8596407e-8834-4a62-8d3c-9231af92d785", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbulkactions.LaunchBulkInstancesOperationClientGetOperationStatusResponse{
	// 	OperationStatusResult: armbulkactions.OperationStatusResult{
	// 		Status: to.Ptr("Succeeded"),
	// 		StartTime: to.Ptr(time.Date(2025, time.November, 25, 22, 19, 26, 906000000, time.UTC)),
	// 		EndTime: to.Ptr(time.Date(2025, time.November, 25, 22, 19, 26, 906000000, time.UTC)),
	// 		Name: to.Ptr("2a3fce8e-874c-45f4-9d27-1a804f3b7a0f"),
	// 	},
	// }
}
