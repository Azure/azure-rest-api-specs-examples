Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Felastic%2Farmelastic%2Fv0.1.0/sdk/resourcemanager/elastic/armelastic/README.md) on how to add the SDK to your project and authenticate.

```go
package armelastic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
)

// x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2020-07-01/examples/TagRules_Get.json
func ExampleTagRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armelastic.NewTagRulesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<rule-set-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MonitoringTagRules.ID: %s\n", *res.ID)
}
```
