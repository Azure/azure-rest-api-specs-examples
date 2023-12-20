package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServiceGetOperation_example.json
func ExampleServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicesClient().Get(ctx, "resRg", "myCluster", "myApp", "myService", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceResource = armservicefabric.ServiceResource{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("services"),
	// 	Etag: to.Ptr("W/\"636462502183671258\""),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp/services/myService"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabric.StatelessServiceProperties{
	// 		DefaultMoveCost: to.Ptr(armservicefabric.MoveCostMedium),
	// 		PlacementConstraints: to.Ptr("NodeType==frontend"),
	// 		ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
	// 			{
	// 				Name: to.Ptr("metric1"),
	// 				Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
	// 		}},
	// 		ServicePlacementPolicies: []armservicefabric.ServicePlacementPolicyDescriptionClassification{
	// 		},
	// 		PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
	// 			PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
	// 		ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
	// 		ServiceTypeName: to.Ptr("myServiceType"),
	// 		InstanceCount: to.Ptr[int32](5),
	// 	},
	// }
}
