package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_CreateIndicator() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewThreatIntelligenceIndicatorClient().CreateIndicator(ctx, "myRg", "myWorkspace", armsecurityinsights.ThreatIntelligenceIndicatorModel{
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
			ValidFrom:  to.Ptr("2021-09-15T17:44:00.114052Z"),
			ValidUntil: to.Ptr(""),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.ThreatIntelligenceIndicatorClientCreateIndicatorResponse{
	// 	                            ThreatIntelligenceInformationClassification: &armsecurityinsights.ThreatIntelligenceIndicatorModel{
	// 		Name: to.Ptr("180105c7-a28d-b1a2-4a78-234f6ec80fd6"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/ThreatIntelligence"),
	// 		ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/ThreatIntelligence/180105c7-a28d-b1a2-4a78-234f6ec80fd6"),
	// 		Etag: to.Ptr("\"0000322c-0000-0800-0000-5e976c960000\""),
	// 		Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
	// 		Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
	// 			Description: to.Ptr("debugging indicators"),
	// 			Confidence: to.Ptr[int32](78),
	// 			Created: to.Ptr("2021-09-15T20:20:38.6160949Z"),
	// 			CreatedByRef: to.Ptr("contoso@contoso.com"),
	// 			DisplayName: to.Ptr("new schema"),
	// 			ExternalID: to.Ptr("indicator--a2b6a95e-2108-4a38-bd49-ef95811bbcd7"),
	// 			ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{
	// 			},
	// 			GranularMarkings: []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{
	// 			},
	// 			KillChainPhases: []*armsecurityinsights.ThreatIntelligenceKillChainPhase{
	// 			},
	// 			LastUpdatedTimeUTC: to.Ptr("2020-04-15T20:20:38.6161887Z"),
	// 			Pattern: to.Ptr("[url:value = 'https://www.contoso.com']"),
	// 			PatternType: to.Ptr("url"),
	// 			Revoked: to.Ptr(false),
	// 			Source: to.Ptr("Azure Sentinel"),
	// 			ThreatIntelligenceTags: []*string{
	// 				to.Ptr("new schema")},
	// 				ThreatTypes: []*string{
	// 					to.Ptr("compromised")},
	// 					ValidFrom: to.Ptr("2021-09-15T17:44:00.114052Z"),
	// 				},
	// 			},
	// 			                        }
}
