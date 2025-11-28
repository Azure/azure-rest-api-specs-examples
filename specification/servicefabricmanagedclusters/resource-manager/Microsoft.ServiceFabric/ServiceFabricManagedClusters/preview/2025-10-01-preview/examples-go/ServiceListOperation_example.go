package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ServiceListOperation_example.json
func ExampleServicesClient_NewListByApplicationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armservicefabricmanagedclusters.ServicesClientListByApplicationsResponse{
		// 	ServiceResourceList: armservicefabricmanagedclusters.ServiceResourceList{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabricmanagedclusters.ServiceResource{
		// 			{
		// 				Name: to.Ptr("myService"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applications/services"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applications/myApp/services/myService"),
		// 				Properties: &armservicefabricmanagedclusters.StatelessServiceProperties{
		// 					InstanceCount: to.Ptr[int32](1),
		// 					PartitionDescription: &armservicefabricmanagedclusters.SingletonPartitionScheme{
		// 						PartitionScheme: to.Ptr(armservicefabricmanagedclusters.PartitionSchemeSingleton),
		// 					},
		// 					ProvisioningState: to.Ptr("Updating"),
		// 					ServiceKind: to.Ptr(armservicefabricmanagedclusters.ServiceKindStateless),
		// 					ServiceLoadMetrics: []*armservicefabricmanagedclusters.ServiceLoadMetric{
		// 						{
		// 							Name: to.Ptr("metric1"),
		// 							Weight: to.Ptr(armservicefabricmanagedclusters.ServiceLoadMetricWeightLow),
		// 						},
		// 					},
		// 					ServicePackageActivationMode: to.Ptr(armservicefabricmanagedclusters.ServicePackageActivationModeSharedProcess),
		// 					ServiceTypeName: to.Ptr("myServiceType"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
