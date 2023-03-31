package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/CustomRollouts_ListByProviderRegistration.json
func ExampleCustomRolloutsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.CustomRolloutArrayResponseWithContinuation = armproviderhub.CustomRolloutArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.CustomRollout{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso/canaryTesting99"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/canaryTesting99"),
		// 			Properties: &armproviderhub.CustomRolloutProperties{
		// 				ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 				Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
		// 					Canary: &armproviderhub.CustomRolloutSpecificationCanary{
		// 						Regions: []*string{
		// 							to.Ptr("eastus2euap"),
		// 							to.Ptr("centraluseuap")},
		// 						},
		// 					},
		// 					Status: &armproviderhub.CustomRolloutPropertiesStatus{
		// 						CompletedRegions: []*string{
		// 							to.Ptr("eastus2euap"),
		// 							to.Ptr("centraluseuap")},
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Microsoft.Contoso/brazilustesting"),
		// 					Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
		// 					ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/brazilustesting"),
		// 					Properties: &armproviderhub.CustomRolloutProperties{
		// 						ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 						Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
		// 							Canary: &armproviderhub.CustomRolloutSpecificationCanary{
		// 								Regions: []*string{
		// 									to.Ptr("brazilus")},
		// 								},
		// 							},
		// 							Status: &armproviderhub.CustomRolloutPropertiesStatus{
		// 								FailedOrSkippedRegions: map[string]*armproviderhub.ExtendedErrorInfo{
		// 									"brazilus": &armproviderhub.ExtendedErrorInfo{
		// 										Code: to.Ptr("RolloutTimedout"),
		// 										Message: to.Ptr("Failed to rollout to specified region."),
		// 									},
		// 								},
		// 							},
		// 						},
		// 				}},
		// 			}
	}
}
