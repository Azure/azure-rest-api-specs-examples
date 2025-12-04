package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/VirtualEndpointsListByServer.json
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
	pager := clientFactory.NewVirtualEndpointsClient().NewListByServerPager("exampleresourcegroup", "exampleserver", nil)
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
		// page.VirtualEndpointsList = armpostgresqlflexibleservers.VirtualEndpointsList{
		// 	Value: []*armpostgresqlflexibleservers.VirtualEndpoint{
		// 		{
		// 			Name: to.Ptr("examplebasename"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualEndpoints"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/virtualEndpoints/examplebasename"),
		// 			Properties: &armpostgresqlflexibleservers.VirtualEndpointResourceProperties{
		// 				EndpointType: to.Ptr(armpostgresqlflexibleservers.VirtualEndpointTypeReadWrite),
		// 				Members: []*string{
		// 					to.Ptr("examplebasename")},
		// 					VirtualEndpoints: []*string{
		// 						to.Ptr("examplebasename.reader.postgres.database.azure.com"),
		// 						to.Ptr("examplebasename.writer.postgres.database.azure.com")},
		// 					},
		// 			}},
		// 		}
	}
}
