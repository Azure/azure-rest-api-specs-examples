Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.3.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_max.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<application-name>",
		"<service-name>",
		armservicefabric.ServiceResource{
			Tags: map[string]*string{},
			Properties: &armservicefabric.StatelessServiceProperties{
				CorrelationScheme: []*armservicefabric.ServiceCorrelationDescription{
					{
						Scheme:      armservicefabric.ServiceCorrelationScheme("Affinity").ToPtr(),
						ServiceName: to.StringPtr("<service-name>"),
					}},
				DefaultMoveCost:      armservicefabric.MoveCost("Medium").ToPtr(),
				PlacementConstraints: to.StringPtr("<placement-constraints>"),
				ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
					{
						Name:   to.StringPtr("<name>"),
						Weight: armservicefabric.ServiceLoadMetricWeight("Low").ToPtr(),
					}},
				ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{},
				PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
					PartitionScheme: armservicefabric.PartitionScheme("Singleton").ToPtr(),
				},
				ServiceDNSName:               to.StringPtr("<service-dnsname>"),
				ServiceKind:                  armservicefabric.ServiceKind("Stateless").ToPtr(),
				ServicePackageActivationMode: armservicefabric.ArmServicePackageActivationMode("SharedProcess").ToPtr(),
				ServiceTypeName:              to.StringPtr("<service-type-name>"),
				InstanceCount:                to.Int32Ptr(5),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
