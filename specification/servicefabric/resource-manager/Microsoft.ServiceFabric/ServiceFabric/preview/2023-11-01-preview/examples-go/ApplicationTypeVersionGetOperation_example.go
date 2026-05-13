package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v3"
)

// Generated from example definition: 2023-11-01-preview/ApplicationTypeVersionGetOperation_example.json
func ExampleApplicationTypeVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationTypeVersionsClient().Get(ctx, "resRg", "myCluster", "myAppType", "1.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicefabric.ApplicationTypeVersionsClientGetResponse{
	// 	ApplicationTypeVersionResource: armservicefabric.ApplicationTypeVersionResource{
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("versions"),
	// 		Etag: to.Ptr("W/\"636462502176040417\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applicationTypes/myAppType/versions/1.0"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armservicefabric.ApplicationTypeVersionResourceProperties{
	// 			AppPackageURL: to.Ptr("http://fakelink.test.com/MyAppType"),
	// 			DefaultParameterList: map[string]*string{
	// 			},
	// 			ProvisioningState: to.Ptr("Updating"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
