package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/VirtualEndpointsListByServer.json
func ExampleVirtualEndpointsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualEndpointsClient().NewListByServerPager("testrg", "pgtestsvc4", nil)
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
		// page.VirtualEndpointsListResult = armpostgresqlflexibleservers.VirtualEndpointsListResult{
		// 	Value: []*armpostgresqlflexibleservers.VirtualEndpointResource{
		// 		{
		// 			Name: to.Ptr("pgVirtualEndpoint1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualEndpoints"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4/virtualEndpoints/pgVirtualEndpoint1"),
		// 			Properties: &armpostgresqlflexibleservers.VirtualEndpointResourceProperties{
		// 				EndpointType: to.Ptr(armpostgresqlflexibleservers.VirtualEndpointTypeReadWrite),
		// 				Members: []*string{
		// 					to.Ptr("testReplica1")},
		// 					VirtualEndpoints: []*string{
		// 						to.Ptr("pgVirtualEndpoint1.reader.postgres.database.azure.com"),
		// 						to.Ptr("pgVirtualEndpoint1.writer.postgres.database.azure.com")},
		// 					},
		// 			}},
		// 		}
	}
}
