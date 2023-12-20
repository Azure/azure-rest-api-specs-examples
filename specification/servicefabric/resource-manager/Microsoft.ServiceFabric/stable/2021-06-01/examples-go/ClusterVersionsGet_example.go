package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterVersionsGet_example.json
func ExampleClusterVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterVersionsClient().Get(ctx, "eastus", "6.1.480.9494", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterCodeVersionsListResult = armservicefabric.ClusterCodeVersionsListResult{
	// 	Value: []*armservicefabric.ClusterCodeVersionsResult{
	// 		{
	// 			Name: to.Ptr("6.1.480.9494"),
	// 			Type: to.Ptr("Microsoft.ServiceFabric/locations/environments/clusterVersions"),
	// 			ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.ServiceFabric/locations/eastus/environments/Windows/clusterVersions/6.1.480.9494"),
	// 			Properties: &armservicefabric.ClusterVersionDetails{
	// 				CodeVersion: to.Ptr("6.1.480.9494"),
	// 				Environment: to.Ptr(armservicefabric.ClusterEnvironmentWindows),
	// 				SupportExpiryUTC: to.Ptr("2018-06-15T23:59:59.9999999"),
	// 			},
	// 	}},
	// }
}
