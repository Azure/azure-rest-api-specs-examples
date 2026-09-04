package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/CustomRollouts_CreateOrUpdate.json
func ExampleCustomRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomRolloutsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", "brazilUsShoeBoxTesting", armproviderhub.CustomRollout{
		Properties: &armproviderhub.CustomRolloutProperties{
			Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
				AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
					Storage:       to.Ptr(true),
					ResourceGraph: to.Ptr(true),
				},
				Canary: &armproviderhub.CustomRolloutSpecificationCanary{
					Regions: []*string{
						to.Ptr("brazilus"),
					},
				},
				ManifestCheckinSpecification: &armproviderhub.ManifestCheckinSpecification{
					ManifestCheckinOption: to.Ptr(armproviderhub.ManifestCheckinOptionAttemptAutomaticManifestCheckin),
					ManifestCheckinParams: &armproviderhub.CheckinManifestParams{
						BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
						Environment:                 to.Ptr("Prod"),
					},
				},
				RefreshSubscriptionRegistration: to.Ptr(true),
				RolloutID:                       to.Ptr("Ev2RolloutIdGuid"),
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
	// res = armproviderhub.CustomRolloutsClientCreateOrUpdateResponse{
	// 	CustomRollout: armproviderhub.CustomRollout{
	// 		ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/brazilUsShoeBoxTesting"),
	// 		Name: to.Ptr("Microsoft.Contoso/brazilUsShoeBoxTesting"),
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
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 			Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
	// 				AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
	// 					Storage: to.Ptr(true),
	// 					ResourceGraph: to.Ptr(true),
	// 				},
	// 				Canary: &armproviderhub.CustomRolloutSpecificationCanary{
	// 					Regions: []*string{
	// 						to.Ptr("brazilus"),
	// 						to.Ptr("eastus2euap"),
	// 						to.Ptr("centraluseuap"),
	// 					},
	// 				},
	// 				ManifestCheckinSpecification: &armproviderhub.ManifestCheckinSpecification{
	// 					ManifestCheckinOption: to.Ptr(armproviderhub.ManifestCheckinOptionAttemptAutomaticManifestCheckin),
	// 					ManifestCheckinParams: &armproviderhub.CheckinManifestParams{
	// 						BaselineArmManifestLocation: to.Ptr("EastUS2EUAP"),
	// 						Environment: to.Ptr("Prod"),
	// 					},
	// 				},
	// 				RefreshSubscriptionRegistration: to.Ptr(true),
	// 				RolloutID: to.Ptr("Ev2RolloutIdGuid"),
	// 			},
	// 			Status: &armproviderhub.CustomRolloutPropertiesStatus{
	// 				CompletedRegions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap"),
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
