package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationGatewayAvailableSslOptionsGet.json
func ExampleApplicationGatewaysClient_ListAvailableSSLOptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewApplicationGatewaysClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ListAvailableSSLOptions(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
