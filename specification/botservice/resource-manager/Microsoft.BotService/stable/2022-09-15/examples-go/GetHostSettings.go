package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/GetHostSettings.json
func ExampleHostSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHostSettingsClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HostSettingsResponse = armbotservice.HostSettingsResponse{
	// 	BotOpenIDMetadata: to.Ptr("https://login.botframework.com/v1/.well-known/openidconfiguration"),
	// 	OAuthURL: to.Ptr("https://token.botframework.com/"),
	// 	ToBotFromChannelOpenIDMetadataURL: to.Ptr("https://login.botframework.com/v1/.well-known/openidconfiguration"),
	// 	ToBotFromChannelTokenIssuer: to.Ptr("https://api.botframework.com"),
	// 	ToBotFromEmulatorOpenIDMetadataURL: to.Ptr("https://login.microsoftonline.com/botframework.com/v2.0/.well-known/openid-configuration"),
	// 	ToChannelFromBotLoginURL: to.Ptr("https://login.microsoftonline.com/botframework.com"),
	// 	ToChannelFromBotOAuthScope: to.Ptr("https://api.botframework.com"),
	// 	ValidateAuthority: to.Ptr(true),
	// }
}
