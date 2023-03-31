package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/PrivateEndpointConnections_Delete.json
func ExamplePrivateEndpointConnectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionClient().Delete(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "custestpece80project3980pe.7e35576b-3df4-478e-9759-f64351cf4f43", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
