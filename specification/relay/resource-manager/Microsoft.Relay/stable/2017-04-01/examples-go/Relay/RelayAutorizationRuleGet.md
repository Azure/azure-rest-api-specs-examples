Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frelay%2Farmrelay%2Fv0.3.1/sdk/resourcemanager/relay/armrelay/README.md) on how to add the SDK to your project and authenticate.

```go
package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleGet.json
func ExampleWCFRelaysClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrelay.NewWCFRelaysClient("<subscription-id>", cred, nil)
	res, err := client.GetAuthorizationRule(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<relay-name>",
		"<authorization-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WCFRelaysClientGetAuthorizationRuleResult)
}
```
