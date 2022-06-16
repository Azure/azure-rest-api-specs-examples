package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredisenterprise.NewPrivateEndpointConnectionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"rg1",
		"cache1",
		"pectest01",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
