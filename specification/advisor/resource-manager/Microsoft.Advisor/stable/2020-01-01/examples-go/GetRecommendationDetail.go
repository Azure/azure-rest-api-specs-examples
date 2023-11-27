package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationDetail.json
func ExampleRecommendationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecommendationsClient().Get(ctx, "resourceUri", "recommendationId", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceRecommendationBase = armadvisor.ResourceRecommendationBase{
	// 	Name: to.Ptr("recommendationId"),
	// 	Type: to.Ptr("Microsoft.Advisor/recommendations"),
	// 	ID: to.Ptr("/resourceUri/providers/Microsoft.Advisor/recommendations/recommendationId"),
	// 	Properties: &armadvisor.RecommendationProperties{
	// 		Description: to.Ptr("After enabling Soft Delete, deleted data transitions to a soft deleted state instead of being permanently deleted. When data is overwritten, a soft deleted snapshot is generated to save the state of the overwritten data. You can configure the amount of time soft deleted data is recoverable before it permanently expires."),
	// 		Actions: []map[string]any{
	// 			map[string]any{
	// 				"description": "Enable Soft Delete to protect blob data",
	// 				"actionType": "Document",
	// 				"caption": "Enable Soft Delete to protect blob data",
	// 				"link": "https://link1",
	// 				"metadata": map[string]any{
	// 					"id": "/subscriptions/subscriptionId",
	// 				},
	// 		}},
	// 		Category: to.Ptr(armadvisor.CategorySecurity),
	// 		Impact: to.Ptr(armadvisor.ImpactMedium),
	// 		ImpactedField: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		ImpactedValue: to.Ptr("armavset"),
	// 		Label: to.Ptr("Enable Soft Delete"),
	// 		LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T22:24:43.321Z"); return t}()),
	// 		LearnMoreLink: to.Ptr("https://link2"),
	// 		PotentialBenefits: to.Ptr("Save and recover your data when blobs or blob snapshots are accidentally overwritten or deleted"),
	// 		Remediation: map[string]any{
	// 			"httpMethod": "POST",
	// 			"uri": "uri",
	// 			"details": "link to document",
	// 		},
	// 		ResourceMetadata: &armadvisor.ResourceMetadata{
	// 			Action: map[string]any{
	// 				"description": "Enable Soft Delete to protect blob data",
	// 				"actionType": "Document",
	// 				"caption": "Enable Soft Delete to protect blob data",
	// 				"link": "https://link3",
	// 				"metadata": map[string]any{
	// 					"id": "/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.Compute/virtualMachines/xyz",
	// 				},
	// 			},
	// 			Plural: to.Ptr("Virtual machines"),
	// 			ResourceID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.Compute/virtualMachines/xyz"),
	// 			Singular: to.Ptr("Virtual machine"),
	// 			Source: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.Compute/virtualMachines/xyz/providers/Microsoft.Security/assessments/assessmentGuid"),
	// 		},
	// 		Risk: to.Ptr(armadvisor.RiskWarning),
	// 		ShortDescription: &armadvisor.ShortDescription{
	// 			Problem: to.Ptr("Monitoring agent is not installed on your machines"),
	// 			Solution: to.Ptr("Monitoring agent should be installed on your machines"),
	// 		},
	// 	},
	// }
}
