package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ResourceTypeRegistrations_ListByProviderRegistration.json
func ExampleResourceTypeRegistrationsClient_NewListByProviderRegistrationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceTypeRegistrationsClient().NewListByProviderRegistrationPager("Microsoft.Contoso", nil)
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
		// page.ResourceTypeRegistrationArrayResponseWithContinuation = armproviderhub.ResourceTypeRegistrationArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.ResourceTypeRegistration{
		// 		{
		// 			Name: to.Ptr("employees"),
		// 			Properties: &armproviderhub.ResourceTypeRegistrationProperties{
		// 				EnableAsyncOperation: to.Ptr(false),
		// 				EnableThirdPartyS2S: to.Ptr(false),
		// 				Endpoints: []*armproviderhub.ResourceTypeEndpoint{
		// 					{
		// 						APIVersions: []*string{
		// 							to.Ptr("2018-11-01-preview"),
		// 							to.Ptr("2020-01-01-preview"),
		// 							to.Ptr("2019-01-01")},
		// 							Locations: []*string{
		// 								to.Ptr("East Asia"),
		// 								to.Ptr("East US"),
		// 								to.Ptr("North Europe"),
		// 								to.Ptr("Southeast Asia"),
		// 								to.Ptr("East US 2 EUAP"),
		// 								to.Ptr("Central US EUAP"),
		// 								to.Ptr("West Europe"),
		// 								to.Ptr("West US"),
		// 								to.Ptr("West Central US"),
		// 								to.Ptr("West US 2")},
		// 								RequiredFeatures: []*string{
		// 									to.Ptr("Microsoft.Contoso/RPaaSSampleApp")},
		// 							}},
		// 							ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 							Regionality: to.Ptr(armproviderhub.RegionalityRegional),
		// 							RoutingType: to.Ptr(armproviderhub.RoutingTypeDefault),
		// 							SwaggerSpecifications: []*armproviderhub.SwaggerSpecification{
		// 								{
		// 									APIVersions: []*string{
		// 										to.Ptr("2018-11-01-preview"),
		// 										to.Ptr("2020-01-01-preview"),
		// 										to.Ptr("2019-01-01")},
		// 										SwaggerSpecFolderURI: to.Ptr("https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/"),
		// 								}},
		// 							},
		// 					}},
		// 				}
	}
}
