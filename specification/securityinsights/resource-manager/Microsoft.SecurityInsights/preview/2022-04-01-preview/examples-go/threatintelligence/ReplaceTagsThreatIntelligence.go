package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_ReplaceTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ReplaceTags(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Etag: to.Ptr("<etag>"),
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				ThreatIntelligenceTags: []*string{
					to.Ptr("patching tags")},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
