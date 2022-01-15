Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flogic%2Farmlogic%2Fv0.3.0/sdk/resourcemanager/logic/armlogic/README.md) on how to add the SDK to your project and authenticate.

```go
package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ValidateByResourceGroup.json
func ExampleWorkflowsClient_ValidateByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewWorkflowsClient("<subscription-id>", cred, nil)
	_, err = client.ValidateByResourceGroup(ctx,
		"<resource-group-name>",
		"<workflow-name>",
		armlogic.Workflow{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armlogic.WorkflowProperties{
				Definition: map[string]interface{}{
					"$schema":        "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
					"actions":        map[string]interface{}{},
					"contentVersion": "1.0.0.0",
					"outputs":        map[string]interface{}{},
					"parameters":     map[string]interface{}{},
					"triggers":       map[string]interface{}{},
				},
				IntegrationAccount: &armlogic.ResourceReference{
					ID: to.StringPtr("<id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
