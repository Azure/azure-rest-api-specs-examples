package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_Stop.json
func ExampleSAPVirtualInstancesClient_BeginStop_sapVirtualInstancesStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginStop(ctx, "test-rg", "X00", &armworkloadssapvirtualinstance.SAPVirtualInstancesClientBeginStopOptions{
		Body: &armworkloadssapvirtualinstance.StopRequest{
			SoftStopTimeoutSeconds: to.Ptr[int64](0),
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientStopResponse{
	// 	OperationStatusResult: &armworkloadssapvirtualinstance.OperationStatusResult{
	// 		Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:38:07.000Z"); return t}()),
	// 		ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/operations/locations/centralus/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T03:36:07.000Z"); return t}()),
	// 		Status: to.Ptr("Succeeded"),
	// 	},
	// }
}
