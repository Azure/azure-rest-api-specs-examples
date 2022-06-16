package armblueprint_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blueprint/armblueprint"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Delete.json
func ExampleAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armblueprint.NewAssignmentsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"managementGroups/ContosoOnlineGroup",
		"assignSimpleBlueprint",
		&armblueprint.AssignmentsClientDeleteOptions{DeleteBehavior: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
