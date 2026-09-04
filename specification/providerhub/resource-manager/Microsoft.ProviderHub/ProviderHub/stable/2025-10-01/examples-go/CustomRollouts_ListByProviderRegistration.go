package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v4"
)

// Generated from example definition: 2025-10-01/CustomRollouts_ListByProviderRegistration.json
func ExampleCustomRolloutsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("ab7a8701-f7ef-471a-a2f4-d0ebbf494f77", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomRolloutsClient().NewListByProviderRegistrationPager("Microsoft.Contoso", nil)
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
		// page = armproviderhub.CustomRolloutsClientListByProviderRegistrationResponse{
		// 	CustomRolloutArrayResponseWithContinuation: armproviderhub.CustomRolloutArrayResponseWithContinuation{
		// 		Value: []*armproviderhub.CustomRollout{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/canaryTesting99"),
		// 				Name: to.Ptr("Microsoft.Contoso/canaryTesting99"),
		// 				Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
		// 				SystemData: &armproviderhub.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				Properties: &armproviderhub.CustomRolloutProperties{
		// 					ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 					Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
		// 						AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
		// 							Storage: to.Ptr(true),
		// 							ResourceGraph: to.Ptr(true),
		// 						},
		// 						Canary: &armproviderhub.CustomRolloutSpecificationCanary{
		// 							Regions: []*string{
		// 								to.Ptr("eastus2euap"),
		// 								to.Ptr("centraluseuap"),
		// 							},
		// 						},
		// 						RefreshSubscriptionRegistration: to.Ptr(true),
		// 						RolloutID: to.Ptr("Ev2RolloutIdGuid"),
		// 					},
		// 					Status: &armproviderhub.CustomRolloutPropertiesStatus{
		// 						CompletedRegions: []*string{
		// 							to.Ptr("eastus2euap"),
		// 							to.Ptr("centraluseuap"),
		// 						},
		// 						CompletedRegionsInfo: []*armproviderhub.AppliedManifestInfo{
		// 							{
		// 								Region: to.Ptr("eastus2euap"),
		// 								ManifestAppliedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 								PreviousCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 								AppliedCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 							},
		// 							{
		// 								Region: to.Ptr("centraluseuap"),
		// 								ManifestAppliedAt: to.Ptr(time.Date(2020, time.February, 1, 9, 1, 1, 107505600, time.UTC)),
		// 								PreviousCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 								AppliedCommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 							},
		// 						},
		// 						ManifestCheckinStatus: &armproviderhub.CustomRolloutStatusManifestCheckinStatus{
		// 							IsCheckedIn: to.Ptr(true),
		// 							StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
		// 							CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/brazilustesting"),
		// 				Name: to.Ptr("Microsoft.Contoso/brazilustesting"),
		// 				Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
		// 				SystemData: &armproviderhub.SystemData{
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2020, time.February, 1, 1, 1, 1, 107505600, time.UTC)),
		// 				},
		// 				Properties: &armproviderhub.CustomRolloutProperties{
		// 					ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateFailed),
		// 					Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
		// 						AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
		// 							Storage: to.Ptr(true),
		// 							ResourceGraph: to.Ptr(true),
		// 						},
		// 						Canary: &armproviderhub.CustomRolloutSpecificationCanary{
		// 							Regions: []*string{
		// 								to.Ptr("brazilus"),
		// 							},
		// 						},
		// 						RefreshSubscriptionRegistration: to.Ptr(true),
		// 						RolloutID: to.Ptr("Ev2RolloutIdGuid2"),
		// 					},
		// 					Status: &armproviderhub.CustomRolloutPropertiesStatus{
		// 						FailedOrSkippedRegions: map[string]*armproviderhub.ExtendedErrorInfo{
		// 							"brazilus": &armproviderhub.ExtendedErrorInfo{
		// 								Code: to.Ptr("RolloutTimedout"),
		// 								Message: to.Ptr("Failed to rollout to specified region."),
		// 							},
		// 						},
		// 						ManifestCheckinStatus: &armproviderhub.CustomRolloutStatusManifestCheckinStatus{
		// 							IsCheckedIn: to.Ptr(true),
		// 							StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
		// 							CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
