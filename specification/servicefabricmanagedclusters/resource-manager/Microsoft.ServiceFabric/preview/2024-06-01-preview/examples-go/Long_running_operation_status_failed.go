package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/Long_running_operation_status_failed.json
func ExampleOperationStatusClient_Get_getFailedOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.LongRunningOperationResult = armservicefabricmanagedclusters.LongRunningOperationResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000001234"),
	// 	EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-04T00:13:03.279Z"); return t}()),
	// 	Error: &armservicefabricmanagedclusters.ErrorModelError{
	// 		Code: to.Ptr("-2146233029"),
	// 		Message: to.Ptr("A task was canceled."),
	// 	},
	// 	PercentComplete: to.Ptr[float64](100),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-03T23:58:02.250Z"); return t}()),
	// 	Status: to.Ptr("Failed"),
	// }
}
