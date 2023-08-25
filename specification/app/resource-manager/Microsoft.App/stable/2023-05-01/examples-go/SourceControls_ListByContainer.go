package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/SourceControls_ListByContainer.json
func ExampleContainerAppsSourceControlsClient_NewListByContainerAppPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsSourceControlsClient().NewListByContainerAppPager("workerapps-rg-xj", "testcanadacentral", nil)
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
		// page.SourceControlCollection = armappcontainers.SourceControlCollection{
		// 	Value: []*armappcontainers.SourceControl{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.App/containerapps/sourcecontrols"),
		// 			ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/sourcecontrols/current"),
		// 			Properties: &armappcontainers.SourceControlProperties{
		// 				Branch: to.Ptr("master"),
		// 				GithubActionConfiguration: &armappcontainers.GithubActionConfiguration{
		// 					ContextPath: to.Ptr("./"),
		// 					Image: to.Ptr("image/tag"),
		// 					RegistryInfo: &armappcontainers.RegistryInfo{
		// 						RegistryURL: to.Ptr("xwang971reg.azurecr.io"),
		// 						RegistryUserName: to.Ptr("xwang971reg"),
		// 					},
		// 				},
		// 				RepoURL: to.Ptr("https://github.com/xwang971/ghatest"),
		// 			},
		// 	}},
		// }
	}
}
