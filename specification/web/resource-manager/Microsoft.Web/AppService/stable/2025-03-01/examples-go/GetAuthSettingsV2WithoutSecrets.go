package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetAuthSettingsV2WithoutSecrets.json
func ExampleWebAppsClient_GetAuthSettingsV2WithoutSecrets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().GetAuthSettingsV2WithoutSecrets(ctx, "testrg123", "sitef6141", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SiteAuthSettingsV2 = armappservice.SiteAuthSettingsV2{
	// 	Name: to.Ptr("authsettingsv2"),
	// 	Type: to.Ptr("Microsoft.Web/sites/authsettingsv2"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/config/authsettingsv2"),
	// 	Kind: to.Ptr("app"),
	// 	Properties: &armappservice.SiteAuthSettingsV2Properties{
	// 		GlobalValidation: &armappservice.GlobalValidation{
	// 			ExcludedPaths: []*string{
	// 				to.Ptr("/nosecrets/Path")},
	// 				RequireAuthentication: to.Ptr(true),
	// 				UnauthenticatedClientAction: to.Ptr(armappservice.UnauthenticatedClientActionV2Return403),
	// 			},
	// 			HTTPSettings: &armappservice.HTTPSettings{
	// 				ForwardProxy: &armappservice.ForwardProxy{
	// 					Convention: to.Ptr(armappservice.ForwardProxyConventionStandard),
	// 					CustomHostHeaderName: to.Ptr("authHeader"),
	// 					CustomProtoHeaderName: to.Ptr("customProtoHeader"),
	// 				},
	// 				RequireHTTPS: to.Ptr(true),
	// 				Routes: &armappservice.HTTPSettingsRoutes{
	// 					APIPrefix: to.Ptr("/authv2/"),
	// 				},
	// 			},
	// 			IdentityProviders: &armappservice.IdentityProviders{
	// 				Google: &armappservice.Google{
	// 					Enabled: to.Ptr(true),
	// 					Login: &armappservice.LoginScopes{
	// 						Scopes: []*string{
	// 							to.Ptr("admin")},
	// 						},
	// 						Registration: &armappservice.ClientRegistration{
	// 							ClientID: to.Ptr("42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com"),
	// 							ClientSecretSettingName: to.Ptr("ClientSecret"),
	// 						},
	// 						Validation: &armappservice.AllowedAudiencesValidation{
	// 							AllowedAudiences: []*string{
	// 								to.Ptr("https://example.com")},
	// 							},
	// 						},
	// 					},
	// 					Login: &armappservice.Login{
	// 						AllowedExternalRedirectUrls: []*string{
	// 							to.Ptr("https://someurl.com")},
	// 							CookieExpiration: &armappservice.CookieExpiration{
	// 								Convention: to.Ptr(armappservice.CookieExpirationConventionIdentityProviderDerived),
	// 								TimeToExpiration: to.Ptr("2022:09-01T00:00Z"),
	// 							},
	// 							Nonce: &armappservice.Nonce{
	// 								ValidateNonce: to.Ptr(true),
	// 							},
	// 							PreserveURLFragmentsForLogins: to.Ptr(true),
	// 							Routes: &armappservice.LoginRoutes{
	// 								LogoutEndpoint: to.Ptr("https://app.com/logout"),
	// 							},
	// 							TokenStore: &armappservice.TokenStore{
	// 								Enabled: to.Ptr(true),
	// 								FileSystem: &armappservice.FileSystemTokenStore{
	// 									Directory: to.Ptr("/wwwroot/sites/example"),
	// 								},
	// 								TokenRefreshExtensionHours: to.Ptr[float64](96),
	// 							},
	// 						},
	// 						Platform: &armappservice.AuthPlatform{
	// 							ConfigFilePath: to.Ptr("/auth/config.json"),
	// 							Enabled: to.Ptr(true),
	// 							RuntimeVersion: to.Ptr("~1"),
	// 						},
	// 					},
	// 				}
}
