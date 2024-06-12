package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
func ExampleThreatIntelligenceIndicatorClient_ReplaceTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewThreatIntelligenceIndicatorClient().ReplaceTags(ctx, "myRg", "myWorkspace", "d9cd6f0b-96b9-3984-17cd-a779d1e15a93", armsecurityinsights.ThreatIntelligenceIndicatorModel{
		Etag: to.Ptr("\"0000262c-0000-0800-0000-5e9767060000\""),
		Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
		Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
			ThreatIntelligenceTags: []*string{
				to.Ptr("patching tags")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.ThreatIntelligenceIndicatorClientReplaceTagsResponse{
	// 	                            ThreatIntelligenceInformationClassification: &armsecurityinsights.ThreatIntelligenceIndicatorModel{
	// 		Name: to.Ptr("e16ef847-962e-d7b6-9c8b-a33e4bd30e47"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/ThreatIntelligence"),
	// 		ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/ThreatIntelligence/e16ef847-962e-d7b6-9c8b-a33e4bd30e47"),
	// 		Etag: to.Ptr("\"00002a2c-0000-0800-0000-5e97683b0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.ThreatIntelligenceResourceKindEnumIndicator),
	// 		Properties: &armsecurityinsights.ThreatIntelligenceIndicatorProperties{
	// 			Description: to.Ptr("debugging indicators"),
	// 			Confidence: to.Ptr[int32](78),
	// 			Created: to.Ptr("2021-04-15T19:51:17.1050923Z"),
	// 			CreatedByRef: to.Ptr("aztestConnectors@dataconnector.ccsctp.net"),
	// 			DisplayName: to.Ptr("updated indicator"),
	// 			ExternalID: to.Ptr("indicator--73be1729-babb-4348-a6c4-94621cae2530"),
	// 			ExternalReferences: []*armsecurityinsights.ThreatIntelligenceExternalReference{
	// 			},
	// 			GranularMarkings: []*armsecurityinsights.ThreatIntelligenceGranularMarkingModel{
	// 			},
	// 			KillChainPhases: []*armsecurityinsights.ThreatIntelligenceKillChainPhase{
	// 			},
	// 			LastUpdatedTimeUTC: to.Ptr("2021-04-15T19:56:08.828946Z"),
	// 			Pattern: to.Ptr("[url:value = 'https://abc.com']"),
	// 			PatternType: to.Ptr("url"),
	// 			Revoked: to.Ptr(false),
	// 			Source: to.Ptr("Azure Sentinel"),
	// 			ThreatIntelligenceTags: []*string{
	// 				to.Ptr("patching tags")},
	// 				ThreatTypes: []*string{
	// 					to.Ptr("compromised")},
	// 					ValidFrom: to.Ptr("2021-04-15T17:44:00.114052Z"),
	// 				},
	// 			},
	// 			                        }
}
