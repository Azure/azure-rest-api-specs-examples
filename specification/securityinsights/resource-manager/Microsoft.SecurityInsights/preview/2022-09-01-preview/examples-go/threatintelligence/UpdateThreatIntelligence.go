package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/UpdateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewThreatIntelligenceIndicatorClient("bd794837-4d29-4647-9105-6339bfdb4e6a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "myRg", "myWorkspace", "d9cd6f0b-96b9-3984-17cd-a779d1e15a93", armsecurityinsights.ThreatIntelligenceIndicatorModel{
		Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
		Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
			Description:        to.Ptr("debugging indicators"),
			Confidence:         to.Ptr[int32](78),
			CreatedByRef:       to.Ptr("contoso@contoso.com"),
			DisplayName:        to.Ptr("new schema"),
			ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{},
			GranularMarkings:   []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{},
			KillChainPhases:    []*armsecurityinsights.ThreatIntelligenceKillChainPhase{},
			Labels:             []*string{},
			Modified:           to.Ptr(""),
			Pattern:            to.Ptr("[url:value = 'https://www.contoso.com']"),
			PatternType:        to.Ptr("url"),
			Revoked:            to.Ptr(false),
			Source:             to.Ptr("Azure Sentinel"),
			ThreatIntelligenceTags: []*string{
				to.Ptr("new schema")},
			ThreatTypes: []*string{
				to.Ptr("compromised")},
			ValidFrom:  to.Ptr("2020-04-15T17:44:00.114052Z"),
			ValidUntil: to.Ptr(""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
