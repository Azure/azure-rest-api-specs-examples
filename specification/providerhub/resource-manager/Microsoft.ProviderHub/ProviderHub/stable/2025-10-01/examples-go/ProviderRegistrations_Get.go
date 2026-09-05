package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/ProviderRegistrations_Get.json
func ExampleProviderRegistrationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderRegistrationsClient().Get(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.ProviderRegistrationsClientGetResponse{
	// 	ProviderRegistration: armproviderhub.ProviderRegistration{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
	// 		Name: to.Ptr("Microsoft.Contoso"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
	// 		Properties: &armproviderhub.ProviderRegistrationProperties{
	// 			ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
	// 				ProviderAuthentication: &armproviderhub.MetadataProviderAuthentication{
	// 					AllowedAudiences: []*string{
	// 						to.Ptr("https://management.core.windows.net/"),
	// 					},
	// 				},
	// 				DirectRpRoleDefinitionID: to.Ptr("1x86y807-6zx0-40y3-8z5x-686y7z43x0y2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 				{
	// 					ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 					RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 				},
	// 			},
	// 			ResourceProviderAuthorizationRules: &armproviderhub.ResourceProviderAuthorizationRules{
	// 				AsyncOperationPollingRules: &armproviderhub.AsyncOperationPollingRules{
	// 					AuthorizationActions: []*string{
	// 						to.Ptr("Microsoft.Contoso/classicAdministrators/operationStatuses/read"),
	// 					},
	// 				},
	// 			},
	// 			Namespace: to.Ptr("microsoft.contoso"),
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
	// 			Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 				ManifestOwners: []*string{
	// 					to.Ptr("Contoso-PlatformServiceAdministrator"),
	// 				},
	// 				AuthorizationOwners: []*string{
	// 					to.Ptr("RPAAS-PlatformServiceAdministrator"),
	// 				},
	// 				IncidentRoutingService: to.Ptr(""),
	// 				IncidentRoutingTeam: to.Ptr(""),
	// 				IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 				ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
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
	// 			Metadata: nil,
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
