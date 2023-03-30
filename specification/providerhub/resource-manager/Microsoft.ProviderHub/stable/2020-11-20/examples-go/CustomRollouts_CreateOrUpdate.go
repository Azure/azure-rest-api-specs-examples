package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/CustomRollouts_CreateOrUpdate.json
func ExampleCustomRolloutsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomRolloutsClient().CreateOrUpdate(ctx, "Microsoft.Contoso", "brazilUsShoeBoxTesting", armproviderhub.CustomRollout{
		Properties: &armproviderhub.CustomRolloutProperties{
			Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
				Canary: &armproviderhub.CustomRolloutSpecificationCanary{
					Regions: []*string{
						to.Ptr("brazilus")},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomRollout = armproviderhub.CustomRollout{
	// 	Name: to.Ptr("Microsoft.Contoso/brazilUsShoeBoxTesting"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/customRollouts"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/customRollouts/brazilUsShoeBoxTesting"),
	// 	Properties: &armproviderhub.CustomRolloutProperties{
	// 		ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		Specification: &armproviderhub.CustomRolloutPropertiesSpecification{
	// 			Canary: &armproviderhub.CustomRolloutSpecificationCanary{
	// 				Regions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 				},
	// 			},
	// 			Status: &armproviderhub.CustomRolloutPropertiesStatus{
	// 				CompletedRegions: []*string{
	// 					to.Ptr("brazilus"),
	// 					to.Ptr("eastus2euap"),
	// 					to.Ptr("centraluseuap")},
	// 				},
	// 			},
	// 		}
}
