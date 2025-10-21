package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-06-01-preview/ApplicationTypeNameListOperation_example.json
func ExampleApplicationTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationTypesClient().NewListPager("resRg", "myCluster", nil)
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
		// page = armservicefabricmanagedclusters.ApplicationTypesClientListResponse{
		// 	ApplicationTypeResourceList: armservicefabricmanagedclusters.ApplicationTypeResourceList{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabricmanagedclusters.ApplicationTypeResource{
		// 			{
		// 				Name: to.Ptr("myAppType"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applicationTypes"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType"),
		// 				Properties: &armservicefabricmanagedclusters.ApplicationTypeResourceProperties{
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
