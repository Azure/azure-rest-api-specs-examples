Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Falertsmanagement%2Farmalertsmanagement%2Fv0.3.0/sdk/resourcemanager/alertsmanagement/armalertsmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// x-ms-original-file: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_ChangeState.json
func ExampleSmartGroupsClient_ChangeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armalertsmanagement.NewSmartGroupsClient("<subscription-id>", cred, nil)
	res, err := client.ChangeState(ctx,
		"<smart-group-id>",
		armalertsmanagement.AlertState("Acknowledged"),
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SmartGroupsClientChangeStateResult)
}
```
