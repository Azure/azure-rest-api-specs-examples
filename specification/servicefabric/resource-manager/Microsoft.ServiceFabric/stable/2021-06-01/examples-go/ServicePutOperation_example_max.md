Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.6.0/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_max.json
func ExampleServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armservicefabric.NewServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
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
						Scheme:      to.Ptr(armservicefabric.ServiceCorrelationSchemeAffinity),
						ServiceName: to.Ptr("<service-name>"),
					}},
				DefaultMoveCost:      to.Ptr(armservicefabric.MoveCostMedium),
				PlacementConstraints: to.Ptr("<placement-constraints>"),
				ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
					{
						Name:   to.Ptr("<name>"),
						Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
					}},
				ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{},
				PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
					PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
				},
				ServiceDNSName:               to.Ptr("<service-dnsname>"),
				ServiceKind:                  to.Ptr(armservicefabric.ServiceKindStateless),
				ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
				ServiceTypeName:              to.Ptr("<service-type-name>"),
				InstanceCount:                to.Ptr[int32](5),
			},
		},
		&armservicefabric.ServicesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
