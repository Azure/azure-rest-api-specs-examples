package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/threatintelligence/DeleteThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
