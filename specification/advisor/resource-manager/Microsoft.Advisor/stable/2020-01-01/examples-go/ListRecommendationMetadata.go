package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListRecommendationMetadata.json
func ExampleRecommendationMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecommendationMetadataClient().NewListPager(nil)
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
		// page.MetadataEntityListResult = armadvisor.MetadataEntityListResult{
		// 	Value: []*armadvisor.MetadataEntity{
		// 		{
		// 			Name: to.Ptr("recommendationType"),
		// 			Type: to.Ptr("Microsoft.Advisor/metadata"),
		// 			ID: to.Ptr("providers/Microsoft.Advisor/metadata/recommendationType"),
		// 			Properties: &armadvisor.MetadataEntityProperties{
		// 				ApplicableScenarios: []*armadvisor.Scenario{
		// 					to.Ptr(armadvisor.ScenarioAlerts)},
		// 					DependsOn: []*string{
		// 						to.Ptr("category"),
		// 						to.Ptr("impact")},
		// 						DisplayName: to.Ptr("Recommendation Type"),
		// 						SupportedValues: []*armadvisor.MetadataSupportedValueDetail{
		// 							{
		// 								DisplayName: to.Ptr("Upgrade your SKU or add more instances to ensure fault tolerance"),
		// 								ID: to.Ptr("6a2b1e70-bd4c-4163-86de-5243d7ac05ee"),
		// 							},
		// 							{
		// 								DisplayName: to.Ptr("Delete ExpressRoute circuits in the provider status of Not Provisioned"),
		// 								ID: to.Ptr("da6630fb-4286-4996-92a3-a43f5f26dd34"),
		// 						}},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("recommendationCategory"),
		// 					Type: to.Ptr("Microsoft.Advisor/metadata"),
		// 					ID: to.Ptr("providers/Microsoft.Advisor/metadata/recommendationCategory"),
		// 					Properties: &armadvisor.MetadataEntityProperties{
		// 						ApplicableScenarios: []*armadvisor.Scenario{
		// 							to.Ptr(armadvisor.ScenarioAlerts)},
		// 							DisplayName: to.Ptr("Category"),
		// 							SupportedValues: []*armadvisor.MetadataSupportedValueDetail{
		// 								{
		// 									DisplayName: to.Ptr("Cost"),
		// 									ID: to.Ptr("Cost"),
		// 								},
		// 								{
		// 									DisplayName: to.Ptr("Performance"),
		// 									ID: to.Ptr("Performance"),
		// 							}},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("recommendationImpact"),
		// 						Type: to.Ptr("Microsoft.Advisor/metadata"),
		// 						ID: to.Ptr("providers/Microsoft.Advisor/metadata/recommendationImpact"),
		// 						Properties: &armadvisor.MetadataEntityProperties{
		// 							DisplayName: to.Ptr("Impact"),
		// 							SupportedValues: []*armadvisor.MetadataSupportedValueDetail{
		// 								{
		// 									DisplayName: to.Ptr("High"),
		// 									ID: to.Ptr("High"),
		// 								},
		// 								{
		// 									DisplayName: to.Ptr("Medium"),
		// 									ID: to.Ptr("Medium"),
		// 								},
		// 								{
		// 									DisplayName: to.Ptr("Low"),
		// 									ID: to.Ptr("Low"),
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
