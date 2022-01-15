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

// x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateConnection.json
func ExampleBotConnectionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbotservice.NewBotConnectionClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<connection-name>",
		armbotservice.ConnectionSetting{
			Etag:     to.StringPtr("<etag>"),
			Location: to.StringPtr("<location>"),
			Properties: &armbotservice.ConnectionSettingProperties{
				ClientID:     to.StringPtr("<client-id>"),
				ClientSecret: to.StringPtr("<client-secret>"),
				Parameters: []*armbotservice.ConnectionSettingParameter{
					{
						Key:   to.StringPtr("<key>"),
						Value: to.StringPtr("<value>"),
					},
					{
						Key:   to.StringPtr("<key>"),
						Value: to.StringPtr("<value>"),
					}},
				Scopes:                     to.StringPtr("<scopes>"),
				ServiceProviderDisplayName: to.StringPtr("<service-provider-display-name>"),
				ServiceProviderID:          to.StringPtr("<service-provider-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BotConnectionClientUpdateResult)
}
```
