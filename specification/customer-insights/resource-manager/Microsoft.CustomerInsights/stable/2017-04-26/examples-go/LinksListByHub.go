package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksListByHub.json
func ExampleLinksClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinksClient().NewListByHubPager("TestHubRG", "sdkTestHub", nil)
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
		// page.LinkListResult = armcustomerinsights.LinkListResult{
		// 	Value: []*armcustomerinsights.LinkResourceFormat{
		// 		{
		// 			Name: to.Ptr("azSdkTestHub/linkTest4806"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/links"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/links/linkTest4806"),
		// 			Properties: &armcustomerinsights.LinkDefinition{
		// 				Description: map[string]*string{
		// 					"en-us": to.Ptr("Link Description"),
		// 				},
		// 				DisplayName: map[string]*string{
		// 					"en-us": to.Ptr("Link DisplayName"),
		// 				},
		// 				LinkName: to.Ptr("linkTest4806"),
		// 				Mappings: []*armcustomerinsights.TypePropertiesMapping{
		// 					{
		// 						LinkType: to.Ptr(armcustomerinsights.LinkTypesUpdateAlways),
		// 						SourcePropertyName: to.Ptr("testInteraction1949"),
		// 						TargetPropertyName: to.Ptr("testProfile1446"),
		// 				}},
		// 				ParticipantPropertyReferences: []*armcustomerinsights.ParticipantPropertyReference{
		// 					{
		// 						SourcePropertyName: to.Ptr("testInteraction1949"),
		// 						TargetPropertyName: to.Ptr("ProfileId"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 				ReferenceOnly: to.Ptr(false),
		// 				SourceEntityType: to.Ptr(armcustomerinsights.EntityTypeInteraction),
		// 				SourceEntityTypeName: to.Ptr("testInteraction1949"),
		// 				TargetEntityType: to.Ptr(armcustomerinsights.EntityTypeProfile),
		// 				TargetEntityTypeName: to.Ptr("testProfile1446"),
		// 				TenantID: to.Ptr("azsdktesthub"),
		// 			},
		// 	}},
		// }
	}
}
