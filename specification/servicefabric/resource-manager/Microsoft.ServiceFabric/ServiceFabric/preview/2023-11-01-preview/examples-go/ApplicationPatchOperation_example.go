package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ApplicationPatchOperation_example.json
func ExampleApplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResourceUpdate{
		Properties: &armservicefabric.ApplicationResourceUpdateProperties{
			Metrics: []*armservicefabric.ApplicationMetricDescription{
				{
					Name:                     to.Ptr("metric1"),
					MaximumCapacity:          to.Ptr[int64](3),
					ReservationCapacity:      to.Ptr[int64](1),
					TotalApplicationCapacity: to.Ptr[int64](5),
				},
			},
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
