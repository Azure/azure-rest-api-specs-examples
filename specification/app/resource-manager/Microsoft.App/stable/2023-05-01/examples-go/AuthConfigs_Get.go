package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/58be094c6b365f8d4d73a91e293dfb4818e57cf6/specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/AuthConfigs_Get.json
func ExampleContainerAppsAuthConfigsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsAuthConfigsClient().Get(ctx, "workerapps-rg-xj", "testcanadacentral", "current", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthConfig = armappcontainers.AuthConfig{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/authconfigs"),
	// 	ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/testcanadacentral/authconfigs/current"),
	// 	Properties: &armappcontainers.AuthConfigProperties{
	// 		GlobalValidation: &armappcontainers.GlobalValidation{
	// 			UnauthenticatedClientAction: to.Ptr(armappcontainers.UnauthenticatedClientActionV2AllowAnonymous),
	// 		},
	// 		IdentityProviders: &armappcontainers.IdentityProviders{
	// 			Facebook: &armappcontainers.Facebook{
	// 				Registration: &armappcontainers.AppRegistration{
	// 					AppID: to.Ptr("123"),
	// 					AppSecretSettingName: to.Ptr("facebook-secret"),
	// 				},
	// 			},
	// 		},
	// 		Platform: &armappcontainers.AuthPlatform{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// }
}
