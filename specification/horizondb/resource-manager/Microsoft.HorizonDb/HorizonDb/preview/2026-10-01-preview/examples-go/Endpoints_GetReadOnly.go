package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Endpoints_GetReadOnly.json
func ExampleEndpointsClient_Get_getAReadOnlyHorizonDbPoolEndpoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEndpointsClient().Get(ctx, "example-resource-group", "example-cluster", "example-pool-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.EndpointsClientGetResponse{
	// 	Endpoint: armhorizondb.Endpoint{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/endpoints/example-pool-01"),
	// 		Name: to.Ptr("example-pool-01"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/endpoints"),
	// 		Properties: &armhorizondb.EndpointProperties{
	// 			Scope: to.Ptr(armhorizondb.EndpointScopePool),
	// 			AssociatedResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01"),
	// 			FullyQualifiedDomainName: to.Ptr("example-pool-01.example-cluster.ffffffffffff.eastus2.horizondb.azure.com"),
	// 			Role: to.Ptr(armhorizondb.EndpointRoleReadOnly),
	// 		},
	// 	},
	// }
}
