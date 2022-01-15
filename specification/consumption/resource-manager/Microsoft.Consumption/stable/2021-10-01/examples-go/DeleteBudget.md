Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fconsumption%2Farmconsumption%2Fv0.3.0/sdk/resourcemanager/consumption/armconsumption/README.md) on how to add the SDK to your project and authenticate.

```go
package armconsumption_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
)

// x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
func ExampleBudgetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armconsumption.NewBudgetsClient(cred, nil)
	_, err = client.Delete(ctx,
		"<scope>",
		"<budget-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
