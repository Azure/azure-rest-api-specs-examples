package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/DirectProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate_directProviderRegistrationsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProviderRegistrationsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", armproviderhub.ProviderRegistration{
		Kind: to.Ptr(armproviderhub.ProviderRegistrationKindDirect),
		Properties: &armproviderhub.ProviderRegistrationProperties{
			ProviderType:    to.Ptr(armproviderhub.ResourceProviderTypeInternal),
			ProviderVersion: to.Ptr("2.0"),
			ServiceName:     to.Ptr("root"),
			Services: []*armproviderhub.ResourceProviderService{
				{
					ServiceName: to.Ptr("tags"),
					Status:      to.Ptr(armproviderhub.ServiceStatusInactive),
				},
			},
			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
				IncidentRoutingService: to.Ptr("Contoso Resource Provider"),
				IncidentRoutingTeam:    to.Ptr("Contoso Triage"),
				IncidentContactEmail:   to.Ptr("helpme@contoso.com"),
			},
			Capabilities: []*armproviderhub.ResourceProviderCapabilities{
				{
					QuotaID: to.Ptr("CSP_2015-05-01"),
					Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
				},
				{
					QuotaID: to.Ptr("CSP_MG_2017-12-01"),
					Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
				},
			},
			NotificationSettings: &armproviderhub.ResourceProviderManifestPropertiesNotificationSettings{
				SubscriberSettings: []*armproviderhub.SubscriberSetting{
					{
						FilterRules: []*armproviderhub.FilterRule{
							{
								FilterQuery: to.Ptr("Resources | where event.eventType in ('Microsoft.Network/IpAddresses/write', 'Microsoft.KeyVault/vaults/move/action')"),
								EndpointInformation: []*armproviderhub.EndpointInformation{
									{
										Endpoint:      to.Ptr("https://userrp.azure.com/arnnotify"),
										EndpointType:  to.Ptr(armproviderhub.NotificationEndpointTypeWebhook),
										SchemaVersion: to.Ptr("3.0"),
									},
									{
										Endpoint:      to.Ptr("https://userrp.azure.com/arnnotify"),
										EndpointType:  to.Ptr(armproviderhub.NotificationEndpointTypeEventhub),
										SchemaVersion: to.Ptr("3.0"),
									},
								},
							},
						},
					},
				},
			},
			NotificationOptions: to.Ptr(armproviderhub.NotificationOptionsEmitSpendingLimit),
			ResourceHydrationAccounts: []*armproviderhub.ResourceHydrationAccount{
				{
					SubscriptionID: to.Ptr("e4eae963-2d15-43e6-a097-98bd75b33edd"),
					AccountName:    to.Ptr("classichydrationprodsn01"),
				},
				{
					SubscriptionID: to.Ptr("69e69ecb-e69c-41d4-99b8-87dd12781067"),
					AccountName:    to.Ptr("classichydrationprodch01"),
				},
			},
			ManagementGroupGlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
				{
					EndpointURI: to.Ptr("{your_management_group_notification_endpoint}"),
				},
			},
			OptionalFeatures: []*string{
				to.Ptr("Microsoft.Resources/PlatformSubscription"),
			},
			ResourceGroupLockOptionDuringMove: &armproviderhub.ResourceProviderManifestPropertiesResourceGroupLockOptionDuringMove{
				BlockActionVerb: to.Ptr(armproviderhub.BlockActionVerbAction),
			},
			ResponseOptions: &armproviderhub.ResourceProviderManifestPropertiesResponseOptions{
				ServiceClientOptionsType: to.Ptr(armproviderhub.ServiceClientOptionsTypeDisableAutomaticDecompression),
			},
			LegacyNamespace: to.Ptr("legacyNamespace"),
			LegacyRegistrations: []*string{
				to.Ptr("legacyRegistration"),
			},
			CustomManifestVersion: to.Ptr("2.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.ProviderRegistrationsClientCreateOrUpdateResponse{
	// 	ProviderRegistration: armproviderhub.ProviderRegistration{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
	// 		Name: to.Ptr("Microsoft.Contoso"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
	// 		Properties: &armproviderhub.ProviderRegistrationProperties{
	// 			Namespace: to.Ptr("Microsoft.Contoso"),
	// 			ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 				},
	// 			},
	// 			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 				ManifestOwners: []*string{
	// 					to.Ptr("manifestOwners-group"),
	// 				},
	// 				AuthorizationOwners: []*string{
	// 					to.Ptr("authorizationOwners-group"),
	// 				},
	// 				IncidentRoutingService: to.Ptr(""),
	// 				IncidentRoutingTeam: to.Ptr(""),
	// 				IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 			},
	// 			Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 				{
	// 					QuotaID: to.Ptr("CSP_2015-05-01"),
	// 					Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				},
	// 				{
	// 					QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 					Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				},
	// 			},
	// 			Metadata: map[string]any{
	// 				"onboardedVia": "ProviderHub",
	// 			},
	// 			ProviderVersion: to.Ptr("2.0"),
	// 			ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 			ServiceName: to.Ptr("root"),
	// 			Services: []*armproviderhub.ResourceProviderService{
	// 				{
	// 					ServiceName: to.Ptr("tags"),
	// 					Status: to.Ptr(armproviderhub.ServiceStatusInactive),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			NotificationSettings: &armproviderhub.ResourceProviderManifestPropertiesNotificationSettings{
	// 				SubscriberSettings: []*armproviderhub.SubscriberSetting{
	// 					{
	// 						FilterRules: []*armproviderhub.FilterRule{
	// 							{
	// 								FilterQuery: to.Ptr("Resources | where event.eventType in ('Microsoft.Network/IpAddresses/write', 'Microsoft.KeyVault/vaults/move/action')"),
	// 								EndpointInformation: []*armproviderhub.EndpointInformation{
	// 									{
	// 										Endpoint: to.Ptr("https://userrp.azure.com/arnnotify"),
	// 										EndpointType: to.Ptr(armproviderhub.NotificationEndpointTypeWebhook),
	// 										SchemaVersion: to.Ptr("3.0"),
	// 									},
	// 									{
	// 										Endpoint: to.Ptr("https://userrp.azure.com/arnnotify"),
	// 										EndpointType: to.Ptr(armproviderhub.NotificationEndpointTypeEventhub),
	// 										SchemaVersion: to.Ptr("3.0"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			NotificationOptions: to.Ptr(armproviderhub.NotificationOptionsEmitSpendingLimit),
	// 			ResourceHydrationAccounts: []*armproviderhub.ResourceHydrationAccount{
	// 				{
	// 					SubscriptionID: to.Ptr("e4eae963-2d15-43e6-a097-98bd75b33edd"),
	// 					AccountName: to.Ptr("classichydrationprodsn01"),
	// 				},
	// 				{
	// 					SubscriptionID: to.Ptr("69e69ecb-e69c-41d4-99b8-87dd12781067"),
	// 					AccountName: to.Ptr("classichydrationprodch01"),
	// 				},
	// 			},
	// 			ManagementGroupGlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 				{
	// 					EndpointURI: to.Ptr("{your_management_group_notification_endpoint}"),
	// 				},
	// 			},
	// 			OptionalFeatures: []*string{
	// 				to.Ptr("Microsoft.Resources/PlatformSubscription"),
	// 			},
	// 			ResourceGroupLockOptionDuringMove: &armproviderhub.ResourceProviderManifestPropertiesResourceGroupLockOptionDuringMove{
	// 				BlockActionVerb: to.Ptr(armproviderhub.BlockActionVerbAction),
	// 			},
	// 			ResponseOptions: &armproviderhub.ResourceProviderManifestPropertiesResponseOptions{
	// 				ServiceClientOptionsType: to.Ptr(armproviderhub.ServiceClientOptionsTypeDisableAutomaticDecompression),
	// 			},
	// 			LegacyNamespace: to.Ptr("legacyNamespace"),
	// 			LegacyRegistrations: []*string{
	// 				to.Ptr("legacyRegistration"),
	// 			},
	// 			CustomManifestVersion: to.Ptr("2.0"),
	// 		},
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 	},
	// }
}
