package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/AuthConfigs_CreateOrUpdate.json
func ExampleContainerAppsAuthConfigsClient_CreateOrUpdate_createOrUpdateContainerAppAuthConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("651f8027-33e8-4ec4-97b4-f6e9f3dc8744", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsAuthConfigsClient().CreateOrUpdate(ctx, "workerapps-rg-xj", "testcanadacentral", "current", armappcontainers.AuthConfig{
		Properties: &armappcontainers.AuthConfigProperties{
			EncryptionSettings: &armappcontainers.EncryptionSettings{
				ContainerAppAuthEncryptionSecretName: to.Ptr("testEncryptionSecretName"),
				ContainerAppAuthSigningSecretName:    to.Ptr("testSigningSecretName"),
			},
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.ContainerAppsAuthConfigsClientCreateOrUpdateResponse{
	// 	AuthConfig: armappcontainers.AuthConfig{
	// 		Name: to.Ptr("current"),
	// 		Type: to.Ptr("Microsoft.App/containerapps/authconfigs"),
	// 		ID: to.Ptr("/subscriptions/651f8027-33e8-4ec4-97b4-f6e9f3dc8744/resourceGroups/workerapps-rg-xj/providers/Microsoft.App/containerApps/myapp/authconfigs/current"),
	// 		Properties: &armappcontainers.AuthConfigProperties{
	// 			EncryptionSettings: &armappcontainers.EncryptionSettings{
	// 				ContainerAppAuthEncryptionSecretName: to.Ptr("testEncryptionSecretName"),
	// 				ContainerAppAuthSigningSecretName: to.Ptr("testSigningSecretName"),
	// 			},
	// 			GlobalValidation: &armappcontainers.GlobalValidation{
	// 				UnauthenticatedClientAction: to.Ptr(armappcontainers.UnauthenticatedClientActionV2AllowAnonymous),
	// 			},
	// 			IdentityProviders: &armappcontainers.IdentityProviders{
	// 				Facebook: &armappcontainers.Facebook{
	// 					Registration: &armappcontainers.AppRegistration{
	// 						AppID: to.Ptr("123"),
	// 						AppSecretSettingName: to.Ptr("facebook-secret"),
	// 					},
	// 				},
	// 			},
	// 			Platform: &armappcontainers.AuthPlatform{
	// 				Enabled: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
