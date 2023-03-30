package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsCreateOrUpdate.json
func ExampleRelationshipsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRelationshipsClient().BeginCreateOrUpdate(ctx, "TestHubRG", "sdkTestHub", "SomeRelationship", armcustomerinsights.RelationshipResourceFormat{
		Properties: &armcustomerinsights.RelationshipDefinition{
			Description: map[string]*string{
				"en-us": to.Ptr("Relationship Description"),
			},
			Cardinality: to.Ptr(armcustomerinsights.CardinalityTypesOneToOne),
			DisplayName: map[string]*string{
				"en-us": to.Ptr("Relationship DisplayName"),
			},
			Fields:             []*armcustomerinsights.PropertyDefinition{},
			ProfileType:        to.Ptr("testProfile2326994"),
			RelatedProfileType: to.Ptr("testProfile2326994"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RelationshipResourceFormat = armcustomerinsights.RelationshipResourceFormat{
	// 	Name: to.Ptr("sdkTestHub/testProfile2326994"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/relationships"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/relationships/SomeRelationship"),
	// 	Properties: &armcustomerinsights.RelationshipDefinition{
	// 		Description: map[string]*string{
	// 			"en-us": to.Ptr("Relationship Description"),
	// 		},
	// 		Cardinality: to.Ptr(armcustomerinsights.CardinalityTypesOneToOne),
	// 		DisplayName: map[string]*string{
	// 			"en-us": to.Ptr("Relationship DisplayName"),
	// 		},
	// 		Fields: []*armcustomerinsights.PropertyDefinition{
	// 		},
	// 		LookupMappings: []*armcustomerinsights.RelationshipTypeMapping{
	// 		},
	// 		ProfileType: to.Ptr("testProfile2326994"),
	// 		ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 		RelatedProfileType: to.Ptr("testProfile2326994"),
	// 		RelationshipName: to.Ptr("SomeRelationship"),
	// 		TenantID: to.Ptr("sdktesthub"),
	// 	},
	// }
}
