package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-04-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_CreateIndicator() {
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
	res, err := client.CreateIndicator(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armsecurityinsights.ThreatIntelligenceIndicatorModel{
			Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
			Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
				Description:        to.Ptr("<description>"),
				Confidence:         to.Ptr[int32](78),
				CreatedByRef:       to.Ptr("<created-by-ref>"),
				DisplayName:        to.Ptr("<display-name>"),
				ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{},
				GranularMarkings:   []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{},
				KillChainPhases:    []*armsecurityinsights.ThreatIntelligenceKillChainPhase{},
				Labels:             []*string{},
				Modified:           to.Ptr("<modified>"),
				Pattern:            to.Ptr("<pattern>"),
				PatternType:        to.Ptr("<pattern-type>"),
				Revoked:            to.Ptr(false),
				Source:             to.Ptr("<source>"),
				ThreatIntelligenceTags: []*string{
					to.Ptr("new schema")},
				ThreatTypes: []*string{
					to.Ptr("compromised")},
				ValidFrom:  to.Ptr("<valid-from>"),
				ValidUntil: to.Ptr("<valid-until>"),
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
