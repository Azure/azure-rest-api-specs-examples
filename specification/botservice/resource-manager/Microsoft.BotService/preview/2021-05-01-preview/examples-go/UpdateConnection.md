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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateConnection.json
func ExampleBotConnectionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbotservice.NewBotConnectionClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<connection-name>",
		armbotservice.ConnectionSetting{
			Etag:     to.Ptr("<etag>"),
			Location: to.Ptr("<location>"),
			Properties: &armbotservice.ConnectionSettingProperties{
				ClientID:     to.Ptr("<client-id>"),
				ClientSecret: to.Ptr("<client-secret>"),
				Parameters: []*armbotservice.ConnectionSettingParameter{
					{
						Key:   to.Ptr("<key>"),
						Value: to.Ptr("<value>"),
					},
					{
						Key:   to.Ptr("<key>"),
						Value: to.Ptr("<value>"),
					}},
				Scopes:                     to.Ptr("<scopes>"),
				ServiceProviderDisplayName: to.Ptr("<service-provider-display-name>"),
				ServiceProviderID:          to.Ptr("<service-provider-id>"),
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
