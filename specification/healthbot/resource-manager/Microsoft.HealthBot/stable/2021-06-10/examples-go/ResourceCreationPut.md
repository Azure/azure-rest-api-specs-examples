Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthbot%2Farmhealthbot%2Fv0.2.1/sdk/resourcemanager/healthbot/armhealthbot/README.md) on how to add the SDK to your project and authenticate.

```go
package armhealthbot_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot"
)

// x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ResourceCreationPut.json
func ExampleBotsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthbot.NewBotsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<bot-name>",
		armhealthbot.HealthBot{
			Location: to.StringPtr("<location>"),
			Identity: &armhealthbot.Identity{
				Type: armhealthbot.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armhealthbot.UserAssignedIdentity{
					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi":  {},
					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2": {},
				},
			},
			SKU: &armhealthbot.SKU{
				Name: armhealthbot.SKUNameF0.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.BotsClientCreateResult)
}
```
