```go
package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_max.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicefabric.NewServicesClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"resRg",
		"myCluster",
		"myApp",
		"myService",
		armservicefabric.ServiceResource{
			Tags: map[string]*string{},
			Properties: &armservicefabric.StatelessServiceProperties{
				CorrelationScheme: []*armservicefabric.ServiceCorrelationDescription{
					{
						Scheme:      to.Ptr(armservicefabric.ServiceCorrelationSchemeAffinity),
						ServiceName: to.Ptr("fabric:/app1/app1~svc1"),
					}},
				DefaultMoveCost:      to.Ptr(armservicefabric.MoveCostMedium),
				PlacementConstraints: to.Ptr("NodeType==frontend"),
				ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
					{
						Name:   to.Ptr("metric1"),
						Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
					}},
				ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{},
				PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
					PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
				},
				ServiceDNSName:               to.Ptr("my.service.dns"),
				ServiceKind:                  to.Ptr(armservicefabric.ServiceKindStateless),
				ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
				ServiceTypeName:              to.Ptr("myServiceType"),
				InstanceCount:                to.Ptr[int32](5),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv1.0.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.
