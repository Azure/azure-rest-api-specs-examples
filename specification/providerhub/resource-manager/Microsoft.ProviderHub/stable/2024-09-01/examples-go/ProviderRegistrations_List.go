package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/ProviderRegistrations_List.json
func ExampleProviderRegistrationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderRegistrationsClient().NewListPager(nil)
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
		// page.ProviderRegistrationArrayResponseWithContinuation = armproviderhub.ProviderRegistrationArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.ProviderRegistration{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
		// 			SystemData: &armproviderhub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 			},
		// 			Properties: &armproviderhub.ProviderRegistrationProperties{
		// 				Capabilities: []*armproviderhub.ResourceProviderCapabilities{
		// 					{
		// 						Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
		// 						QuotaID: to.Ptr("CSP_2015-05-01"),
		// 					},
		// 					{
		// 						Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
		// 						QuotaID: to.Ptr("CSP_MG_2017-12-01"),
		// 				}},
		// 				CrossTenantTokenValidation: to.Ptr(armproviderhub.CrossTenantTokenValidationEnsureSecureValidation),
		// 				Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
		// 					AuthorizationOwners: []*string{
		// 						to.Ptr("authorizationOwners-group")},
		// 						IncidentContactEmail: to.Ptr("helpme@contoso.com"),
		// 						IncidentRoutingService: to.Ptr(""),
		// 						IncidentRoutingTeam: to.Ptr(""),
		// 						ManifestOwners: []*string{
		// 							to.Ptr("manifestOwners-group")},
		// 							PcCode: to.Ptr("P1234"),
		// 							ProfitCenterProgramID: to.Ptr("1234"),
		// 							ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
		// 							ServiceTreeInfos: []*armproviderhub.ServiceTreeInfo{
		// 								{
		// 									ComponentID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
		// 									Readiness: to.Ptr(armproviderhub.ReadinessInDevelopment),
		// 									ServiceID: to.Ptr("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69"),
		// 							}},
		// 						},
		// 						Namespace: to.Ptr("microsoft.contoso"),
		// 						ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
		// 							{
		// 								ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
		// 								RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
		// 						}},
		// 						ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
		// 						ProviderVersion: to.Ptr("2.0"),
		// 						ResourceProviderAuthorizationRules: &armproviderhub.ResourceProviderAuthorizationRules{
		// 							AsyncOperationPollingRules: &armproviderhub.AsyncOperationPollingRules{
		// 								AuthorizationActions: []*string{
		// 									to.Ptr("Microsoft.Contoso/classicAdministrators/operationStatuses/read")},
		// 								},
		// 							},
		// 							ServiceName: to.Ptr("root"),
		// 							Services: []*armproviderhub.ResourceProviderService{
		// 								{
		// 									ServiceName: to.Ptr("tags"),
		// 									Status: to.Ptr(armproviderhub.ServiceStatusInactive),
		// 							}},
		// 							ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
		// 								DirectRpRoleDefinitionID: to.Ptr("1x86y807-6zx0-40y3-8z5x-686y7z43x0y2"),
		// 								ProviderAuthentication: &armproviderhub.MetadataProviderAuthentication{
		// 									AllowedAudiences: []*string{
		// 										to.Ptr("https://management.core.windows.net/")},
		// 									},
		// 								},
		// 								ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 							},
		// 					}},
		// 				}
	}
}
