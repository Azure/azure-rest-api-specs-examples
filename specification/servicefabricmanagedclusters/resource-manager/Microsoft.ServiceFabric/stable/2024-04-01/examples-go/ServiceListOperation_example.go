package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a651ba25cda4eec698a3a4e35f867ecc2681d126/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ServiceListOperation_example.json
func ExampleServicesClient_NewListByApplicationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListByApplicationsPager("resRg", "myCluster", "myApp", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ServiceResourceList = armservicefabricmanagedclusters.ServiceResourceList{
		// 	Value: []*armservicefabricmanagedclusters.ServiceResource{
		// 		{
		// 			Name: to.Ptr("myService"),
		// 			Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applications/services"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService"),
		// 			Properties: &armservicefabricmanagedclusters.StatelessServiceProperties{
		// 				ServiceLoadMetrics: []*armservicefabricmanagedclusters.ServiceLoadMetric{
		// 					{
		// 						Name: to.Ptr("metric1"),
		// 						Weight: to.Ptr(armservicefabricmanagedclusters.ServiceLoadMetricWeightLow),
		// 				}},
		// 				PartitionDescription: &armservicefabricmanagedclusters.SingletonPartitionScheme{
		// 					PartitionScheme: to.Ptr(armservicefabricmanagedclusters.PartitionSchemeSingleton),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServiceKind: to.Ptr(armservicefabricmanagedclusters.ServiceKindStateless),
		// 				ServicePackageActivationMode: to.Ptr(armservicefabricmanagedclusters.ServicePackageActivationModeSharedProcess),
		// 				ServiceTypeName: to.Ptr("myServiceType"),
		// 				InstanceCount: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}
