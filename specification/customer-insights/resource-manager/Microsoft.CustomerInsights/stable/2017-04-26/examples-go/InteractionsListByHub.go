package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsListByHub.json
func ExampleInteractionsClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInteractionsClient().NewListByHubPager("TestHubRG", "sdkTestHub", &armcustomerinsights.InteractionsClientListByHubOptions{LocaleCode: nil})
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
		// page.InteractionListResult = armcustomerinsights.InteractionListResult{
		// 	Value: []*armcustomerinsights.InteractionResourceFormat{
		// 		{
		// 			Name: to.Ptr("azSdkTestHub/TestInteractionType6358"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/interactions"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/interactions/TestInteractionType6358"),
		// 			Properties: &armcustomerinsights.InteractionTypeDefinition{
		// 				LargeImage: to.Ptr("\\\\Images\\\\LargeImage"),
		// 				MediumImage: to.Ptr("\\\\Images\\\\MediumImage"),
		// 				SmallImage: to.Ptr("\\\\Images\\\\smallImage"),
		// 				APIEntitySetName: to.Ptr("TestInteractionType6358"),
		// 				EntityType: to.Ptr(armcustomerinsights.EntityTypesInteraction),
		// 				Fields: []*armcustomerinsights.PropertyDefinition{
		// 					{
		// 						DataSourcePrecedenceRules: []*armcustomerinsights.DataSourcePrecedence{
		// 						},
		// 						FieldName: to.Ptr("TestInteractionType6358"),
		// 						FieldType: to.Ptr("Edm.String"),
		// 						IsArray: to.Ptr(false),
		// 						IsEnum: to.Ptr(false),
		// 						IsFlagEnum: to.Ptr(false),
		// 						IsImage: to.Ptr(false),
		// 						IsLocalizedString: to.Ptr(false),
		// 						IsName: to.Ptr(false),
		// 						IsRequired: to.Ptr(true),
		// 						PropertyID: to.Ptr("id1"),
		// 					},
		// 					{
		// 						DataSourcePrecedenceRules: []*armcustomerinsights.DataSourcePrecedence{
		// 							{
		// 								DataSource: &armcustomerinsights.DataSource{
		// 									Name: to.Ptr("crm2"),
		// 									DataSourceReferenceID: to.Ptr("70"),
		// 									DataSourceType: to.Ptr(armcustomerinsights.DataSourceTypeConnector),
		// 									ID: to.Ptr[int32](1),
		// 									Status: to.Ptr(armcustomerinsights.StatusActive),
		// 								},
		// 								Precedence: to.Ptr[int32](50),
		// 						}},
		// 						FieldName: to.Ptr("profile1"),
		// 						FieldType: to.Ptr("Edm.String"),
		// 						IsArray: to.Ptr(false),
		// 						IsEnum: to.Ptr(false),
		// 						IsFlagEnum: to.Ptr(false),
		// 						IsImage: to.Ptr(false),
		// 						IsLocalizedString: to.Ptr(false),
		// 						IsName: to.Ptr(false),
		// 						IsRequired: to.Ptr(true),
		// 						PropertyID: to.Ptr("id1"),
		// 				}},
		// 				InstancesCount: to.Ptr[int32](0),
		// 				ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 				TenantID: to.Ptr("azsdktesthub"),
		// 				TypeName: to.Ptr("TestInteractionType6358"),
		// 				DataSourcePrecedenceRules: []*armcustomerinsights.DataSourcePrecedence{
		// 					{
		// 						DataSource: &armcustomerinsights.DataSource{
		// 							Name: to.Ptr("crm2"),
		// 							DataSourceReferenceID: to.Ptr("70"),
		// 							DataSourceType: to.Ptr(armcustomerinsights.DataSourceTypeConnector),
		// 							ID: to.Ptr[int32](1),
		// 							Status: to.Ptr(armcustomerinsights.StatusActive),
		// 						},
		// 						Precedence: to.Ptr[int32](50),
		// 				}},
		// 				DefaultDataSource: &armcustomerinsights.DataSource{
		// 					Name: to.Ptr("crm2"),
		// 					DataSourceReferenceID: to.Ptr("70"),
		// 					DataSourceType: to.Ptr(armcustomerinsights.DataSourceTypeConnector),
		// 					ID: to.Ptr[int32](1),
		// 					Status: to.Ptr(armcustomerinsights.StatusActive),
		// 				},
		// 				IDPropertyNames: []*string{
		// 					to.Ptr("TestInteractionType6358")},
		// 					IsActivity: to.Ptr(true),
		// 					PrimaryParticipantProfilePropertyName: to.Ptr("profile1"),
		// 				},
		// 		}},
		// 	}
	}
}
