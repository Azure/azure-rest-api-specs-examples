package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/NewRegionFrontloadRelease_CreateOrUpdate.json
func ExampleNewRegionFrontloadReleaseClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNewRegionFrontloadReleaseClient().CreateOrUpdate(ctx, "Microsoft.Contoso", "2020week10", armproviderhub.FrontloadPayload{
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
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 				},
	// 				HighTraffic: &armproviderhub.DefaultRolloutSpecificationHighTraffic{
	// 					Regions: []*string{
	// 						to.Ptr("australiasoutheast"),
	// 						to.Ptr("otherhightraficregions")},
	// 						WaitDuration: to.Ptr("PT24H"),
	// 					},
	// 					LowTraffic: &armproviderhub.DefaultRolloutSpecificationLowTraffic{
	// 						Regions: []*string{
	// 							to.Ptr("southeastasia")},
	// 							WaitDuration: to.Ptr("PT24H"),
	// 						},
	// 						MediumTraffic: &armproviderhub.DefaultRolloutSpecificationMediumTraffic{
	// 							Regions: []*string{
	// 								to.Ptr("uksouth"),
	// 								to.Ptr("indiawest")},
	// 								WaitDuration: to.Ptr("PT24H"),
	// 							},
	// 							RestOfTheWorldGroupOne: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupOne{
	// 								Regions: []*string{
	// 									to.Ptr("koreacentral"),
	// 									to.Ptr("francecentral"),
	// 									to.Ptr("australiacentral"),
	// 									to.Ptr("westus"),
	// 									to.Ptr("allotherregions")},
	// 									WaitDuration: to.Ptr("PT4H"),
	// 								},
	// 								RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
	// 									Regions: []*string{
	// 										to.Ptr("germanynorth"),
	// 										to.Ptr("norwayeast"),
	// 										to.Ptr("allotherregions")},
	// 										WaitDuration: to.Ptr("PT4H"),
	// 									},
	// 								},
	// 								Status: &armproviderhub.DefaultRolloutPropertiesStatus{
	// 									CompletedRegions: []*string{
	// 										to.Ptr("brazilus"),
	// 										to.Ptr("eastus2euap"),
	// 										to.Ptr("centraluseuap"),
	// 										to.Ptr("allcompletedregions")},
	// 										ManifestCheckinStatus: &armproviderhub.DefaultRolloutStatusManifestCheckinStatus{
	// 											CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 											IsCheckedIn: to.Ptr(true),
	// 											StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 										},
	// 									},
	// 								},
	// 							}
}
