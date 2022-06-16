package armengagementfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/engagementfabric/resource-manager/Microsoft.EngagementFabric/preview/2018-09-01/examples/AccountsDeleteExample.json
func ExampleAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armengagementfabric.NewAccountsClient("EDBF0095-A524-4A84-95FB-F72DA41AA6A1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"ExampleRg",
		"ExampleAccount",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
