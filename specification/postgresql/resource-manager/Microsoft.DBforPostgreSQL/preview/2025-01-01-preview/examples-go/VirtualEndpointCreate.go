package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/VirtualEndpointCreate.json
func ExampleVirtualEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualEndpointsClient().BeginCreate(ctx, "testrg", "pgtestsvc4", "pgVirtualEndpoint1", armpostgresqlflexibleservers.VirtualEndpointResource{
		Properties: &armpostgresqlflexibleservers.VirtualEndpointResourceProperties{
			EndpointType: to.Ptr(armpostgresqlflexibleservers.VirtualEndpointTypeReadWrite),
			Members: []*string{
				to.Ptr("testPrimary1")},
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
	// res.VirtualEndpointResource = armpostgresqlflexibleservers.VirtualEndpointResource{
	// 	Name: to.Ptr("pgVirtualEndpoint1"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualEndpoints"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4/virtualEndpoints/pgVirtualEndpoint1"),
	// 	Properties: &armpostgresqlflexibleservers.VirtualEndpointResourceProperties{
	// 		EndpointType: to.Ptr(armpostgresqlflexibleservers.VirtualEndpointTypeReadWrite),
	// 		Members: []*string{
	// 			to.Ptr("testPrimary1")},
	// 			VirtualEndpoints: []*string{
	// 				to.Ptr("pgVirtualEndpoint1.reader.postgres.database.azure.com"),
	// 				to.Ptr("pgVirtualEndpoint1.writer.postgres.database.azure.com")},
	// 			},
	// 		}
}
