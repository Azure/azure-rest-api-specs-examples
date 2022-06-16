package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePatchOperation_example.json
func ExampleServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabric.NewServicesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"resRg",
		"myCluster",
		"myApp",
		"myService",
		armservicefabric.ServiceResourceUpdate{
			Tags: map[string]*string{},
			Properties: &armservicefabric.StatelessServiceUpdateProperties{
				ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
					{
						Name:   to.Ptr("metric1"),
						Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
					}},
				ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
