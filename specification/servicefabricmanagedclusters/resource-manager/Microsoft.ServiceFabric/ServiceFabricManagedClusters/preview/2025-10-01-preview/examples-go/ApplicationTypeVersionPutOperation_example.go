package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: 2025-10-01-preview/ApplicationTypeVersionPutOperation_example.json
func ExampleApplicationTypeVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationTypeVersionsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myAppType", "1.0", armservicefabricmanagedclusters.ApplicationTypeVersionResource{
		Location: to.Ptr("eastus"),
		Properties: &armservicefabricmanagedclusters.ApplicationTypeVersionResourceProperties{
			AppPackageURL: to.Ptr("http://fakelink.test.com/MyAppType"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabricmanagedclusters.ApplicationTypeVersionsClientCreateOrUpdateResponse{
	// 	ApplicationTypeVersionResource: &armservicefabricmanagedclusters.ApplicationTypeVersionResource{
	// 		Name: to.Ptr("1.0"),
	// 		Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applicationTypes/versions"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
	// 		Properties: &armservicefabricmanagedclusters.ApplicationTypeVersionResourceProperties{
	// 			AppPackageURL: to.Ptr("http://fakelink.test.com/MyAppType"),
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 	},
	// }
}
