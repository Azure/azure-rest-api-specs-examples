package armadvisor_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/UpdateTrackedRecommendation.json
func ExampleRecommendationsClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecommendationsClient().Patch(ctx, "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/xyz", "c5532a76-8605-4328-ad27-f37ae87c086c", armadvisor.TrackedRecommendationPropertiesPayload{
		Properties: &armadvisor.TrackedRecommendationPropertiesPayloadProperties{
			TrackedProperties: &armadvisor.TrackedRecommendationProperties{
				State:         to.Ptr(armadvisor.StatePostponed),
				PostponedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-01T00:00:00Z"); return t }()),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armadvisor.RecommendationsClientPatchResponse{
	// 	ResourceRecommendationBase: armadvisor.ResourceRecommendationBase{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/xyz/providers/Microsoft.Advisor/recommendations/c5532a76-8605-4328-ad27-f37ae87c086c"),
	// 		Name: to.Ptr("c5532a76-8605-4328-ad27-f37ae87c086c"),
	// 		Type: to.Ptr("Microsoft.Advisor/recommendations"),
	// 		Properties: &armadvisor.RecommendationProperties{
	// 			Category: to.Ptr(armadvisor.CategorySecurity),
	// 			Impact: to.Ptr(armadvisor.ImpactMedium),
	// 			ImpactedField: to.Ptr("Microsoft.Compute/virtualMachines"),
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
	// 				ResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/xyz"),
	// 				Source: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/xyz/providers/Microsoft.Security/assessments/7d5f6ba5-307e-452e-b3c4-7d636f722055"),
	// 				Action: map[string]any{
	// 					"additionalProperties": map[string]any{
	// 						"actionType": "Document",
	// 						"link": "https://link3",
	// 						"caption": "Enable Soft Delete to protect blob data",
	// 						"description": "Enable Soft Delete to protect blob data",
	// 					},
	// 					"metadata": map[string]any{
	// 						"id": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.Compute/virtualMachines/xyz",
	// 					},
	// 				},
	// 				Singular: to.Ptr("Virtual machine"),
	// 				Plural: to.Ptr("Virtual machines"),
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
	// 						"id": "/subscriptions/00000000-1111-2222-3333-444444444444",
	// 					},
	// 				},
	// 			},
	// 			Description: to.Ptr("After enabling Soft Delete, deleted data transitions to a soft deleted state instead of being permanently deleted. When data is overwritten, a soft deleted snapshot is generated to save the state of the overwritten data. You can configure the amount of time soft deleted data is recoverable before it permanently expires."),
	// 			Label: to.Ptr("Enable Soft Delete"),
	// 			LearnMoreLink: to.Ptr("https://link2"),
	// 			PotentialBenefits: to.Ptr("Save and recover your data when blobs or blob snapshots are accidentally overwritten or deleted"),
	// 			Tracked: to.Ptr(true),
	// 			TrackedProperties: &armadvisor.TrackedRecommendationProperties{
	// 				State: to.Ptr(armadvisor.StatePostponed),
	// 				PostponedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-01T00:00:00Z"); return t}()),
	// 				Priority: to.Ptr(armadvisor.PriorityHigh),
	// 			},
	// 			ResourceWorkload: &armadvisor.RecommendationPropertiesResourceWorkload{
	// 				ID: to.Ptr("0560dbdb-f207-4690-b4ba-1fc0c3a0f082"),
	// 				Name: to.Ptr("Critical VM Workload"),
	// 			},
	// 			Review: &armadvisor.RecommendationPropertiesReview{
	// 				ID: to.Ptr("/subscriptions/a5481ee1-95df-47d0-85d4-dd3f0dfa19bc/providers/Microsoft.Advisor/assessments/MCWAR1"),
	// 				Name: to.Ptr("Monthly WAF Review 09/23"),
	// 			},
	// 			SourceSystem: to.Ptr("CxObserve"),
	// 			Notes: to.Ptr("This is a note"),
	// 		},
	// 	},
	// }
}
