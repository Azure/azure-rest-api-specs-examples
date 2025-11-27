package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationTypeVersionListOperation_example.json
func ExampleApplicationTypeVersionsClient_NewListByApplicationTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationTypeVersionsClient().NewListByApplicationTypesPager("resRg", "myCluster", "myAppType", nil)
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
		// page = armservicefabricmanagedclusters.ApplicationTypeVersionsClientListByApplicationTypesResponse{
		// 	ApplicationTypeVersionResourceList: armservicefabricmanagedclusters.ApplicationTypeVersionResourceList{
		// 		NextLink: to.Ptr("http://examplelink.com"),
		// 		Value: []*armservicefabricmanagedclusters.ApplicationTypeVersionResource{
		// 			{
		// 				Name: to.Ptr("1.0"),
		// 				Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applicationTypes/versions"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
		// 				Properties: &armservicefabricmanagedclusters.ApplicationTypeVersionResourceProperties{
		// 					AppPackageURL: to.Ptr("http://fakelink.test.com/MyAppType"),
		// 					ProvisioningState: to.Ptr("Updating"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
