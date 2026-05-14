package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ServiceListOperation_example.json
func ExampleServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicesClient().NewListPager("resRg", "myCluster", "myApp", nil)
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
		// page = armservicefabric.ServicesClientListResponse{
		// 	ServiceResourceList: armservicefabric.ServiceResourceList{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabric.ServiceResource{
		// 			{
		// 				Name: to.Ptr("myCluster"),
		// 				Type: to.Ptr("services"),
		// 				Etag: to.Ptr("W/\"636462502183671257\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp/services/myService"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armservicefabric.StatelessServiceProperties{
		// 					InstanceCount: to.Ptr[int32](1),
		// 					PartitionDescription: &armservicefabric.SingletonPartitionSchemeDescription{
		// 						PartitionScheme: to.Ptr(armservicefabric.PartitionSchemeSingleton),
		// 					},
		// 					ProvisioningState: to.Ptr("Updating"),
		// 					ServiceKind: to.Ptr(armservicefabric.ServiceKindStateless),
		// 					ServiceLoadMetrics: []*armservicefabric.ServiceLoadMetricDescription{
		// 						{
		// 							Name: to.Ptr("metric1"),
		// 							Weight: to.Ptr(armservicefabric.ServiceLoadMetricWeightLow),
		// 						},
		// 					},
		// 					ServicePackageActivationMode: to.Ptr(armservicefabric.ArmServicePackageActivationModeSharedProcess),
		// 					ServiceTypeName: to.Ptr("myServiceType"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
