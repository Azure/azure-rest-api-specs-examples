Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappcontainers%2Farmappcontainers%2Fv1.0.0/sdk/resourcemanager/appcontainers/armappcontainers/README.md) on how to add the SDK to your project and authenticate.

```go
package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/AuthConfigs_CreateOrUpdate.json
func ExampleContainerAppsAuthConfigsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armappcontainers.NewContainerAppsAuthConfigsClient("651f8027-33e8-4ec4-97b4-f6e9f3dc8744", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"workerapps-rg-xj",
		"testcanadacentral",
		"current",
		armappcontainers.AuthConfig{
			Properties: &armappcontainers.AuthConfigProperties{
				GlobalValidation: &armappcontainers.GlobalValidation{
					UnauthenticatedClientAction: to.Ptr(armappcontainers.UnauthenticatedClientActionV2AllowAnonymous),
				},
				IdentityProviders: &armappcontainers.IdentityProviders{
					Facebook: &armappcontainers.Facebook{
						Registration: &armappcontainers.AppRegistration{
							AppID:                to.Ptr("123"),
							AppSecretSettingName: to.Ptr("facebook-secret"),
						},
					},
				},
				Platform: &armappcontainers.AuthPlatform{
					Enabled: to.Ptr(true),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
