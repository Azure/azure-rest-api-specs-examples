package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/NewRegionFrontloadRelease_Get.json
func ExampleNewRegionFrontloadReleaseClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNewRegionFrontloadReleaseClient().Get(ctx, "Microsoft.Contoso", "2020week10", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefaultRollout = armproviderhub.DefaultRollout{
	// 	Name: to.Ptr("Microsoft.Contoso/2020week10"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/defaultRollouts"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/defaultRollouts/2020week10"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.DefaultRolloutProperties{
	// 		ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
	// 			Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
	// 				Regions: []*string{
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 					SkipRegions: []*string{
	// 						to.Ptr("brazilus")},
	// 					},
	// 					ExpeditedRollout: &armproviderhub.DefaultRolloutSpecificationExpeditedRollout{
	// 						Enabled: to.Ptr(true),
	// 					},
	// 					HighTraffic: &armproviderhub.DefaultRolloutSpecificationHighTraffic{
	// 						Regions: []*string{
	// 							to.Ptr("australiasoutheast"),
	// 							to.Ptr("otherhightraficregions")},
	// 							WaitDuration: to.Ptr("PT24H"),
	// 						},
	// 						LowTraffic: &armproviderhub.DefaultRolloutSpecificationLowTraffic{
	// 							Regions: []*string{
	// 								to.Ptr("southeastasia")},
	// 								WaitDuration: to.Ptr("PT24H"),
	// 							},
	// 							MediumTraffic: &armproviderhub.DefaultRolloutSpecificationMediumTraffic{
	// 								Regions: []*string{
	// 									to.Ptr("uksouth"),
	// 									to.Ptr("indiawest")},
	// 									WaitDuration: to.Ptr("PT24H"),
	// 								},
	// 								ProviderRegistration: &armproviderhub.DefaultRolloutSpecificationProviderRegistration{
	// 									Name: to.Ptr("Microsoft.Contoso"),
	// 									Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations"),
	// 									ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso"),
	// 									Properties: &armproviderhub.ProviderRegistrationProperties{
	// 										Capabilities: []*armproviderhub.ResourceProviderCapabilities{
	// 											{
	// 												Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 												QuotaID: to.Ptr("CSP_2015-05-01"),
	// 											},
	// 											{
	// 												Effect: to.Ptr(armproviderhub.ResourceProviderCapabilitiesEffectAllow),
	// 												QuotaID: to.Ptr("CSP_MG_2017-12-01"),
	// 										}},
	// 										Management: &armproviderhub.ResourceProviderManifestPropertiesManagement{
	// 											AuthorizationOwners: []*string{
	// 												to.Ptr("RPAAS-PlatformServiceAdministrator")},
	// 												IncidentContactEmail: to.Ptr("helpme@contoso.com"),
	// 												IncidentRoutingService: to.Ptr(""),
	// 												IncidentRoutingTeam: to.Ptr(""),
	// 												ManifestOwners: []*string{
	// 													to.Ptr("SPARTA-PlatformServiceAdministrator")},
	// 													ResourceAccessPolicy: to.Ptr(armproviderhub.ResourceAccessPolicyNotSpecified),
	// 												},
	// 												Namespace: to.Ptr("microsoft.contoso"),
	// 												ProviderAuthorizations: []*armproviderhub.ResourceProviderAuthorization{
	// 													{
	// 														ApplicationID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 														RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 												}},
	// 												ProviderType: to.Ptr(armproviderhub.ResourceProviderType("Internal, Hidden")),
	// 												ProviderVersion: to.Ptr("2.0"),
	// 												ProviderHubMetadata: &armproviderhub.ProviderRegistrationPropertiesProviderHubMetadata{
	// 													ProviderAuthentication: &armproviderhub.MetadataProviderAuthentication{
	// 														AllowedAudiences: []*string{
	// 															to.Ptr("https://management.core.windows.net/")},
	// 														},
	// 													},
	// 													ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 												},
	// 											},
	// 											ResourceTypeRegistrations: []*armproviderhub.ResourceTypeRegistration{
	// 												{
	// 													Name: to.Ptr("Microsoft.Contoso/employees"),
	// 													Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/resourceTypeRegistrations"),
	// 													ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/resourceTypeRegistrations/employees"),
	// 													Properties: &armproviderhub.ResourceTypeRegistrationProperties{
	// 														EnableAsyncOperation: to.Ptr(false),
	// 														EnableThirdPartyS2S: to.Ptr(false),
	// 														Endpoints: []*armproviderhub.ResourceTypeEndpoint{
	// 															{
	// 																APIVersions: []*string{
	// 																	to.Ptr("2018-11-01-preview"),
	// 																	to.Ptr("2020-01-01-preview"),
	// 																	to.Ptr("2019-01-01")},
	// 																	Locations: []*string{
	// 																		to.Ptr("East Asia"),
	// 																		to.Ptr("East US"),
	// 																		to.Ptr("North Europe"),
	// 																		to.Ptr("Southeast Asia"),
	// 																		to.Ptr("East US 2 EUAP"),
	// 																		to.Ptr("Central US EUAP"),
	// 																		to.Ptr("West Europe"),
	// 																		to.Ptr("West US"),
	// 																		to.Ptr("West Central US"),
	// 																		to.Ptr("West US 2")},
	// 																		RequiredFeatures: []*string{
	// 																			to.Ptr("Microsoft.Contoso/RPaaSSampleApp")},
	// 																	}},
	// 																	ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 																	Regionality: to.Ptr(armproviderhub.RegionalityRegional),
	// 																	RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
	// 																	SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
	// 																		{
	// 																			APIVersions: []*string{
	// 																				to.Ptr("2018-11-01-preview"),
	// 																				to.Ptr("2020-01-01-preview"),
	// 																				to.Ptr("2019-01-01")},
	// 																				SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
	// 																		}},
	// 																	},
	// 															}},
	// 															RestOfTheWorldGroupOne: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupOne{
	// 																Regions: []*string{
	// 																	to.Ptr("koreacentral"),
	// 																	to.Ptr("francecentral"),
	// 																	to.Ptr("australiacentral"),
	// 																	to.Ptr("westus"),
	// 																	to.Ptr("allotherregions")},
	// 																	WaitDuration: to.Ptr("PT4H"),
	// 																},
	// 																RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
	// 																	Regions: []*string{
	// 																		to.Ptr("germanynorth"),
	// 																		to.Ptr("norwayeast"),
	// 																		to.Ptr("allotherregions")},
	// 																		WaitDuration: to.Ptr("PT4H"),
	// 																	},
	// 																},
	// 																Status: &armproviderhub.DefaultRolloutPropertiesStatus{
	// 																	CompletedRegions: []*string{
	// 																		to.Ptr("brazilus"),
	// 																		to.Ptr("eastus2euap"),
	// 																		to.Ptr("centraluseuap"),
	// 																		to.Ptr("allcompletedregions")},
	// 																		FailedOrSkippedRegions: map[string]*armproviderhub.ExtendedErrorInfo{
	// 																			"westus2": &armproviderhub.ExtendedErrorInfo{
	// 																				Code: to.Ptr("RolloutStoppedByUser"),
	// 																				Message: to.Ptr("Rollout was explicitly stopped by the user."),
	// 																			},
	// 																		},
	// 																		ManifestCheckinStatus: &armproviderhub.DefaultRolloutStatusManifestCheckinStatus{
	// 																			CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 																			IsCheckedIn: to.Ptr(true),
	// 																			StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 																		},
	// 																		SubscriptionReregistrationResult: to.Ptr(armproviderhub.SubscriptionReregistrationResultConditionalUpdate),
	// 																	},
	// 																},
	// 															}
}
