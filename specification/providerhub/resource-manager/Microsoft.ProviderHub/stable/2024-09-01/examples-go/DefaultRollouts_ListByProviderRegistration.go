package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/DefaultRollouts_ListByProviderRegistration.json
func ExampleDefaultRolloutsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefaultRolloutsClient().NewListByProviderRegistrationPager("Microsoft.Contoso", nil)
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
		// page.DefaultRolloutArrayResponseWithContinuation = armproviderhub.DefaultRolloutArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.DefaultRollout{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso/2020week01"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/defaultRollouts"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/defaultRollouts/2020week01"),
		// 			SystemData: &armproviderhub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 			},
		// 			Properties: &armproviderhub.DefaultRolloutProperties{
		// 				ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 				Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
		// 					Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
		// 						Regions: []*string{
		// 							to.Ptr("brazilus"),
		// 							to.Ptr("eastus2euap"),
		// 							to.Ptr("centraluseuap")},
		// 						},
		// 						ExpeditedRollout: &armproviderhub.DefaultRolloutSpecificationExpeditedRollout{
		// 							Enabled: to.Ptr(true),
		// 						},
		// 						HighTraffic: &armproviderhub.DefaultRolloutSpecificationHighTraffic{
		// 							Regions: []*string{
		// 								to.Ptr("australiasoutheast"),
		// 								to.Ptr("otherhightraficregions")},
		// 								WaitDuration: to.Ptr("PT24H"),
		// 							},
		// 							LowTraffic: &armproviderhub.DefaultRolloutSpecificationLowTraffic{
		// 								Regions: []*string{
		// 									to.Ptr("southeastasia")},
		// 									WaitDuration: to.Ptr("PT24H"),
		// 								},
		// 								MediumTraffic: &armproviderhub.DefaultRolloutSpecificationMediumTraffic{
		// 									Regions: []*string{
		// 										to.Ptr("uksouth"),
		// 										to.Ptr("indiawest")},
		// 										WaitDuration: to.Ptr("PT24H"),
		// 									},
		// 									RestOfTheWorldGroupOne: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupOne{
		// 										Regions: []*string{
		// 											to.Ptr("koreacentral"),
		// 											to.Ptr("francecentral"),
		// 											to.Ptr("australiacentral"),
		// 											to.Ptr("westus"),
		// 											to.Ptr("allotherregions")},
		// 											WaitDuration: to.Ptr("PT4H"),
		// 										},
		// 										RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
		// 											Regions: []*string{
		// 												to.Ptr("germanynorth"),
		// 												to.Ptr("norwayeast"),
		// 												to.Ptr("allotherregions")},
		// 												WaitDuration: to.Ptr("PT4H"),
		// 											},
		// 										},
		// 										Status: &armproviderhub.DefaultRolloutPropertiesStatus{
		// 											CompletedRegions: []*string{
		// 												to.Ptr("brazilus"),
		// 												to.Ptr("eastus2euap"),
		// 												to.Ptr("centraluseuap"),
		// 												to.Ptr("allcompletedregions")},
		// 												ManifestCheckinStatus: &armproviderhub.DefaultRolloutStatusManifestCheckinStatus{
		// 													CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 													IsCheckedIn: to.Ptr(true),
		// 													StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
		// 												},
		// 											},
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("Microsoft.Contoso/2020week10"),
		// 										Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/defaultRollouts"),
		// 										ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/defaultRollouts/2020week10"),
		// 										SystemData: &armproviderhub.SystemData{
		// 											CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 											CreatedBy: to.Ptr("string"),
		// 											CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 											LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 											LastModifiedBy: to.Ptr("string"),
		// 											LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 										},
		// 										Properties: &armproviderhub.DefaultRolloutProperties{
		// 											ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 											Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
		// 												Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
		// 													Regions: []*string{
		// 														to.Ptr("brazilus"),
		// 														to.Ptr("eastus2euap"),
		// 														to.Ptr("centraluseuap")},
		// 													},
		// 													HighTraffic: &armproviderhub.DefaultRolloutSpecificationHighTraffic{
		// 														Regions: []*string{
		// 															to.Ptr("australiasoutheast"),
		// 															to.Ptr("otherhightraficregions")},
		// 															WaitDuration: to.Ptr("PT24H"),
		// 														},
		// 														LowTraffic: &armproviderhub.DefaultRolloutSpecificationLowTraffic{
		// 															Regions: []*string{
		// 																to.Ptr("southeastasia")},
		// 																WaitDuration: to.Ptr("PT24H"),
		// 															},
		// 															MediumTraffic: &armproviderhub.DefaultRolloutSpecificationMediumTraffic{
		// 																Regions: []*string{
		// 																	to.Ptr("uksouth"),
		// 																	to.Ptr("indiawest")},
		// 																	WaitDuration: to.Ptr("PT24H"),
		// 																},
		// 																RestOfTheWorldGroupOne: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupOne{
		// 																	Regions: []*string{
		// 																		to.Ptr("koreacentral"),
		// 																		to.Ptr("francecentral"),
		// 																		to.Ptr("australiacentral"),
		// 																		to.Ptr("westus"),
		// 																		to.Ptr("allotherregions")},
		// 																		WaitDuration: to.Ptr("PT4H"),
		// 																	},
		// 																	RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
		// 																		Regions: []*string{
		// 																			to.Ptr("germanynorth"),
		// 																			to.Ptr("norwayeast"),
		// 																			to.Ptr("allotherregions")},
		// 																			WaitDuration: to.Ptr("PT4H"),
		// 																		},
		// 																	},
		// 																	Status: &armproviderhub.DefaultRolloutPropertiesStatus{
		// 																		CompletedRegions: []*string{
		// 																			to.Ptr("brazilus"),
		// 																			to.Ptr("eastus2euap"),
		// 																			to.Ptr("centraluseuap"),
		// 																			to.Ptr("allcompletedregions")},
		// 																			FailedOrSkippedRegions: map[string]*armproviderhub.ExtendedErrorInfo{
		// 																				"westus2": &armproviderhub.ExtendedErrorInfo{
		// 																					Code: to.Ptr("RolloutStoppedByUser"),
		// 																					Message: to.Ptr("Rollout was explicitly stopped by the user."),
		// 																				},
		// 																			},
		// 																			ManifestCheckinStatus: &armproviderhub.DefaultRolloutStatusManifestCheckinStatus{
		// 																				CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 																				IsCheckedIn: to.Ptr(true),
		// 																				StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
		// 																			},
		// 																		},
		// 																	},
		// 															}},
		// 														}
	}
}
