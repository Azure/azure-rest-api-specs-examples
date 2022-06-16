package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListResourceSkus.json
func ExampleKustoPoolsClient_ListSKUsByResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewKustoPoolsClient("<subscription-id>", cred, nil)
	_, err = client.ListSKUsByResource(ctx,
		"<workspace-name>",
		"<kusto-pool-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
