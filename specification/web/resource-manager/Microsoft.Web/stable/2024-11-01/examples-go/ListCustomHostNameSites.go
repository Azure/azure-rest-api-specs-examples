package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/ListCustomHostNameSites.json
func ExampleWebSiteManagementClient_NewListCustomHostNameSitesPager_getCustomHostnamesUnderSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebSiteManagementClient().NewListCustomHostNameSitesPager(&armappservice.WebSiteManagementClientListCustomHostNameSitesOptions{Hostname: nil})
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
		// page.CustomHostnameSitesCollection = armappservice.CustomHostnameSitesCollection{
		// 	Value: []*armappservice.CustomHostnameSites{
		// 		{
		// 			Name: to.Ptr("mywebapp.azurewebsites.net"),
		// 			Type: to.Ptr("Microsoft.Web/customhostnameSites"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/customhostnameSites/mywebapp.azurewebsites.net"),
		// 			Properties: &armappservice.CustomHostnameSitesProperties{
		// 				CustomHostname: to.Ptr("mywebapp.azurewebsites.net"),
		// 				Region: to.Ptr("West US"),
		// 				SiteResourceIDs: []*armappservice.Identifier{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus-rg/providers/Microsoft.Web/sites/mywebapp"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("www.example.com"),
		// 			Type: to.Ptr("Microsoft.Web/customhostnameSites"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/customhostnameSites/www.example.com"),
		// 			Properties: &armappservice.CustomHostnameSitesProperties{
		// 				CustomHostname: to.Ptr("www.example.com"),
		// 				Region: to.Ptr("West US 2"),
		// 				SiteResourceIDs: []*armappservice.Identifier{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
