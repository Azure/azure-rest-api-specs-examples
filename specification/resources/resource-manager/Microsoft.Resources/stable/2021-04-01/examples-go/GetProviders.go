package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProviders.json
func ExampleProvidersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProvidersClient().NewListPager(&armresources.ProvidersClientListOptions{Expand: nil})
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
		// page.ProviderListResult = armresources.ProviderListResult{
		// 	Value: []*armresources.Provider{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.TestRP1"),
		// 			Namespace: to.Ptr("Microsoft.TestRP1"),
		// 			RegistrationPolicy: to.Ptr("RegistrationRequired"),
		// 			RegistrationState: to.Ptr("Registering"),
		// 			ResourceTypes: []*armresources.ProviderResourceType{
		// 				{
		// 					APIVersions: []*string{
		// 						to.Ptr("2018-01-01"),
		// 						to.Ptr("2015-05-01")},
		// 						Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 						LocationMappings: []*armresources.ProviderExtendedLocation{
		// 							{
		// 								Type: to.Ptr("EdgeZone"),
		// 								ExtendedLocations: []*string{
		// 									to.Ptr("LosAngeles"),
		// 									to.Ptr("LosAngeles2")},
		// 									Location: to.Ptr("West US"),
		// 							}},
		// 							Locations: []*string{
		// 								to.Ptr("West US")},
		// 								ResourceType: to.Ptr("TestResourceType"),
		// 							},
		// 							{
		// 								APIVersions: []*string{
		// 									to.Ptr("2018-01-01"),
		// 									to.Ptr("2015-05-01")},
		// 									Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 									Locations: []*string{
		// 										to.Ptr("West US")},
		// 										ResourceType: to.Ptr("TestResourceTypeSibling"),
		// 								}},
		// 							},
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources"),
		// 								Namespace: to.Ptr("Microsoft.Resources"),
		// 								RegistrationPolicy: to.Ptr("RegistrationFree"),
		// 								RegistrationState: to.Ptr("Registered"),
		// 								ResourceTypes: []*armresources.ProviderResourceType{
		// 									{
		// 										Aliases: []*armresources.Alias{
		// 										},
		// 										APIVersions: []*string{
		// 											to.Ptr("2016-09-01"),
		// 											to.Ptr("2014-04-01-preview")},
		// 											Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 											Locations: []*string{
		// 												to.Ptr("eastus"),
		// 												to.Ptr("eastus2"),
		// 												to.Ptr("westus")},
		// 												ResourceType: to.Ptr("subscriptions"),
		// 											},
		// 											{
		// 												Aliases: []*armresources.Alias{
		// 												},
		// 												APIVersions: []*string{
		// 													to.Ptr("2016-09-01"),
		// 													to.Ptr("2014-04-01-preview")},
		// 													Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 													Locations: []*string{
		// 														to.Ptr("centralus"),
		// 														to.Ptr("eastasia"),
		// 														to.Ptr("southeastasia")},
		// 														ResourceType: to.Ptr("resourceGroups"),
		// 													},
		// 													{
		// 														Aliases: []*armresources.Alias{
		// 														},
		// 														APIVersions: []*string{
		// 															to.Ptr("2016-09-01"),
		// 															to.Ptr("2014-04-01-preview")},
		// 															Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 															Locations: []*string{
		// 																to.Ptr("eastus"),
		// 																to.Ptr("eastus2"),
		// 																to.Ptr("westus")},
		// 																ResourceType: to.Ptr("subscriptions/resourceGroups"),
		// 															},
		// 															{
		// 																Aliases: []*armresources.Alias{
		// 																},
		// 																APIVersions: []*string{
		// 																	to.Ptr("2014-04-01-preview")},
		// 																	Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 																	Locations: []*string{
		// 																		to.Ptr("centralus"),
		// 																		to.Ptr("eastasia")},
		// 																		ResourceType: to.Ptr("bulkDelete"),
		// 																	},
		// 																	{
		// 																		Aliases: []*armresources.Alias{
		// 																		},
		// 																		APIVersions: []*string{
		// 																			to.Ptr("2017-08-01"),
		// 																			to.Ptr("2017-06-01")},
		// 																			Capabilities: to.Ptr("SupportsTags"),
		// 																			Locations: []*string{
		// 																			},
		// 																			ResourceType: to.Ptr("deployments"),
		// 																		},
		// 																		{
		// 																			Aliases: []*armresources.Alias{
		// 																			},
		// 																			APIVersions: []*string{
		// 																			},
		// 																			Capabilities: to.Ptr("SupportsExtension"),
		// 																			Locations: []*string{
		// 																				to.Ptr("DevFabric")},
		// 																				ResourceType: to.Ptr("tags"),
		// 																		}},
		// 																	},
		// 																	{
		// 																		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.TestRP2"),
		// 																		Namespace: to.Ptr("Microsoft.TestRP2"),
		// 																		RegistrationPolicy: to.Ptr("RegistrationRequired"),
		// 																		RegistrationState: to.Ptr("NotRegistered"),
		// 																		ResourceTypes: []*armresources.ProviderResourceType{
		// 																			{
		// 																				APIVersions: []*string{
		// 																					to.Ptr("2018-01-01"),
		// 																					to.Ptr("2015-05-01")},
		// 																					Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 																					LocationMappings: []*armresources.ProviderExtendedLocation{
		// 																						{
		// 																							Type: to.Ptr("EdgeZone"),
		// 																							ExtendedLocations: []*string{
		// 																								to.Ptr("LosAngeles"),
		// 																								to.Ptr("LosAngeles2")},
		// 																								Location: to.Ptr("West US"),
		// 																						}},
		// 																						Locations: []*string{
		// 																							to.Ptr("West US")},
		// 																							ResourceType: to.Ptr("TestResourceType"),
		// 																						},
		// 																						{
		// 																							APIVersions: []*string{
		// 																								to.Ptr("2018-01-01"),
		// 																								to.Ptr("2015-05-01")},
		// 																								Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
		// 																								Locations: []*string{
		// 																									to.Ptr("West US")},
		// 																									ResourceType: to.Ptr("TestResourceTypeSibling"),
		// 																							}},
		// 																					}},
		// 																				}
	}
}
