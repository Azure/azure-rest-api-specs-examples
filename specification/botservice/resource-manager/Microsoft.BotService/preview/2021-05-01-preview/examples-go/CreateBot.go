package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/CreateBot.json
func ExampleBotsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewBotsClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"OneResourceGroupName",
		"samplebotname",
		armbotservice.Bot{
			Etag:     to.Ptr("etag1"),
			Kind:     to.Ptr(armbotservice.KindSdk),
			Location: to.Ptr("West US"),
			SKU: &armbotservice.SKU{
				Name: to.Ptr(armbotservice.SKUNameS1),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armbotservice.BotProperties{
				Description:                       to.Ptr("The description of the bot"),
				CmekKeyVaultURL:                   to.Ptr("https://myCmekKey"),
				DeveloperAppInsightKey:            to.Ptr("appinsightskey"),
				DeveloperAppInsightsAPIKey:        to.Ptr("appinsightsapikey"),
				DeveloperAppInsightsApplicationID: to.Ptr("appinsightsappid"),
				DisableLocalAuth:                  to.Ptr(true),
				DisplayName:                       to.Ptr("The Name of the bot"),
				Endpoint:                          to.Ptr("http://mybot.coffee"),
				IconURL:                           to.Ptr("http://myicon"),
				IsCmekEnabled:                     to.Ptr(true),
				LuisAppIDs: []*string{
					to.Ptr("luisappid1"),
					to.Ptr("luisappid2")},
				LuisKey:                     to.Ptr("luiskey"),
				MsaAppID:                    to.Ptr("exampleappid"),
				MsaAppMSIResourceID:         to.Ptr("/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId"),
				MsaAppTenantID:              to.Ptr("exampleapptenantid"),
				MsaAppType:                  to.Ptr(armbotservice.MsaAppTypeUserAssignedMSI),
				PublicNetworkAccess:         to.Ptr(armbotservice.PublicNetworkAccessEnabled),
				SchemaTransformationVersion: to.Ptr("1.0"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
