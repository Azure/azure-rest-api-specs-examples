Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbotservice%2Farmbotservice%2Fv0.4.0/sdk/resourcemanager/botservice/armbotservice/README.md) on how to add the SDK to your project and authenticate.

```go
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
		return
	}
	ctx := context.Background()
	client, err := armbotservice.NewBotsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armbotservice.Bot{
			Etag:     to.Ptr("<etag>"),
			Kind:     to.Ptr(armbotservice.KindSdk),
			Location: to.Ptr("<location>"),
			SKU: &armbotservice.SKU{
				Name: to.Ptr(armbotservice.SKUNameS1),
			},
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armbotservice.BotProperties{
				Description:                       to.Ptr("<description>"),
				CmekKeyVaultURL:                   to.Ptr("<cmek-key-vault-url>"),
				DeveloperAppInsightKey:            to.Ptr("<developer-app-insight-key>"),
				DeveloperAppInsightsAPIKey:        to.Ptr("<developer-app-insights-apikey>"),
				DeveloperAppInsightsApplicationID: to.Ptr("<developer-app-insights-application-id>"),
				DisableLocalAuth:                  to.Ptr(true),
				DisplayName:                       to.Ptr("<display-name>"),
				Endpoint:                          to.Ptr("<endpoint>"),
				IconURL:                           to.Ptr("<icon-url>"),
				IsCmekEnabled:                     to.Ptr(true),
				LuisAppIDs: []*string{
					to.Ptr("luisappid1"),
					to.Ptr("luisappid2")},
				LuisKey:                     to.Ptr("<luis-key>"),
				MsaAppID:                    to.Ptr("<msa-app-id>"),
				MsaAppMSIResourceID:         to.Ptr("<msa-app-msiresource-id>"),
				MsaAppTenantID:              to.Ptr("<msa-app-tenant-id>"),
				MsaAppType:                  to.Ptr(armbotservice.MsaAppTypeUserAssignedMSI),
				PublicNetworkAccess:         to.Ptr(armbotservice.PublicNetworkAccessEnabled),
				SchemaTransformationVersion: to.Ptr("<schema-transformation-version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
