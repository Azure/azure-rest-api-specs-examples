package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateAuthSettings.json
func ExampleWebAppsClient_UpdateAuthSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().UpdateAuthSettings(ctx, "testrg123", "sitef6141", armappservice.SiteAuthSettings{
		Properties: &armappservice.SiteAuthSettingsProperties{
			AllowedExternalRedirectUrls: []*string{
				to.Ptr("sitef6141.customdomain.net"),
				to.Ptr("sitef6141.customdomain.info")},
			ClientID:                    to.Ptr("42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com"),
			DefaultProvider:             to.Ptr(armappservice.BuiltInAuthenticationProviderGoogle),
			Enabled:                     to.Ptr(true),
			RuntimeVersion:              to.Ptr("~1"),
			TokenRefreshExtensionHours:  to.Ptr[float64](120),
			TokenStoreEnabled:           to.Ptr(true),
			UnauthenticatedClientAction: to.Ptr(armappservice.UnauthenticatedClientActionRedirectToLoginPage),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SiteAuthSettings = armappservice.SiteAuthSettings{
	// 	Name: to.Ptr("authsettings"),
	// 	Type: to.Ptr("Microsoft.Web/sites/authsettings"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/sitef6141/config/authsettings"),
	// 	Kind: to.Ptr("app"),
	// 	Properties: &armappservice.SiteAuthSettingsProperties{
	// 		AllowedExternalRedirectUrls: []*string{
	// 			to.Ptr("sitef6141.customdomain.net"),
	// 			to.Ptr("sitef6141.customdomain.info")},
	// 			ClientID: to.Ptr("42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com"),
	// 			DefaultProvider: to.Ptr(armappservice.BuiltInAuthenticationProviderGoogle),
	// 			Enabled: to.Ptr(true),
	// 			RuntimeVersion: to.Ptr("~1"),
	// 			TokenRefreshExtensionHours: to.Ptr[float64](120),
	// 			TokenStoreEnabled: to.Ptr(true),
	// 			UnauthenticatedClientAction: to.Ptr(armappservice.UnauthenticatedClientActionRedirectToLoginPage),
	// 		},
	// 	}
}
