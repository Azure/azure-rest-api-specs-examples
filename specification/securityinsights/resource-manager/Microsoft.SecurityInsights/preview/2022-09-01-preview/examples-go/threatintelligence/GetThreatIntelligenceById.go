package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/GetThreatIntelligenceById.json
func ExampleThreatIntelligenceIndicatorClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewThreatIntelligenceIndicatorClient().Get(ctx, "myRg", "myWorkspace", "e16ef847-962e-d7b6-9c8b-a33e4bd30e47", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.ThreatIntelligenceIndicatorClientGetResponse{
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
	// 			LastUpdatedTimeUTC: to.Ptr("2021-04-15T20:18:49.2259902Z"),
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
