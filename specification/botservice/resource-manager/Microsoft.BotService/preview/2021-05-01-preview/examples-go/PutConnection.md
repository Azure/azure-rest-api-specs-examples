```go
package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutConnection.json
func ExampleBotConnectionClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewBotConnectionClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"OneResourceGroupName",
		"samplebotname",
		"sampleConnection",
		armbotservice.ConnectionSetting{
			Etag:     to.Ptr("etag1"),
			Location: to.Ptr("West US"),
			Properties: &armbotservice.ConnectionSettingProperties{
				ClientID:     to.Ptr("sampleclientid"),
				ClientSecret: to.Ptr("samplesecret"),
				Parameters: []*armbotservice.ConnectionSettingParameter{
					{
						Key:   to.Ptr("key1"),
						Value: to.Ptr("value1"),
					},
					{
						Key:   to.Ptr("key2"),
						Value: to.Ptr("value2"),
					}},
				Scopes:            to.Ptr("samplescope"),
				ServiceProviderID: to.Ptr("serviceproviderid"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbotservice%2Farmbotservice%2Fv0.5.0/sdk/resourcemanager/botservice/armbotservice/README.md) on how to add the SDK to your project and authenticate.
