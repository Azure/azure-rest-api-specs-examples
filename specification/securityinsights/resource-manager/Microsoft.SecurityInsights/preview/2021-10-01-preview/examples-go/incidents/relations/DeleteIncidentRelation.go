package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/incidents/relations/DeleteIncidentRelation.json
func ExampleIncidentRelationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewIncidentRelationsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<incident-id>",
		"<relation-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
