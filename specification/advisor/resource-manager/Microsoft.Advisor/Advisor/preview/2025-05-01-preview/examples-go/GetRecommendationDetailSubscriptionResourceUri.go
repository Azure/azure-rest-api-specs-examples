package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/GetRecommendationDetailSubscriptionResourceUri.json
func ExampleRecommendationsClient_Get_getRecommendationDetailSubscriptionResourceUri() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecommendationsClient().Get(ctx, "subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset", "bd27ddc6-1312-4067-b4af-cbb45e32cfd7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.RecommendationsClientGetResponse{
	// 	ResourceRecommendationBase: armadvisor.ResourceRecommendationBase{
	// 		ID: to.Ptr("/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset/providers/Microsoft.Advisor/recommendations/bd27ddc6-1312-4067-b4af-cbb45e32cfd7"),
	// 		Name: to.Ptr("bd27ddc6-1312-4067-b4af-cbb45e32cfd7"),
	// 		Properties: &armadvisor.RecommendationProperties{
	// 			Category: to.Ptr(armadvisor.CategorySecurity),
	// 			Impact: to.Ptr(armadvisor.ImpactMedium),
	// 			ImpactedField: to.Ptr("Microsoft.Compute/availabilitysets"),
	// 			ImpactedValue: to.Ptr("armavset"),
	// 			LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T22:24:43.3216408Z"); return t}()),
	// 			Risk: to.Ptr(armadvisor.RiskWarning),
	// 			ShortDescription: &armadvisor.ShortDescription{
	// 				Problem: to.Ptr("Monitoring agent should be installed on your machines"),
	// 				Solution: to.Ptr("Monitoring agent should be installed on your machines"),
	// 			},
	// 			Remediation: map[string]any{
	// 				"additionalProperties": map[string]any{
	// 					"httpMethod": "POST",
	// 					"uri": "uri",
	// 					"details": "link to document",
	// 				},
	// 			},
	// 			ResourceMetadata: &armadvisor.ResourceMetadata{
	// 				ResourceID: to.Ptr("/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset"),
	// 				Source: to.Ptr("/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset/providers/Microsoft.Security/assessments/15e912ef-221a-4efe-80d2-df5671cdd5f5"),
	// 				Action: map[string]any{
	// 					"additionalProperties": map[string]any{
	// 						"actionType": "Document",
	// 						"link": "https://link3",
	// 						"caption": "Enable Soft Delete to protect blob data",
	// 						"description": "Enable Soft Delete to protect blob data",
	// 					},
	// 					"metadata": map[string]any{
	// 						"id": "/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/resourceGroups/resourceGroup/providers/Microsoft.Compute/availabilitysets/armavset",
	// 					},
	// 				},
	// 				Singular: to.Ptr("Availability set"),
	// 				Plural: to.Ptr("Availability sets"),
	// 			},
	// 			Actions: []map[string]any{
	// 				map[string]any{
	// 					"additionalProperties": map[string]any{
	// 						"actionType": "Document",
	// 						"link": "https://link1",
	// 						"caption": "Enable Soft Delete to protect blob data",
	// 						"description": "Enable Soft Delete to protect blob data",
	// 					},
	// 					"metadata": map[string]any{
	// 						"id": "/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc",
	// 					},
	// 				},
	// 			},
	// 			Description: to.Ptr("After enabling Soft Delete, deleted data transitions to a soft deleted state instead of being permanently deleted. When data is overwritten, a soft deleted snapshot is generated to save the state of the overwritten data. You can configure the amount of time soft deleted data is recoverable before it permanently expires."),
	// 			Label: to.Ptr("Enable Soft Delete"),
	// 			LearnMoreLink: to.Ptr("https://link2"),
	// 			PotentialBenefits: to.Ptr("Save and recover your data when blobs or blob snapshots are accidentally overwritten or deleted"),
	// 			Tracked: to.Ptr(false),
	// 		},
	// 		Type: to.Ptr("Microsoft.Advisor/recommendations"),
	// 	},
	// }
}
