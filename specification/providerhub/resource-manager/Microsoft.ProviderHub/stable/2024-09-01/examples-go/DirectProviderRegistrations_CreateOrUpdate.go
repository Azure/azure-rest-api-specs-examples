package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/DirectProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate_directProviderRegistrationsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProviderRegistrationsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", armproviderhub.ProviderRegistration{
		Kind: to.Ptr(armproviderhub.ProviderRegistrationKindDirect),
		Properties: &armproviderhub.ProviderRegistrationProperties{
			Capabilities: []*armproviderhub.ResourceProviderCapabilities{
				{
					Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
					QuotaID: to.Ptr("CSP_2015-05-01"),
				},
				{
					Effect:  to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
					QuotaID: to.Ptr("CSP_MG_2017-12-01"),
				}},
			CustomManifestVersion: to.Ptr("2.0"),
			DstsConfiguration: &armproviderhub.ResourceProviderManifestPropertiesDstsConfiguration{
				ServiceDNSName: to.Ptr("prds.sparta.azure.com"),
				ServiceName:    to.Ptr("prds-shim"),
			},
			LegacyNamespace: to.Ptr("legacyNamespace"),
			LegacyRegistrations: []*string{
				to.Ptr("legacyRegistration")},
			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
				IncidentContactEmail:   to.Ptr("helpme@contoso.com"),
				IncidentRoutingService: to.Ptr("Contoso Resource Provider"),
				IncidentRoutingTeam:    to.Ptr("Contoso Triage"),
				ServiceTreeInfos: []*armproviderhub.ServiceTreeInfo{
					{
						ComponentID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
						Readiness:   to.Ptr(armproviderhub.ReadinessInDevelopment),
						ServiceID:   to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
					}},
			},
			ManagementGroupGlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
				{
					EndpointURI: to.Ptr("{your_management_group_notification_endpoint}"),
				}},
			NotificationOptions: to.Ptr(armproviderhub.NotificationOptionsEmitSpendingLimit),
			NotificationSettings: &armproviderhub.ResourceProviderManifestPropertiesNotificationSettings{
				SubscriberSettings: []*armproviderhub.SubscriberSetting{
					{
						FilterRules: []*armproviderhub.FilterRule{
							{
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
									}},
								FilterQuery: to.Ptr("Resources | where event.eventType in ('Microsoft.Network/IpAddresses/write', 'Microsoft.KeyVault/vaults/move/action')"),
							}},
					}},
			},
			OptionalFeatures: []*string{
				to.Ptr("Microsoft.Resources/PlatformSubscription")},
			ProviderType:    to.Ptr(armproviderhub.ResourceProviderTypeInternal),
			ProviderVersion: to.Ptr("2.0"),
			ResourceGroupLockOptionDuringMove: &armproviderhub.ResourceProviderManifestPropertiesResourceGroupLockOptionDuringMove{
				BlockActionVerb: to.Ptr(armproviderhub.BlockActionVerbAction),
			},
			ResourceHydrationAccounts: []*armproviderhub.ResourceHydrationAccount{
				{
					AccountName:    to.Ptr("classichydrationprodsn01"),
					SubscriptionID: to.Ptr("e4eae963-2d15-43e6-a097-98bd75b33edd"),
				},
				{
					AccountName:    to.Ptr("classichydrationprodch01"),
					SubscriptionID: to.Ptr("69e69ecb-e69c-41d4-99b8-87dd12781067"),
				}},
			ResponseOptions: &armproviderhub.ResourceProviderManifestPropertiesResponseOptions{
				ServiceClientOptionsType: to.Ptr(armproviderhub.ServiceClientOptionsTypeDisableAutomaticDecompression),
			},
			ServiceName: to.Ptr("root"),
			Services: []*armproviderhub.ResourceProviderService{
				{
					ServiceName: to.Ptr("tags"),
					Status:      to.Ptr(armproviderhub.ServiceStatusInactive),
				}},
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
	// res.ProviderRegistration = armproviderhub.ProviderRegistration{
	// 	Name: to.Ptr("Microsoft.Contoso"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.ProviderRegistrationProperties{
	// 		Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 			{
	// 				Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				QuotaID: to.Ptr("CSP_2015-05-01"),
	// 			},
	// 			{
	// 				Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 				QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 		}},
	// 		CustomManifestVersion: to.Ptr("2.0"),
	// 		DstsConfiguration: &armproviderhub.ResourceProviderManifestPropertiesDstsConfiguration{
	// 			ServiceDNSName: to.Ptr("prds.sparta.azure.com"),
	// 			ServiceName: to.Ptr("prds-shim"),
	// 		},
	// 		LegacyNamespace: to.Ptr("legacyNamespace"),
	// 		LegacyRegistrations: []*string{
	// 			to.Ptr("legacyRegistration")},
	// 			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 				AuthorizationOwners: []*string{
	// 					to.Ptr("authorizationOwners-group")},
	// 					IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 					IncidentRoutingService: to.Ptr(""),
	// 					IncidentRoutingTeam: to.Ptr(""),
	// 					ManifestOwners: []*string{
	// 						to.Ptr("manifestOwners-group")},
	// 						ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 						ServiceTreeInfos: []*armproviderhub.ServiceTreeInfo{
	// 							{
	// 								ComponentID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
	// 								Readiness: to.Ptr(armproviderhub.ReadinessInDevelopment),
	// 								ServiceID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
	// 						}},
	// 					},
	// 					ManagementGroupGlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 						{
	// 							EndpointURI: to.Ptr("{your_management_group_notification_endpoint}"),
	// 					}},
	// 					Metadata: map[string]any{
	// 						"onboardedVia": "ProviderHub",
	// 					},
	// 					Namespace: to.Ptr("Microsoft.Contoso"),
	// 					NotificationOptions: to.Ptr(armproviderhub.NotificationOptionsEmitSpendingLimit),
	// 					NotificationSettings: &armproviderhub.ResourceProviderManifestPropertiesNotificationSettings{
	// 						SubscriberSettings: []*armproviderhub.SubscriberSetting{
	// 							{
	// 								FilterRules: []*armproviderhub.FilterRule{
	// 									{
	// 										EndpointInformation: []*armproviderhub.EndpointInformation{
	// 											{
	// 												Endpoint: to.Ptr("https://userrp.azure.com/arnnotify"),
	// 												EndpointType: to.Ptr(armproviderhub.NotificationEndpointTypeWebhook),
	// 												SchemaVersion: to.Ptr("3.0"),
	// 											},
	// 											{
	// 												Endpoint: to.Ptr("https://userrp.azure.com/arnnotify"),
	// 												EndpointType: to.Ptr(armproviderhub.NotificationEndpointTypeEventhub),
	// 												SchemaVersion: to.Ptr("3.0"),
	// 										}},
	// 										FilterQuery: to.Ptr("Resources | where event.eventType in ('Microsoft.Network/IpAddresses/write', 'Microsoft.KeyVault/vaults/move/action')"),
	// 								}},
	// 						}},
	// 					},
	// 					OptionalFeatures: []*string{
	// 						to.Ptr("Microsoft.Resources/PlatformSubscription")},
	// 						ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 							{
	// 								ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 								RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 						}},
	// 						ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 						ProviderVersion: to.Ptr("2.0"),
	// 						ResourceGroupLockOptionDuringMove: &armproviderhub.ResourceProviderManifestPropertiesResourceGroupLockOptionDuringMove{
	// 							BlockActionVerb: to.Ptr(armproviderhub.BlockActionVerbAction),
	// 						},
	// 						ResourceHydrationAccounts: []*armproviderhub.ResourceHydrationAccount{
	// 							{
	// 								AccountName: to.Ptr("classichydrationprodsn01"),
	// 								SubscriptionID: to.Ptr("e4eae963-2d15-43e6-a097-98bd75b33edd"),
	// 							},
	// 							{
	// 								AccountName: to.Ptr("classichydrationprodch01"),
	// 								SubscriptionID: to.Ptr("69e69ecb-e69c-41d4-99b8-87dd12781067"),
	// 						}},
	// 						ResponseOptions: &armproviderhub.ResourceProviderManifestPropertiesResponseOptions{
	// 							ServiceClientOptionsType: to.Ptr(armproviderhub.ServiceClientOptionsTypeDisableAutomaticDecompression),
	// 						},
	// 						ServiceName: to.Ptr("root"),
	// 						Services: []*armproviderhub.ResourceProviderService{
	// 							{
	// 								ServiceName: to.Ptr("tags"),
	// 								Status: to.Ptr(armproviderhub.ServiceStatusInactive),
	// 						}},
	// 						ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 					},
	// 				}
}
