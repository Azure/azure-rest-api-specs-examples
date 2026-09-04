package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/DefaultRollouts_CreateOrUpdate.json
func ExampleDefaultRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDefaultRolloutsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", "2020week10", armproviderhub.DefaultRollout{
		Properties: &armproviderhub.DefaultRolloutProperties{
			Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
				ExpeditedRollout: &armproviderhub.DefaultRolloutSpecificationExpeditedRollout{
					Enabled: to.Ptr(true),
				},
				Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
					SkipRegions: []*string{
						to.Ptr("eastus2euap"),
					},
				},
				RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
					WaitDuration: to.Ptr("PT4H"),
				},
				ManifestCheckinSpecification: &armproviderhub.ManifestCheckinSpecification{
					ManifestCheckinOption: to.Ptr(armproviderhub.ManifestCheckinOptionAttemptAutomaticManifestCheckin),
					ManifestCheckinParams: &armproviderhub.CheckinManifestParams{
						BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
						Environment:                 to.Ptr("Prod"),
					},
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
	// res = armproviderhub.DefaultRolloutsClientCreateOrUpdateResponse{
	// 	DefaultRollout: armproviderhub.DefaultRollout{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/defaultRollouts/2020week10"),
	// 		Name: to.Ptr("Microsoft.Contoso/2020week10"),
	// 		Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/defaultRollouts"),
	// 		SystemData: &armproviderhub.SystemData{
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
	// 		},
	// 		Properties: &armproviderhub.DefaultRolloutProperties{
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateRolloutInProgress),
	// 			Specification: &armproviderhub.DefaultRolloutPropertiesSpecification{
	// 				Canary: &armproviderhub.DefaultRolloutSpecificationCanary{
	// 					Regions: []*string{
	// 						to.Ptr("brazilus"),
	// 						to.Ptr("eastus2euap"),
	// 						to.Ptr("centraluseuap"),
	// 					},
	// 				},
	// 				LowTraffic: &armproviderhub.DefaultRolloutSpecificationLowTraffic{
	// 					Regions: []*string{
	// 						to.Ptr("southeastasia"),
	// 					},
	// 					WaitDuration: to.Ptr("PT24H"),
	// 				},
	// 				MediumTraffic: &armproviderhub.DefaultRolloutSpecificationMediumTraffic{
	// 					Regions: []*string{
	// 						to.Ptr("uksouth"),
	// 						to.Ptr("indiawest"),
	// 					},
	// 					WaitDuration: to.Ptr("PT24H"),
	// 				},
	// 				HighTraffic: &armproviderhub.DefaultRolloutSpecificationHighTraffic{
	// 					Regions: []*string{
	// 						to.Ptr("australiasoutheast"),
	// 						to.Ptr("otherhightraficregions"),
	// 					},
	// 					WaitDuration: to.Ptr("PT24H"),
	// 				},
	// 				RestOfTheWorldGroupOne: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupOne{
	// 					Regions: []*string{
	// 						to.Ptr("koreacentral"),
	// 						to.Ptr("francecentral"),
	// 						to.Ptr("australiacentral"),
	// 						to.Ptr("westus"),
	// 						to.Ptr("allotherregions"),
	// 					},
	// 					WaitDuration: to.Ptr("PT4H"),
	// 				},
	// 				RestOfTheWorldGroupTwo: &armproviderhub.DefaultRolloutSpecificationRestOfTheWorldGroupTwo{
	// 					Regions: []*string{
	// 						to.Ptr("germanynorth"),
	// 						to.Ptr("norwayeast"),
	// 						to.Ptr("allotherregions"),
	// 					},
	// 					WaitDuration: to.Ptr("PT4H"),
	// 				},
	// 				ManifestCheckinSpecification: &armproviderhub.ManifestCheckinSpecification{
	// 					ManifestCheckinOption: to.Ptr(armproviderhub.ManifestCheckinOptionAttemptAutomaticManifestCheckin),
	// 					ManifestCheckinParams: &armproviderhub.CheckinManifestParams{
	// 						BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
	// 						Environment: to.Ptr("Prod"),
	// 					},
	// 				},
	// 			},
	// 			Status: &armproviderhub.DefaultRolloutPropertiesStatus{
	// 				CompletedRegions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap"),
	// 					to.Ptr("allcompletedregions"),
	// 				},
	// 				ManifestCheckinStatus: &armproviderhub.DefaultRolloutStatusManifestCheckinStatus{
	// 					IsCheckedIn: to.Ptr(true),
	// 					StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 					CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
