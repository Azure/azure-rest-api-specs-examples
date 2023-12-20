package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPatchOperation_example.json
func ExampleApplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResourceUpdate{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceUpdateProperties{
			Metrics: []*armservicefabric.ApplicationMetricDescription{
				{
					Name:                     to.Ptr("metric1"),
					MaximumCapacity:          to.Ptr[int64](3),
					ReservationCapacity:      to.Ptr[int64](1),
					TotalApplicationCapacity: to.Ptr[int64](5),
				}},
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
