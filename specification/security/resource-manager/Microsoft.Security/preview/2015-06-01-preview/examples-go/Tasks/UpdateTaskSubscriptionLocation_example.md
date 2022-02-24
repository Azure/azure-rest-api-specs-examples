Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Tasks/UpdateTaskSubscriptionLocation_example.json
func ExampleTasksClient_UpdateSubscriptionLevelTaskState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewTasksClient("<subscription-id>",
		"<asc-location>", cred, nil)
	_, err = client.UpdateSubscriptionLevelTaskState(ctx,
		"<task-name>",
		armsecurity.TaskUpdateActionType("Dismiss"),
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
