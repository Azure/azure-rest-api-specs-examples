package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2021-08-01/examples/RedisEnterpriseDeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredisenterprise.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<private-endpoint-connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
