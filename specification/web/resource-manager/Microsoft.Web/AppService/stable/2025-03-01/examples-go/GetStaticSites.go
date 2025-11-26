package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetStaticSites.json
func ExampleStaticSitesClient_NewGetStaticSitesByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStaticSitesClient().NewGetStaticSitesByResourceGroupPager("rg", nil)
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
		// page.StaticSiteCollection = armappservice.StaticSiteCollection{
		// 	Value: []*armappservice.StaticSiteARMResource{
		// 		{
		// 			Name: to.Ptr("testStaticSite0"),
		// 			Type: to.Ptr("Microsoft.Web/staticSites"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/staticSites/testStaticSite0"),
		// 			Location: to.Ptr("West US 2"),
		// 			Properties: &armappservice.StaticSite{
		// 				AllowConfigFileUpdates: to.Ptr(true),
		// 				Branch: to.Ptr("demo"),
		// 				ContentDistributionEndpoint: to.Ptr(""),
		// 				CustomDomains: []*string{
		// 				},
		// 				DefaultHostname: to.Ptr("happy-sea-15afae3e.azurestaticwebsites.net"),
		// 				KeyVaultReferenceIdentity: to.Ptr("SystemAssigned"),
		// 				LinkedBackends: []*armappservice.StaticSiteLinkedBackend{
		// 				},
		// 				PrivateEndpointConnections: []*armappservice.ResponseMessageEnvelopeRemotePrivateEndpointConnection{
		// 				},
		// 				RepositoryURL: to.Ptr("https://github.com/username/repo"),
		// 				StagingEnvironmentPolicy: to.Ptr(armappservice.StagingEnvironmentPolicyEnabled),
		// 			},
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("Standard"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
