package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor/v2"
)

// Generated from example definition: 2025-05-01-preview/ListRecommendationsServiceGroupResourceUri.json
func ExampleRecommendationsClient_NewListByTenantPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecommendationsClient().NewListByTenantPager("providers/microsoft.management/serviceGroups/serviceGroupXYZ", &armadvisor.RecommendationsClientListByTenantOptions{
		Top: to.Ptr[int32](10)})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armadvisor.RecommendationsClientListByTenantResponse{
		// 	ResourceRecommendationBaseListResult: armadvisor.ResourceRecommendationBaseListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/providers/microsoft.management/serviceGroup/serviceGroupXYZ/providers/Microsoft.Advisor/recommendations?api-version=2025-05-01-preview&$top=10&$skiptoken=skiptoken"),
		// 		Value: []*armadvisor.ResourceRecommendationBase{
		// 			{
		// 				ID: to.Ptr("/providers/microsoft.management/serviceGroup/serviceGroupXYZ/providers/Microsoft.Advisor/recommendations/37c93209-4bfb-4f3b-8874-ccc718f7a467"),
		// 				Name: to.Ptr("37c93209-4bfb-4f3b-8874-ccc718f7a467"),
		// 				Type: to.Ptr("Microsoft.Advisor/recommendations"),
		// 				Properties: &armadvisor.RecommendationProperties{
		// 					Category: to.Ptr(armadvisor.CategoryHighAvailability),
		// 					Impact: to.Ptr(armadvisor.ImpactMedium),
		// 					ImpactedField: to.Ptr("Microsoft.Management/serviceGroup"),
		// 					ImpactedValue: to.Ptr("serviceGroupXYZ"),
		// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-24T22:24:43.3216408Z"); return t}()),
		// 					Risk: to.Ptr(armadvisor.RiskWarning),
		// 					ShortDescription: &armadvisor.ShortDescription{
		// 						Problem: to.Ptr("To ensure high availability add one or more virtual machines to this availability set"),
		// 						Solution: to.Ptr("To ensure high availability add one or more virtual machines to this availability set"),
		// 					},
		// 					ResourceMetadata: &armadvisor.ResourceMetadata{
		// 						ResourceID: to.Ptr("/providers/microsoft.management/serviceGroup/serviceGroupXYZ"),
		// 						Action: map[string]any{
		// 							"additionalProperties": map[string]any{
		// 								"actionType": "Document",
		// 								"link": "https://link3",
		// 								"caption": "Enable Soft Delete to protect blob data",
		// 								"description": "Enable Soft Delete to protect blob data",
		// 							},
		// 							"metadata": map[string]any{
		// 								"id": "/providers/microsoft.management/serviceGroup/serviceGroupXYZ",
		// 							},
		// 						},
		// 						Singular: to.Ptr("Availability set"),
		// 						Plural: to.Ptr("Availability sets"),
		// 					},
		// 					Actions: []map[string]any{
		// 						map[string]any{
		// 							"additionalProperties": map[string]any{
		// 								"actionType": "Document",
		// 								"link": "https://link1",
		// 								"caption": "Enable Soft Delete to protect blob data",
		// 								"description": "Enable Soft Delete to protect blob data",
		// 							},
		// 							"metadata": map[string]any{
		// 								"id": "/providers/microsoft.management/serviceGroup/serviceGroupXYZ",
		// 							},
		// 						},
		// 					},
		// 					Description: to.Ptr("After enabling Soft Delete, deleted data transitions to a soft deleted state instead of being permanently deleted. When data is overwritten, a soft deleted snapshot is generated to save the state of the overwritten data. You can configure the amount of time soft deleted data is recoverable before it permanently expires."),
		// 					Label: to.Ptr("Enable Soft Delete"),
		// 					LearnMoreLink: to.Ptr("https://link2"),
		// 					PotentialBenefits: to.Ptr("Save and recover your data when blobs or blob snapshots are accidentally overwritten or deleted"),
		// 					Tracked: to.Ptr(false),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/providers/microsoft.management/serviceGroup/serviceGroupXYZ/providers/Microsoft.Advisor/recommendations/de0eb863-583e-ac10-eb7f-cb0f6b570e2b"),
		// 				Name: to.Ptr("de0eb863-583e-ac10-eb7f-cb0f6b570e2b"),
		// 				Type: to.Ptr("Microsoft.Advisor/recommendations"),
		// 				Properties: &armadvisor.RecommendationProperties{
		// 					Category: to.Ptr(armadvisor.CategoryHighAvailability),
		// 					Impact: to.Ptr(armadvisor.ImpactMedium),
		// 					ImpactedField: to.Ptr("Microsoft.Management/serviceGroup"),
		// 					ImpactedValue: to.Ptr("serviceGroupXYZ"),
		// 					LastUpdated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-04-24T22:24:43.3216408Z"); return t}()),
		// 					Risk: to.Ptr(armadvisor.RiskWarning),
		// 					ShortDescription: &armadvisor.ShortDescription{
		// 						Problem: to.Ptr("To ensure high availability add one or more virtual machines to this availability set"),
		// 						Solution: to.Ptr("To ensure high availability add one or more virtual machines to this availability set"),
		// 					},
		// 					ResourceMetadata: &armadvisor.ResourceMetadata{
		// 						ResourceID: to.Ptr("/providers/microsoft.management/serviceGroup/serviceGroupXYZ"),
		// 						Action: map[string]any{
		// 							"additionalProperties": map[string]any{
		// 								"actionType": "Document",
		// 								"link": "https://link3",
		// 								"caption": "Enable Soft Delete to protect blob data",
		// 								"description": "Enable Soft Delete to protect blob data",
		// 							},
		// 							"metadata": map[string]any{
		// 								"id": "/providers/microsoft.management/serviceGroup/serviceGroupXYZ",
		// 							},
		// 						},
		// 						Singular: to.Ptr("Availability set"),
		// 						Plural: to.Ptr("Availability sets"),
		// 					},
		// 					Actions: []map[string]any{
		// 						map[string]any{
		// 							"additionalProperties": map[string]any{
		// 								"actionType": "Document",
		// 								"link": "https://link1",
		// 								"caption": "Enable Soft Delete to protect blob data",
		// 								"description": "Enable Soft Delete to protect blob data",
		// 							},
		// 							"metadata": map[string]any{
		// 								"id": "/providers/microsoft.management/serviceGroup/serviceGroupXYZ",
		// 							},
		// 						},
		// 					},
		// 					Description: to.Ptr("After enabling Soft Delete, deleted data transitions to a soft deleted state instead of being permanently deleted. When data is overwritten, a soft deleted snapshot is generated to save the state of the overwritten data. You can configure the amount of time soft deleted data is recoverable before it permanently expires."),
		// 					Label: to.Ptr("Enable Soft Delete"),
		// 					LearnMoreLink: to.Ptr("https://link2"),
		// 					PotentialBenefits: to.Ptr("Save and recover your data when blobs or blob snapshots are accidentally overwritten or deleted"),
		// 					Tracked: to.Ptr(false),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
