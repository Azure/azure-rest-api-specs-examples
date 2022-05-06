Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhealthbot%2Farmhealthbot%2Fv0.4.0/sdk/resourcemanager/healthbot/armhealthbot/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ResourceCreationPut.json
func ExampleBotsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhealthbot.NewBotsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<bot-name>",
		armhealthbot.HealthBot{
			Location: to.Ptr("<location>"),
			Identity: &armhealthbot.Identity{
				Type: to.Ptr(armhealthbot.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armhealthbot.UserAssignedIdentity{
					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi":  {},
					"/subscriptions/subscription-id/resourcegroups/myrg/providers/microsoft.managedidentity/userassignedidentities/my-mi2": {},
				},
			},
			SKU: &armhealthbot.SKU{
				Name: to.Ptr(armhealthbot.SKUNameF0),
			},
		},
		&armhealthbot.BotsClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
