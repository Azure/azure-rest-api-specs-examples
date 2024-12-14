package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/ServicePatchOperation_example.json
func ExampleServicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Update(ctx, "resRg", "myCluster", "myApp", "myService", armservicefabricmanagedclusters.ServiceUpdateParameters{
		Tags: map[string]*string{
			"a": to.Ptr("b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armservicefabricmanagedclusters.ServiceResource{
	// 	Name: to.Ptr("myService"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applications/services"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService"),
	// 	Tags: map[string]*string{
	// 		"a": to.Ptr("b"),
	// 	},
	// 	Properties: &armservicefabricmanagedclusters.StatelessServiceProperties{
	// 		CorrelationScheme: []*armservicefabricmanagedclusters.ServiceCorrelation{
	// 			{
	// 				Scheme: to.Ptr(armservicefabricmanagedclusters.ServiceCorrelationSchemeAlignedAffinity),
	// 				ServiceName: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService1"),
	// 		}},
	// 		DefaultMoveCost: to.Ptr(armservicefabricmanagedclusters.MoveCostMedium),
	// 		PlacementConstraints: to.Ptr("NodeType==frontend"),
	// 		ScalingPolicies: []*armservicefabricmanagedclusters.ScalingPolicy{
	// 			{
	// 				ScalingMechanism: &armservicefabricmanagedclusters.PartitionInstanceCountScaleMechanism{
	// 					Kind: to.Ptr(armservicefabricmanagedclusters.ServiceScalingMechanismKindScalePartitionInstanceCount),
	// 					MaxInstanceCount: to.Ptr[int32](9),
	// 					MinInstanceCount: to.Ptr[int32](3),
	// 					ScaleIncrement: to.Ptr[int32](2),
	// 				},
	// 				ScalingTrigger: &armservicefabricmanagedclusters.AveragePartitionLoadScalingTrigger{
	// 					Kind: to.Ptr(armservicefabricmanagedclusters.ServiceScalingTriggerKindAveragePartitionLoadTrigger),
	// 					LowerLoadThreshold: to.Ptr[float64](2),
	// 					MetricName: to.Ptr("metricName"),
	// 					ScaleInterval: to.Ptr("00:01:00"),
	// 					UpperLoadThreshold: to.Ptr[float64](8),
	// 				},
	// 		}},
	// 		ServiceLoadMetrics: []*armservicefabricmanagedclusters.ServiceLoadMetric{
	// 			{
	// 				Name: to.Ptr("metric1"),
	// 				DefaultLoad: to.Ptr[int32](3),
	// 				Weight: to.Ptr(armservicefabricmanagedclusters.ServiceLoadMetricWeightLow),
	// 		}},
	// 		ServicePlacementPolicies: []armservicefabricmanagedclusters.ServicePlacementPolicyClassification{
	// 			&armservicefabricmanagedclusters.ServicePlacementNonPartiallyPlaceServicePolicy{
	// 				Type: to.Ptr(armservicefabricmanagedclusters.ServicePlacementPolicyTypeNonPartiallyPlaceService),
	// 		}},
	// 		PartitionDescription: &armservicefabricmanagedclusters.SingletonPartitionScheme{
	// 			PartitionScheme: to.Ptr(armservicefabricmanagedclusters.PartitionSchemeSingleton),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceKind: to.Ptr(armservicefabricmanagedclusters.ServiceKindStateless),
	// 		ServicePackageActivationMode: to.Ptr(armservicefabricmanagedclusters.ServicePackageActivationModeSharedProcess),
	// 		ServiceTypeName: to.Ptr("myServiceType"),
	// 		InstanceCount: to.Ptr[int32](5),
	// 		MinInstanceCount: to.Ptr[int32](3),
	// 		MinInstancePercentage: to.Ptr[int32](30),
	// 	},
	// }
}
