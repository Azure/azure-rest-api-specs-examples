package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/ProviderRegistrations_CreateOrUpdate.json
func ExampleProviderRegistrationsClient_BeginCreateOrUpdate_providerRegistrationsCreateOrUpdate() {
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
			CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
				IncidentRoutingService: to.Ptr("Contoso Resource Provider"),
				IncidentRoutingTeam:    to.Ptr("Contoso Triage"),
				IncidentContactEmail:   to.Ptr("helpme@contoso.com"),
				ExpeditedRolloutSubmitters: []*string{
					to.Ptr("Contoso-PlatformServiceOperator"),
				},
				ExpeditedRolloutMetadata: &armproviderhub.ResourceProviderManagementExpeditedRolloutMetadata{
					Enabled:                to.Ptr(false),
					ExpeditedRolloutIntent: to.Ptr(armproviderhub.ExpeditedRolloutIntentHotfix),
				},
				ErrorResponseMessageOptions: &armproviderhub.ResourceProviderManagementErrorResponseMessageOptions{
					ServerFailureResponseMessageType: to.Ptr(armproviderhub.ServerFailureResponseMessageTypeOutageReporting),
				},
				CanaryManifestOwners: []*string{
					to.Ptr("Contoso-PlatformServiceAdmin"),
				},
				PcCode:                to.Ptr("P1234"),
				ProfitCenterProgramID: to.Ptr("1234"),
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
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k2"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b3"),
	// 					ManagedByAuthorization: &armproviderhub.ResourceProviderAuthorizationManagedByAuthorization{
	// 						ManagedByResourceRoleDefinitionID: to.Ptr("9e3af657-a8ff-583c-a75c-2fe7c4bcb635"),
	// 						AllowManagedByInheritance: to.Ptr(true),
	// 						AdditionalAuthorizations: []*armproviderhub.AdditionalAuthorization{
	// 							{
	// 								ApplicationID: to.Ptr("3e5aaca6-6470-4be4-8a17-24ab9519414b"),
	// 								RoleDefinitionID: to.Ptr("1e86f807-6ec0-40b3-8b5f-686b7e43a0a2"),
	// 							},
	// 						},
	// 					},
	// 					AllowedThirdPartyExtensions: []*armproviderhub.ThirdPartyExtension{
	// 						{
	// 							Name: to.Ptr("name"),
	// 						},
	// 					},
	// 					GroupingTag: to.Ptr("GroupingTag"),
	// 				},
	// 			},
	// 			ResourceProviderAuthorizationRules: &armproviderhub.ResourceProviderAuthorizationRules{
	// 				AsyncOperationPollingRules: &armproviderhub.AsyncOperationPollingRules{
	// 					AuthorizationActions: []*string{
	// 						to.Ptr("Microsoft.Contoso/classicAdministrators/operationStatuses/read"),
	// 					},
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
	// 				ExpeditedRolloutSubmitters: []*string{
	// 					to.Ptr("Contoso-PlatformServiceOperator"),
	// 				},
	// 				ExpeditedRolloutMetadata: &armproviderhub.ResourceProviderManagementExpeditedRolloutMetadata{
	// 					Enabled: to.Ptr(false),
	// 					ExpeditedRolloutIntent: to.Ptr(armproviderhub.ExpeditedRolloutIntentHotfix),
	// 				},
	// 				ErrorResponseMessageOptions: &armproviderhub.ResourceProviderManagementErrorResponseMessageOptions{
	// 					ServerFailureResponseMessageType: to.Ptr(armproviderhub.ServerFailureResponseMessageTypeOutageReporting),
	// 				},
	// 				CanaryManifestOwners: []*string{
	// 					to.Ptr("Contoso-PlatformServiceAdmin"),
	// 				},
	// 				PcCode: to.Ptr("P1234"),
	// 				ProfitCenterProgramID: to.Ptr("1234"),
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
	// 			CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
	// 				DirectRpRoleDefinitionID: to.Ptr("1x86y807-6zx0-40y3-8z5x-686y7z43x0y2"),
	// 			},
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
