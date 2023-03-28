package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProvider.json
func ExampleProvidersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProvidersClient().Get(ctx, "Microsoft.TestRP1", &armresources.ProvidersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Provider = armresources.Provider{
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.TestRP1"),
	// 	Namespace: to.Ptr("Microsoft.TestRP1"),
	// 	RegistrationPolicy: to.Ptr("RegistrationRequired"),
	// 	RegistrationState: to.Ptr("Registering"),
	// 	ResourceTypes: []*armresources.ProviderResourceType{
	// 		{
	// 			APIVersions: []*string{
	// 				to.Ptr("2018-01-01"),
	// 				to.Ptr("2015-05-01")},
	// 				Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
	// 				LocationMappings: []*armresources.ProviderExtendedLocation{
	// 					{
	// 						Type: to.Ptr("EdgeZone"),
	// 						ExtendedLocations: []*string{
	// 							to.Ptr("LosAngeles"),
	// 							to.Ptr("LosAngeles2")},
	// 							Location: to.Ptr("West US"),
	// 					}},
	// 					Locations: []*string{
	// 						to.Ptr("West US")},
	// 						ResourceType: to.Ptr("TestResourceType"),
	// 					},
	// 					{
	// 						APIVersions: []*string{
	// 							to.Ptr("2018-01-01"),
	// 							to.Ptr("2015-05-01")},
	// 							Capabilities: to.Ptr("CrossResourceGroupResourceMove, CrossSubscriptionResourceMove, SupportsTags, SupportsLocation"),
	// 							Locations: []*string{
	// 								to.Ptr("West US")},
	// 								ResourceType: to.Ptr("TestResourceTypeSibling"),
	// 						}},
	// 					}
}
