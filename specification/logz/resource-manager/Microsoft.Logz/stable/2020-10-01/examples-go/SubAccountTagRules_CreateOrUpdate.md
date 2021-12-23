Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogz%2Farmlogz%2Fv0.1.0/sdk/resourcemanager/logz/armlogz/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_CreateOrUpdate.json
func ExampleSubAccountTagRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogz.NewSubAccountTagRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<monitor-name>",
		"<sub-account-name>",
		"<rule-set-name>",
		&armlogz.SubAccountTagRulesCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("MonitoringTagRules.ID: %s\n", *res.ID)
}
```
