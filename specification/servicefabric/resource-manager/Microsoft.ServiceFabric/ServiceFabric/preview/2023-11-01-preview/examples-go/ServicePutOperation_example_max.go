package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ServicePutOperation_example_max.json
func ExampleServicesClient_BeginCreateOrUpdate_putAServiceWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServicesClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabric.ServiceResource{
		Properties: &armservicefabric.StatelessServiceProperties{
			CorrelationScheme: []*armservicefabric.ServiceCorrelationDescription{
				{
					Scheme:      to.Ptr(armservicefabric.ServiceCorrelationSchemeAffinity),
					ServiceName: to.Ptr("fabric:/app1/app1~svc1"),
				},
			},
			DefaultMoveCost: to.Ptr(armservicefabric.MoveCostMedium),
			InstanceCount:   to.Ptr[int32](5),
			PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
				PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
			},
			PlacementConstraints: to.Ptr("NodeType==frontend"),
			ServiceDNSName:       to.Ptr("my.service.dns"),
			ServiceKind:          to.Ptr(armservicefabric.ServiceKindStateless),
			ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
				{
					Name:   to.Ptr("metric1"),
					Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
				},
			},
			ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
			ServicePlacementPolicies:     []armservicefabric.ServicePlacementPolicyDescriptionClassification{},
			ServiceTypeName:              to.Ptr("myServiceType"),
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
