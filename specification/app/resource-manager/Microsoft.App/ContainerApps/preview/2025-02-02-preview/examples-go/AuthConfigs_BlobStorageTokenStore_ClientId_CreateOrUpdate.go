package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/AuthConfigs_BlobStorageTokenStore_ClientId_CreateOrUpdate.json
func ExampleContainerAppsAuthConfigsClient_CreateOrUpdate_createOrUpdateContainerAppAuthConfigWithMsiClientIdBlobStorageTokenStore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewContainerAppsAuthConfigsClient().CreateOrUpdate(ctx, "rg1", "myapp", "current", armappcontainers.AuthConfig{
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
			Login: &armappcontainers.Login{
				TokenStore: &armappcontainers.TokenStore{
					AzureBlobStorage: &armappcontainers.BlobStorageTokenStore{
						BlobContainerURI: to.Ptr("https://test.blob.core.windows.net/container1"),
						ClientID:         to.Ptr("00000000-0000-0000-0000-000000000000"),
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
	// res.AuthConfig = armappcontainers.AuthConfig{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.App/containerapps/authconfigs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.App/containerApps/myapp/authconfigs/current"),
	// 	Properties: &armappcontainers.AuthConfigProperties{
	// 		EncryptionSettings: &armappcontainers.EncryptionSettings{
	// 			ContainerAppAuthEncryptionSecretName: to.Ptr("testEncryptionSecretName"),
	// 			ContainerAppAuthSigningSecretName: to.Ptr("testSigningSecretName"),
	// 		},
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
	// 		Login: &armappcontainers.Login{
	// 			TokenStore: &armappcontainers.TokenStore{
	// 				AzureBlobStorage: &armappcontainers.BlobStorageTokenStore{
	// 					BlobContainerURI: to.Ptr("https://test.blob.core.windows.net/container1"),
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Platform: &armappcontainers.AuthPlatform{
	// 			Enabled: to.Ptr(true),
	// 		},
	// 	},
	// }
}
