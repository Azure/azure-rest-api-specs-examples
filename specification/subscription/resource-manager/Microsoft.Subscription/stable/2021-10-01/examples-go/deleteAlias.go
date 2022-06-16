package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/deleteAlias.json
func ExampleAliasClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewAliasClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"aliasForNewSub",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
