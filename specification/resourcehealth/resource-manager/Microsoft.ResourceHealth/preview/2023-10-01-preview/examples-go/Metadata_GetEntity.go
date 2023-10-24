package armresourcehealth_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Metadata_GetEntity.json
func ExampleMetadataClient_GetEntity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcehealth.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMetadataClient().GetEntity(ctx, "status", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MetadataEntity = armresourcehealth.MetadataEntity{
	// 	Name: to.Ptr("status"),
	// 	Type: to.Ptr("Microsoft.ResourceHealth/metadata"),
	// 	ID: to.Ptr("/providers/Microsoft.ResourceHealth/metadata/status"),
	// 	Properties: &armresourcehealth.MetadataEntityProperties{
	// 		ApplicableScenarios: []*armresourcehealth.Scenario{
	// 			to.Ptr(armresourcehealth.ScenarioAlerts)},
	// 			DisplayName: to.Ptr("Status"),
	// 			SupportedValues: []*armresourcehealth.MetadataSupportedValueDetail{
	// 				{
	// 					DisplayName: to.Ptr("Active"),
	// 					ID: to.Ptr("Active"),
	// 				},
	// 				{
	// 					DisplayName: to.Ptr("Resolved"),
	// 					ID: to.Ptr("Resolved"),
	// 				},
	// 				{
	// 					DisplayName: to.Ptr("In Progress"),
	// 					ID: to.Ptr("In Progress"),
	// 				},
	// 				{
	// 					DisplayName: to.Ptr("Updated"),
	// 					ID: to.Ptr("Updated"),
	// 			}},
	// 		},
	// 	}
}
