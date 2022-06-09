```go
package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/CalculateTemplateHash.json
func ExampleDeploymentsClient_CalculateTemplateHash() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armresources.NewDeploymentsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CalculateTemplateHash(ctx,
		map[string]interface{}{
			"$schema":        "http://schemas.management.azure.com/deploymentTemplate?api-version=2014-04-01-preview",
			"contentVersion": "1.0.0.0",
			"outputs": map[string]interface{}{
				"string": map[string]interface{}{
					"type":  "string",
					"value": "myvalue",
				},
			},
			"parameters": map[string]interface{}{
				"string": map[string]interface{}{
					"type": "string",
				},
			},
			"resources": []interface{}{},
			"variables": map[string]interface{}{
				"array": []interface{}{
					float64(1),
					float64(2),
					float64(3),
					float64(4),
				},
				"bool": true,
				"int":  float64(42),
				"object": map[string]interface{}{
					"object": map[string]interface{}{
						"location": "West US",
						"vmSize":   "Large",
					},
				},
				"string": "string",
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmresources%2Fv1.0.0/sdk/resourcemanager/resources/armresources/README.md) on how to add the SDK to your project and authenticate.
