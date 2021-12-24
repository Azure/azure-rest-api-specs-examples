Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatadog%2Farmdatadog%2Fv0.1.0/sdk/resourcemanager/datadog/armdatadog/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatadog_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
)

// x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/TagRules_CreateOrUpdate.json
func ExampleTagRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatadog.NewTagRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<rule-set-name>",
		&armdatadog.TagRulesCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MonitoringTagRules.ID: %s\n", *res.ID)
}
```
