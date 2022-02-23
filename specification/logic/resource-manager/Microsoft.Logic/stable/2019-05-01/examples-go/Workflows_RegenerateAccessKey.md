Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.1/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_RegenerateAccessKey.json
func ExampleWorkflowsClient_RegenerateAccessKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewWorkflowsClient("<subscription-id>", cred, nil)
	_, err = client.RegenerateAccessKey(ctx,
		"<resource-group-name>",
		"<workflow-name>",
		armlogic.RegenerateActionParameter{
			KeyType: armlogic.KeyType("Primary").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
