package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2026-05-01-preview/OperationStatusSucceeded_example.json
func ExampleOperationStatusClient_Get_okTheRequestHasSucceeded() {
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
	// 	LongRunningOperationResult: armservicefabricmanagedclusters.LongRunningOperationResult{
	// 		Name: to.Ptr("00000000-0000-0000-0000-000000001234"),
	// 		StartTime: to.Ptr(time.Date(2022, time.January, 3, 23, 58, 2, 250133700, time.UTC)),
	// 		EndTime: to.Ptr(time.Date(2022, time.January, 4, 0, 13, 3, 279095100, time.UTC)),
	// 		PercentComplete: to.Ptr[float64](100),
	// 		Status: to.Ptr("Succeeded"),
	// 	},
	// }
}
