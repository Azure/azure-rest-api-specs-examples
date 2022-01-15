Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_Get.json
func ExampleWorkflowTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewWorkflowTriggersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workflow-name>",
		"<trigger-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkflowTriggersClientGetResult)
}
```
