Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-01-01/examples/Alerts/GetAlertSubscriptionLocation_example.json
func ExampleAlertsClient_GetSubscriptionLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAlertsClient("<subscription-id>",
		"<asc-location>", cred, nil)
	res, err := client.GetSubscriptionLevel(ctx,
		"<alert-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AlertsClientGetSubscriptionLevelResult)
}
```
