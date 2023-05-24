package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a7af0df86022e5e6cc6e8f40ca1981c4557a4bc/specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/AuthConfigs_ListByContainer.json
func ExampleContainerAppsAuthConfigsClient_NewListByContainerAppPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainerAppsAuthConfigsClient().NewListByContainerAppPager("workerapps-rg-xj", "testcanadacentral", nil)
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
		// page.AuthConfigCollection = armappcontainers.AuthConfigCollection{
		// 	Value: []*armappcontainers.AuthConfig{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.App/containerapps/authconfigs"),
		// 			ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/authconfigs/current"),
		// 			Properties: &armappcontainers.AuthConfigProperties{
		// 				GlobalValidation: &armappcontainers.GlobalValidation{
		// 					UnauthenticatedClientAction: to.Ptr(armappcontainers.UnauthenticatedClientActionV2AllowAnonymous),
		// 				},
		// 				IdentityProviders: &armappcontainers.IdentityProviders{
		// 					Facebook: &armappcontainers.Facebook{
		// 						Registration: &armappcontainers.AppRegistration{
		// 							AppID: to.Ptr("123"),
		// 							AppSecretSettingName: to.Ptr("facebook-secret"),
		// 						},
		// 					},
		// 				},
		// 				Platform: &armappcontainers.AuthPlatform{
		// 					Enabled: to.Ptr(true),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
