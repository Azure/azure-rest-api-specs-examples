package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/VirtualEndpointsGet.json
func ExampleVirtualEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualEndpointsClient().Get(ctx, "exampleresourcegroup", "exampleserver", "examplebasename", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualEndpoint = armpostgresqlflexibleservers.VirtualEndpoint{
	// 	Name: to.Ptr("examplebasename"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualEndpoints"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/virtualEndpoints/examplebasename"),
	// 	Properties: &armpostgresqlflexibleservers.VirtualEndpointResourceProperties{
	// 		EndpointType: to.Ptr(armpostgresqlflexibleservers.VirtualEndpointTypeReadWrite),
	// 		Members: []*string{
	// 			to.Ptr("exampleprimaryserver")},
	// 			VirtualEndpoints: []*string{
	// 				to.Ptr("examplebasename.reader.postgres.database.azure.com"),
	// 				to.Ptr("examplebasename.writer.postgres.database.azure.com")},
	// 			},
	// 		}
}
