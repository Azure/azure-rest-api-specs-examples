package armservicefabricmanagedclusters_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmanagedclusters/armservicefabricmanagedclusters"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/ApplicationTypeVersionPutOperation_example.json
func ExampleApplicationTypeVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabricmanagedclusters.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.ApplicationTypeVersionResource = armservicefabricmanagedclusters.ApplicationTypeVersionResource{
	// 	Name: to.Ptr("1.0"),
	// 	Type: to.Ptr("Microsoft.ServiceFabric/managedClusters/applicationTypes/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/managedclusters/myCluster/applicationTypes/myAppType/versions/1.0"),
	// 	Properties: &armservicefabricmanagedclusters.ApplicationTypeVersionResourceProperties{
	// 		AppPackageURL: to.Ptr("http://fakelink.test.com/MyAppType"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
