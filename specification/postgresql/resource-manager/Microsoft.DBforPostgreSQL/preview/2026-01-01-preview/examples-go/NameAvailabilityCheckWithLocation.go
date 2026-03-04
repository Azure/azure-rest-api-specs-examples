package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v6"
)

// Generated from example definition: 2026-01-01-preview/NameAvailabilityCheckWithLocation.json
func ExampleNameAvailabilityClient_CheckWithLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNameAvailabilityClient().CheckWithLocation(ctx, "eastus", armpostgresqlflexibleservers.CheckNameAvailabilityRequest{
		Name: to.Ptr("exampleserver"),
		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpostgresqlflexibleservers.NameAvailabilityClientCheckWithLocationResponse{
	// 	NameAvailabilityModel: &armpostgresqlflexibleservers.NameAvailabilityModel{
	// 		Name: to.Ptr("exampleserver"),
	// 		Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers"),
	// 		Message: to.Ptr(""),
	// 		NameAvailable: to.Ptr(true),
	// 	},
	// }
}
