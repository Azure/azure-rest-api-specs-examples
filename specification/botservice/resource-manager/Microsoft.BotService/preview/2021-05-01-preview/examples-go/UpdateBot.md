Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbotservice%2Farmbotservice%2Fv0.2.0/sdk/resourcemanager/botservice/armbotservice/README.md) on how to add the SDK to your project and authenticate.

```go
package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateBot.json
func ExampleBotsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbotservice.NewBotsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armbotservice.Bot{
			Etag:     to.StringPtr("<etag>"),
			Kind:     armbotservice.Kind("sdk").ToPtr(),
			Location: to.StringPtr("<location>"),
			SKU: &armbotservice.SKU{
				Name: armbotservice.SKUName("S1").ToPtr(),
			},
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armbotservice.BotProperties{
				Description:                       to.StringPtr("<description>"),
				CmekKeyVaultURL:                   to.StringPtr("<cmek-key-vault-url>"),
				DeveloperAppInsightKey:            to.StringPtr("<developer-app-insight-key>"),
				DeveloperAppInsightsAPIKey:        to.StringPtr("<developer-app-insights-apikey>"),
				DeveloperAppInsightsApplicationID: to.StringPtr("<developer-app-insights-application-id>"),
				DisableLocalAuth:                  to.BoolPtr(true),
				DisplayName:                       to.StringPtr("<display-name>"),
				Endpoint:                          to.StringPtr("<endpoint>"),
				IconURL:                           to.StringPtr("<icon-url>"),
				IsCmekEnabled:                     to.BoolPtr(true),
				LuisAppIDs: []*string{
					to.StringPtr("luisappid1"),
					to.StringPtr("luisappid2")},
				LuisKey:                     to.StringPtr("<luis-key>"),
				MsaAppID:                    to.StringPtr("<msa-app-id>"),
				MsaAppMSIResourceID:         to.StringPtr("<msa-app-msiresource-id>"),
				MsaAppTenantID:              to.StringPtr("<msa-app-tenant-id>"),
				MsaAppType:                  armbotservice.MsaAppType("UserAssignedMSI").ToPtr(),
				PublicNetworkAccess:         armbotservice.PublicNetworkAccess("Enabled").ToPtr(),
				SchemaTransformationVersion: to.StringPtr("<schema-transformation-version>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BotsClientUpdateResult)
}
```
