package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ApplicationListOperation_example.json
func ExampleApplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationsClient().NewListPager("resRg", "myCluster", nil)
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
		// page = armservicefabric.ApplicationsClientListResponse{
		// 	ApplicationResourceList: armservicefabric.ApplicationResourceList{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabric.ApplicationResource{
		// 			{
		// 				Name: to.Ptr("myCluster"),
		// 				Type: to.Ptr("applications"),
		// 				Etag: to.Ptr("W/\"636462502180261858\""),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armservicefabric.ApplicationResourceProperties{
		// 					Metrics: []*armservicefabric.ApplicationMetricDescription{
		// 						{
		// 							Name: to.Ptr("metric1"),
		// 							MaximumCapacity: to.Ptr[int64](3),
		// 							ReservationCapacity: to.Ptr[int64](1),
		// 							TotalApplicationCapacity: to.Ptr[int64](5),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Updating"),
		// 					RemoveApplicationCapacity: to.Ptr(false),
		// 					TypeName: to.Ptr("myAppType"),
		// 					TypeVersion: to.Ptr("1.0"),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
