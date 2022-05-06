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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateAlexaChannel.json
func ExampleChannelsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbotservice.NewChannelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armbotservice.ChannelNameAlexaChannel,
		armbotservice.BotChannel{
			Location: to.Ptr("<location>"),
			Properties: &armbotservice.AlexaChannel{
				ChannelName: to.Ptr("<channel-name>"),
				Properties: &armbotservice.AlexaChannelProperties{
					AlexaSkillID: to.Ptr("<alexa-skill-id>"),
					IsEnabled:    to.Ptr(true),
				},
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
