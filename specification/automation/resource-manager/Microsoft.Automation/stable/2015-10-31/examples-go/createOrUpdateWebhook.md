Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fautomation%2Farmautomation%2Fv0.5.0/sdk/resourcemanager/automation/armautomation/README.md) on how to add the SDK to your project and authenticate.

```go
package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/createOrUpdateWebhook.json
func ExampleWebhookClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewWebhookClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<webhook-name>",
		armautomation.WebhookCreateOrUpdateParameters{
			Name: to.Ptr("<name>"),
			Properties: &armautomation.WebhookCreateOrUpdateProperties{
				ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T22:18:13.7002872Z"); return t }()),
				IsEnabled:  to.Ptr(true),
				Runbook: &armautomation.RunbookAssociationProperty{
					Name: to.Ptr("<name>"),
				},
				URI: to.Ptr("<uri>"),
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
