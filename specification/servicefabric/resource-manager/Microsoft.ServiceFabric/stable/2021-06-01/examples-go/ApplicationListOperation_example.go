package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationListOperation_example.json
func ExampleApplicationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().List(ctx, "resRg", "myCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationResourceList = armservicefabric.ApplicationResourceList{
	// 	Value: []*armservicefabric.ApplicationResource{
	// 		{
	// 			Name: to.Ptr("myCluster"),
	// 			Type: to.Ptr("applications"),
	// 			Etag: to.Ptr("W/\"636462502180261858\""),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp"),
	// 			Location: to.Ptr("eastus"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Properties: &armservicefabric.ApplicationResourceProperties{
	// 				Metrics: []*armservicefabric.ApplicationMetricDescription{
	// 					{
	// 						Name: to.Ptr("metric1"),
	// 						MaximumCapacity: to.Ptr[int64](3),
	// 						ReservationCapacity: to.Ptr[int64](1),
	// 						TotalApplicationCapacity: to.Ptr[int64](5),
	// 				}},
	// 				RemoveApplicationCapacity: to.Ptr(false),
	// 				TypeVersion: to.Ptr("1.0"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				TypeName: to.Ptr("myAppType"),
	// 			},
	// 	}},
	// }
}
