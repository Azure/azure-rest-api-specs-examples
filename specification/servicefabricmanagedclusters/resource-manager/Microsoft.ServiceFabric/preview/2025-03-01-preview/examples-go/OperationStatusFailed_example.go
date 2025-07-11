package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-03-01-preview/OperationStatusFailed_example.json
func ExampleOperationStatusClient_Get_errorResponseDescribingWhyTheOperationFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusClient().Get(ctx, "eastus", "00000000-0000-0000-0000-000000001234", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.OperationStatusClientGetResponse{
	// 	LongRunningOperationResult: &armservicefabricmanagedclusters.LongRunningOperationResult{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000001234"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-03T23:58:02.2501337Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-04T00:13:03.2790951Z"); return t}()),
	// 		PercentComplete: to.Ptr[float64](100),
	// 		Status: to.Ptr("Failed"),
	// 		Error: &armservicefabricmanagedclusters.ErrorModelError{
	// 			Code: to.Ptr("-2146233029"),
	// 			Message: to.Ptr("A task was canceled."),
	// 		},
	// 	},
	// }
}
