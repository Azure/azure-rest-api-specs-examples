package armadvisor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/advisor/armadvisor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationMetadataEntity.json
func ExampleRecommendationMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armadvisor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecommendationMetadataClient().Get(ctx, "types", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataEntity = armadvisor.MetadataEntity{
	// 	Name: to.Ptr("recommendationType"),
	// 	Type: to.Ptr("Microsoft.Advisor/metadata"),
	// 	ID: to.Ptr("providers/Microsoft.Advisor/metadata/recommendationType"),
	// 	Properties: &armadvisor.MetadataEntityProperties{
	// 		ApplicableScenarios: []*armadvisor.Scenario{
	// 			to.Ptr(armadvisor.ScenarioAlerts)},
	// 			DependsOn: []*string{
	// 				to.Ptr("category"),
	// 				to.Ptr("impact")},
	// 				DisplayName: to.Ptr("Recommendation Type"),
	// 				SupportedValues: []*armadvisor.MetadataSupportedValueDetail{
	// 					{
	// 						DisplayName: to.Ptr("Upgrade your SKU or add more instances to ensure fault tolerance"),
	// 						ID: to.Ptr("6a2b1e70-bd4c-4163-86de-5243d7ac05ee"),
	// 					},
	// 					{
	// 						DisplayName: to.Ptr("Delete ExpressRoute circuits in the provider status of Not Provisioned"),
	// 						ID: to.Ptr("da6630fb-4286-4996-92a3-a43f5f26dd34"),
	// 				}},
	// 			},
	// 		}
}
