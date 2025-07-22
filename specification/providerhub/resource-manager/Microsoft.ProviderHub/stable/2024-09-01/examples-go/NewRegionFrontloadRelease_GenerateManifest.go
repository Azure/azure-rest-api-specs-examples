package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/NewRegionFrontloadRelease_GenerateManifest.json
func ExampleNewRegionFrontloadReleaseClient_GenerateManifest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNewRegionFrontloadReleaseClient().GenerateManifest(ctx, "Microsoft.Contoso", armproviderhub.FrontloadPayload{
		Properties: &armproviderhub.FrontloadPayloadProperties{
			CopyFromLocation: to.Ptr("eastus"),
			EnvironmentType:  to.Ptr(armproviderhub.AvailableCheckInManifestEnvironmentProd),
			ExcludeResourceTypes: []*string{
				to.Ptr("monitors")},
			FrontloadLocation: to.Ptr("Israel Central"),
			IgnoreFields: []*string{
				to.Ptr("apiversion")},
			IncludeResourceTypes: []*string{
				to.Ptr("servers")},
			OperationType: to.Ptr("Rollout"),
			OverrideEndpointLevelFields: &armproviderhub.FrontloadPayloadPropertiesOverrideEndpointLevelFields{
				APIVersion: to.Ptr("2024-04-01-preview"),
				APIVersions: []*string{
					to.Ptr("2024-04-01-preview")},
				DstsConfiguration: &armproviderhub.ResourceTypeEndpointBaseDstsConfiguration{
					ServiceDNSName: to.Ptr("messaging.azure-ppe.net"),
					ServiceName:    to.Ptr("resourceprovider"),
				},
				Enabled:      to.Ptr(true),
				EndpointType: to.Ptr(armproviderhub.EndpointTypeProduction),
				EndpointURI:  to.Ptr("https://resource-endpoint.com/"),
				FeaturesRule: &armproviderhub.ResourceTypeEndpointBaseFeaturesRule{},
				Locations: []*string{
					to.Ptr("East US")},
				RequiredFeatures: []*string{
					to.Ptr("<feature flag>")},
				SKULink: to.Ptr("http://endpointuri/westus/skus"),
				Timeout: to.Ptr("PT20S"),
				Zones: []*string{
					to.Ptr("zone1")},
			},
			OverrideManifestLevelFields: &armproviderhub.FrontloadPayloadPropertiesOverrideManifestLevelFields{
				ResourceHydrationAccounts: []*armproviderhub.ResourceHydrationAccount{
					{
						AccountName:    to.Ptr("classichydrationprodsn01"),
						SubscriptionID: to.Ptr("e4eae963-2d15-43e6-a097-98bd75b33edd"),
					}},
			},
			ProviderNamespace:  to.Ptr("Microsoft.Contoso"),
			ServiceFeatureFlag: to.Ptr(armproviderhub.ServiceFeatureFlagActionDoNotCreate),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceProviderManifest = armproviderhub.ResourceProviderManifest{
	// 	Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 		{
	// 			Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 			QuotaID: to.Ptr("CSP_2015-05-01"),
	// 		},
	// 		{
	// 			Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 			QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 	}},
	// 	CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
	// 	GlobalNotificationEndpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 		{
	// 			Enabled: to.Ptr(true),
	// 			EndpointURI: to.Ptr("https://notificationendpoint.com"),
	// 	}},
	// 	Management: &armproviderhub.ResourceProviderManifestManagement{
	// 		AuthorizationOwners: []*string{
	// 			to.Ptr("authorizationOwners-group")},
	// 			IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 			IncidentRoutingService: to.Ptr(""),
	// 			IncidentRoutingTeam: to.Ptr(""),
	// 			ManifestOwners: []*string{
	// 				to.Ptr("manifestOwners-group")},
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 			},
	// 			Metadata: map[string]any{
	// 				"onboardedVia": "ProviderHub",
	// 			},
	// 			Namespace: to.Ptr("microsoft.contoso"),
	// 			ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			}},
	// 			ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 			ProviderVersion: to.Ptr("2.0"),
	// 			ReRegisterSubscriptionMetadata: &armproviderhub.ResourceProviderManifestReRegisterSubscriptionMetadata{
	// 				ConcurrencyLimit: to.Ptr[int32](100),
	// 				Enabled: to.Ptr(true),
	// 			},
	// 			ResourceProviderAuthorizationRules: &armproviderhub.ResourceProviderAuthorizationRules{
	// 				AsyncOperationPollingRules: &armproviderhub.AsyncOperationPollingRules{
	// 					AuthorizationActions: []*string{
	// 						to.Ptr("Microsoft.Contoso/classicAdministrators/operationStatuses/read")},
	// 					},
	// 				},
	// 				ResourceTypes: []*armproviderhub.ResourceType{
	// 					{
	// 						Name: to.Ptr("Operations"),
	// 						AllowedUnauthorizedActions: []*string{
	// 							to.Ptr("microsoft.contoso/operations/read")},
	// 							AllowedUnauthorizedActionsExtensions: []*armproviderhub.AllowedUnauthorizedActionsExtension{
	// 								{
	// 									Action: to.Ptr("Microsoft.BizTalkServices/bizTalk/read"),
	// 									Intent: to.Ptr(armproviderhub.IntentDEFERREDACCESSCHECK),
	// 							}},
	// 							Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 								{
	// 									APIVersions: []*string{
	// 										to.Ptr("2020-01-01-preview")},
	// 										EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 										Locations: []*string{
	// 											to.Ptr("")},
	// 											Timeout: to.Ptr("PT20S"),
	// 									}},
	// 									LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 									},
	// 									ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 									RoutingType: to.Ptr(armproviderhub.RoutingType("ProxyOnly, Tenant")),
	// 								},
	// 								{
	// 									Name: to.Ptr("Locations"),
	// 									Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 										{
	// 											APIVersions: []*string{
	// 												to.Ptr("2020-01-01-preview")},
	// 												EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 												Locations: []*string{
	// 													to.Ptr("")},
	// 													Timeout: to.Ptr("PT20S"),
	// 											}},
	// 											LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 											},
	// 											ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 											RoutingType: to.Ptr(armproviderhub.RoutingTypeProxyOnly),
	// 										},
	// 										{
	// 											Name: to.Ptr("Locations/OperationStatuses"),
	// 											AdditionalOptions: to.Ptr(armproviderhub.AdditionalOptionsProtectedAsyncOperationPolling),
	// 											Endpoints: []*armproviderhub.ResourceProviderEndpoint{
	// 												{
	// 													APIVersions: []*string{
	// 														to.Ptr("2020-01-01-preview")},
	// 														EndpointURI: to.Ptr("https://resource-endpoint.com/"),
	// 														Locations: []*string{
	// 															to.Ptr("")},
	// 															Timeout: to.Ptr("PT20S"),
	// 													}},
	// 													LinkedOperationRules: []*armproviderhub.LinkedOperationRule{
	// 													},
	// 													ResourceValidation: to.Ptr(armproviderhub.ResourceValidation("ReservedWords, ProfaneWords")),
	// 													RoutingType: to.Ptr(armproviderhub.RoutingType("ProxyOnly, LocationBased")),
	// 											}},
	// 											ServiceName: to.Ptr("root"),
	// 											Services: []*armproviderhub.ResourceProviderService{
	// 												{
	// 													ServiceName: to.Ptr("tags"),
	// 													Status: to.Ptr(armproviderhub.ServiceStatusInactive),
	// 											}},
	// 										}
}
