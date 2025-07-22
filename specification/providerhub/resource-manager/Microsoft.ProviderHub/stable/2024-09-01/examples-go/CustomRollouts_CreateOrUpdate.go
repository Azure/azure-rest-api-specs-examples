package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/CustomRollouts_CreateOrUpdate.json
func ExampleCustomRolloutsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomRolloutsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", "brazilUsShoeBoxTesting", armproviderhub.CustomRollout{
		Properties: &armproviderhub.CustomRolloutProperties{
			Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
				AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
					ResourceGraph: to.Ptr(true),
					Storage:       to.Ptr(true),
				},
				Canary: &armproviderhub.CustomRolloutSpecificationCanary{
					Regions: []*string{
						to.Ptr("brazilus")},
				},
				RefreshSubscriptionRegistration: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomRollout = armproviderhub.CustomRollout{
	// 	Name: to.Ptr("Microsoft.Contoso/brazilUsShoeBoxTesting"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/brazilUsShoeBoxTesting"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.CustomRolloutProperties{
	// 		ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
	// 			AutoProvisionConfig: &armproviderhub.CustomRolloutSpecificationAutoProvisionConfig{
	// 				ResourceGraph: to.Ptr(true),
	// 				Storage: to.Ptr(true),
	// 			},
	// 			Canary: &armproviderhub.CustomRolloutSpecificationCanary{
	// 				Regions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 				},
	// 				RefreshSubscriptionRegistration: to.Ptr(true),
	// 			},
	// 			Status: &armproviderhub.CustomRolloutPropertiesStatus{
	// 				CompletedRegions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 					ManifestCheckinStatus: &armproviderhub.CustomRolloutStatusManifestCheckinStatus{
	// 						CommitID: to.Ptr("47317892d4edf22f08704f6b595105c4fd7a8db7"),
	// 						IsCheckedIn: to.Ptr(true),
	// 						StatusMessage: to.Ptr("Manifest is successfully merged. Use the Default/Custom rollout (http://aka.ms/rpaasrollout) to roll out the manifest in ARM."),
	// 					},
	// 				},
	// 			},
	// 		}
}
