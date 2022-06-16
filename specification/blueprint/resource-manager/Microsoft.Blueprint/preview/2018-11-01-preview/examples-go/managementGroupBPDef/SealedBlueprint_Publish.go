package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPDef/SealedBlueprint_Publish.json
func ExamplePublishedBlueprintsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewPublishedBlueprintsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Create(ctx,
		"providers/Microsoft.Management/managementGroups/ContosoOnlineGroup",
		"simpleBlueprint",
		"v2",
		&armblueprint.PublishedBlueprintsClientCreateOptions{PublishedBlueprint: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
