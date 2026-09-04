package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/CustomRollouts_Get.json
func ExampleCustomRolloutsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomRolloutsClient().Get(ctx, "Microsoft.Contoso", "canaryTesting99", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armproviderhub.CustomRolloutsClientGetResponse{
	// 	CustomRollout: armproviderhub.CustomRollout{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/canaryTesting99"),
	// 		Name: to.Ptr("Microsoft.Contoso/canaryTesting99"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 		Properties: &armproviderhub.CustomRolloutProperties{
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningState("Completed")),
	// 			Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
	// 				AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
	// 					Storage: to.Ptr(true),
	// 					ResourceGraph: to.Ptr(true),
	// 				},
	// 				Canary: &armproviderhub.CustomRolloutSpecificationCanary{
	// 					Regions: []*string{
	// 						to.Ptr("eastus2euap"),
	// 						to.Ptr("centraluseuap"),
	// 					},
	// 				},
	// 				RefreshSubscriptionRegistration: to.Ptr(true),
	// 				ProviderRegistration: &armproviderhub.CustomRolloutSpecificationProviderRegistration{
	// 					ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
	// 					Name: to.Ptr("Microsoft.Contoso"),
	// 					Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
	// 					Properties: &armproviderhub.ProviderRegistrationProperties{
	// 						ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
	// 							ProviderAuthentication: &armproviderhub.MetadataProviderAuthentication{
	// 								AllowedAudiences: []*string{
	// 									to.Ptr("https://management.core.windows.net/"),
	// 								},
	// 							},
	// 						},
	// 						ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 						ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 							{
	// 								ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 								RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 							},
	// 						},
	// 						Namespace: to.Ptr("microsoft.contoso"),
	// 						ProviderVersion: to.Ptr("2.0"),
	// 						ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 						Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 							ManifestOwners: []*string{
	// 								to.Ptr("Contoso-PlatformServiceAdministrator"),
	// 							},
	// 							AuthorizationOwners: []*string{
	// 								to.Ptr("RPAAS-PlatformServiceAdministrator"),
	// 							},
	// 							IncidentRoutingService: to.Ptr(""),
	// 							IncidentRoutingTeam: to.Ptr(""),
	// 							IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 							ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 						},
	// 						Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 							{
	// 								QuotaID: to.Ptr("CSP_2015-05-01"),
	// 								Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 							},
	// 							{
	// 								QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 								Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 							},
	// 						},
	// 						Metadata: nil,
	// 					},
	// 				},
	// 				ResourceTypeRegistrations: []*armproviderhub.ResourceTypeRegistration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourceTypeRegistrations/employees"),
	// 						Name: to.Ptr("Microsoft.Contoso/employees"),
	// 						Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourceTypeRegistrations"),
	// 						Properties: &armproviderhub.ResourceTypeRegistrationProperties{
	// 							RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
	// 							Regionality: to.Ptr(armproviderhub.RegionalityRegional),
	// 							Endpoints: []*armproviderhub.ResourceTypeEndpoint{
	// 								{
	// 									APIVersions: []*string{
	// 										to.Ptr("2018-11-01-preview"),
	// 										to.Ptr("2020-01-01-preview"),
	// 										to.Ptr("2019-01-01"),
	// 									},
	// 									Locations: []*string{
	// 										to.Ptr("East Asia"),
	// 										to.Ptr("East US"),
	// 										to.Ptr("North Europe"),
	// 										to.Ptr("Southeast Asia"),
	// 										to.Ptr("East US 2 EUAP"),
	// 										to.Ptr("Central US EUAP"),
	// 										to.Ptr("West Europe"),
	// 										to.Ptr("West US"),
	// 										to.Ptr("West Central US"),
	// 										to.Ptr("West US 2"),
	// 									},
	// 									RequiredFeatures: []*string{
	// 										to.Ptr("Microsoft.Contoso/RPaaSSampleApp"),
	// 									},
	// 								},
	// 							},
	// 							SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
	// 								{
	// 									APIVersions: []*string{
	// 										to.Ptr("2018-11-01-preview"),
	// 										to.Ptr("2020-01-01-preview"),
	// 										to.Ptr("2019-01-01"),
	// 									},
	// 									SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
	// 								},
	// 							},
	// 							EnableAsyncOperation: to.Ptr(false),
	// 							ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 							EnableThirdPartyS2S: to.Ptr(false),
	// 						},
	// 					},
	// 				},
	// 				ManifestCheckinSpecification: &armproviderhub.ManifestCheckinSpecification{
	// 					ManifestCheckinOption: to.Ptr(armproviderhub.ManifestCheckinOptionAttemptAutomaticManifestCheckin),
	// 					ManifestCheckinParams: &armproviderhub.CheckinManifestParams{
	// 						BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
	// 						Environment: to.Ptr("Prod"),
	// 					},
	// 				},
	// 			},
	// 			Status: &armproviderhub.CustomRolloutPropertiesStatus{
	// 				CompletedRegions: []*string{
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap"),
	// 				},
	// 				CompletedRegionsInfo: []*armproviderhub.AppliedManifestInfo{
	// 					{
	// 						Region: to.Ptr("eastus2euap"),
	// 						ManifestAppliedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 						PreviousCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 						AppliedCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 					},
	// 					{
	// 						Region: to.Ptr("centraluseuap"),
	// 						ManifestAppliedAt: to.Ptr(time.Date(2020, time.February, 1, 9, 1, 1, 107505600, time.UTC)),
	// 						PreviousCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 						AppliedCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 					},
	// 				},
	// 				ManifestCheckinStatus: &armproviderhub.CustomRolloutStatusManifestCheckinStatus{
	// 					IsCheckedIn: to.Ptr(true),
	// 					StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 					CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
