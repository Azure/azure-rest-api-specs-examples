package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/CreateBot.json
func ExampleBotsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBotsClient().Create(ctx, "OneResourceGroupName", "samplebotname", armbotservice.Bot{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Bot = armbotservice.Bot{
	// 	Name: to.Ptr("samplebotname"),
	// 	Type: to.Ptr("Microsoft.BotService/botServices"),
	// 	Etag: to.Ptr("etag1"),
	// 	ID: to.Ptr("/subscriptions/subscription-id/resourceGroups/OneResourceGroupName/providers/Microsoft.BotService/botServices/samplebotname"),
	// 	Kind: to.Ptr(armbotservice.KindSdk),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armbotservice.BotProperties{
	// 		Description: to.Ptr("The description of the bot"),
	// 		CmekKeyVaultURL: to.Ptr("https://myCmekKey"),
	// 		ConfiguredChannels: []*string{
	// 			to.Ptr("facebook"),
	// 			to.Ptr("groupme")},
	// 			DeveloperAppInsightKey: to.Ptr("appinsightskey"),
	// 			DeveloperAppInsightsApplicationID: to.Ptr("appinsightsappid"),
	// 			DisableLocalAuth: to.Ptr(true),
	// 			DisplayName: to.Ptr("The Name of the bot"),
	// 			EnabledChannels: []*string{
	// 				to.Ptr("facebook")},
	// 				Endpoint: to.Ptr("http://mybot.coffee"),
	// 				EndpointVersion: to.Ptr("version"),
	// 				IconURL: to.Ptr("http://myicon"),
	// 				IsCmekEnabled: to.Ptr(true),
	// 				LuisAppIDs: []*string{
	// 					to.Ptr("luisappid1"),
	// 					to.Ptr("luisappid2")},
	// 					MsaAppID: to.Ptr("msaappid"),
	// 					MsaAppMSIResourceID: to.Ptr("/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId"),
	// 					MsaAppTenantID: to.Ptr("msaapptenantid"),
	// 					MsaAppType: to.Ptr(armbotservice.MsaAppTypeUserAssignedMSI),
	// 					PublicNetworkAccess: to.Ptr(armbotservice.PublicNetworkAccessEnabled),
	// 					SchemaTransformationVersion: to.Ptr("1.0"),
	// 				},
	// 			}
}
