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

// x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_CreateOrUpdate.json
func ExampleWorkflowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlogic.NewWorkflowsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workflow-name>",
		armlogic.Workflow{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armlogic.WorkflowProperties{
				Definition: map[string]interface{}{
					"$schema": "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
					"actions": map[string]interface{}{
						"Find_pet_by_ID": map[string]interface{}{
							"type": "ApiConnection",
							"inputs": map[string]interface{}{
								"path":   "/pet/@{encodeURIComponent('1')}",
								"method": "get",
								"host": map[string]interface{}{
									"connection": map[string]interface{}{
										"name": "@parameters('$connections')['test-custom-connector']['connectionId']",
									},
								},
							},
							"runAfter": map[string]interface{}{},
						},
					},
					"contentVersion": "1.0.0.0",
					"outputs":        map[string]interface{}{},
					"parameters": map[string]interface{}{
						"$connections": map[string]interface{}{
							"type":         "Object",
							"defaultValue": map[string]interface{}{},
						},
					},
					"triggers": map[string]interface{}{
						"manual": map[string]interface{}{
							"type": "Request",
							"inputs": map[string]interface{}{
								"schema": map[string]interface{}{},
							},
							"kind": "Http",
						},
					},
				},
				IntegrationAccount: &armlogic.ResourceReference{
					ID: to.StringPtr("<id>"),
				},
				Parameters: map[string]*armlogic.WorkflowParameter{
					"$connections": {
						Value: map[string]interface{}{
							"test-custom-connector": map[string]interface{}{
								"connectionId":   "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/connections/test-custom-connector",
								"connectionName": "test-custom-connector",
								"id":             "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/brazilsouth/managedApis/test-custom-connector",
							},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkflowsClientCreateOrUpdateResult)
}
```
